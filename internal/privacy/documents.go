package privacy

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"
)

const maxExtractedDocumentBytes = 64 << 20

type ExtractedDocument struct {
	ContentType string
	Text        string
	Extractor   string
}

func ExtractFile(path string) (ExtractedDocument, error) {
	extension := strings.ToLower(filepath.Ext(path))
	switch extension {
	case ".txt", ".md", ".csv", ".json", ".xml", ".yaml", ".yml", ".html", ".htm", ".log":
		data, err := os.ReadFile(path)
		if err != nil {
			return ExtractedDocument{}, err
		}
		if len(data) > maxExtractedDocumentBytes {
			return ExtractedDocument{}, errors.New("document exceeds 64 MiB extraction limit")
		}
		if !utf8.Valid(data) {
			return ExtractedDocument{}, errors.New("text document is not valid UTF-8")
		}
		return ExtractedDocument{ContentType: inferTextMIME(extension), Text: string(data), Extractor: "native-text"}, nil
	case ".docx", ".pptx", ".xlsx":
		text, err := extractOpenXML(path, extension)
		if err != nil {
			return ExtractedDocument{}, err
		}
		return ExtractedDocument{ContentType: "text/plain", Text: text, Extractor: "native-openxml"}, nil
	case ".pdf":
		text, err := runLocalExtractor("pdftotext", []string{"-layout", "-nopgbrk", path, "-"})
		if err != nil {
			return ExtractedDocument{}, fmt.Errorf("extract PDF locally: %w", err)
		}
		return ExtractedDocument{ContentType: "text/plain", Text: text, Extractor: "pdftotext"}, nil
	case ".png", ".jpg", ".jpeg", ".tif", ".tiff", ".webp":
		text, err := runLocalExtractor("tesseract", []string{path, "stdout"})
		if err != nil {
			return ExtractedDocument{}, fmt.Errorf("OCR image locally: %w", err)
		}
		return ExtractedDocument{ContentType: "text/plain", Text: text, Extractor: "tesseract"}, nil
	default:
		return ExtractedDocument{}, fmt.Errorf("unsupported document format %q", extension)
	}
}

func extractOpenXML(path, extension string) (string, error) {
	archive, err := zip.OpenReader(path)
	if err != nil {
		return "", fmt.Errorf("open Office document: %w", err)
	}
	defer archive.Close()

	var files []*zip.File
	for _, file := range archive.File {
		if openXMLPartAllowed(file.Name, extension) {
			files = append(files, file)
		}
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name < files[j].Name })
	if len(files) == 0 {
		return "", errors.New("Office document contains no extractable text parts")
	}

	var output strings.Builder
	for _, file := range files {
		reader, err := file.Open()
		if err != nil {
			return "", fmt.Errorf("open Office part %s: %w", file.Name, err)
		}
		part, err := extractXMLText(io.LimitReader(reader, maxExtractedDocumentBytes))
		_ = reader.Close()
		if err != nil {
			return "", fmt.Errorf("extract Office part %s: %w", file.Name, err)
		}
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if output.Len() > 0 {
			output.WriteString("\n\n")
		}
		output.WriteString(part)
		if output.Len() > maxExtractedDocumentBytes {
			return "", errors.New("extracted Office text exceeds 64 MiB limit")
		}
	}
	if strings.TrimSpace(output.String()) == "" {
		return "", errors.New("Office document yielded no text")
	}
	return output.String(), nil
}

func openXMLPartAllowed(name, extension string) bool {
	name = strings.ToLower(name)
	if !strings.HasSuffix(name, ".xml") {
		return false
	}
	switch extension {
	case ".docx":
		return strings.HasPrefix(name, "word/document.xml") || strings.HasPrefix(name, "word/header") || strings.HasPrefix(name, "word/footer") || strings.HasPrefix(name, "word/comments") || strings.HasPrefix(name, "word/footnotes") || strings.HasPrefix(name, "word/endnotes")
	case ".pptx":
		return strings.HasPrefix(name, "ppt/slides/") || strings.HasPrefix(name, "ppt/notesslides/") || strings.HasPrefix(name, "ppt/comments/")
	case ".xlsx":
		return name == "xl/sharedstrings.xml" || strings.HasPrefix(name, "xl/worksheets/") || strings.HasPrefix(name, "xl/comments")
	default:
		return false
	}
}

func extractXMLText(reader io.Reader) (string, error) {
	decoder := xml.NewDecoder(reader)
	var output strings.Builder
	var captureDepth int
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		switch value := token.(type) {
		case xml.StartElement:
			if isTextElement(value.Name.Local) {
				captureDepth++
			}
		case xml.EndElement:
			if isTextElement(value.Name.Local) && captureDepth > 0 {
				captureDepth--
				if output.Len() > 0 && !strings.HasSuffix(output.String(), " ") {
					output.WriteByte(' ')
				}
			}
		case xml.CharData:
			if captureDepth > 0 {
				text := strings.TrimSpace(string(value))
				if text != "" {
					if output.Len() > 0 && !strings.HasSuffix(output.String(), " ") {
						output.WriteByte(' ')
					}
					output.WriteString(text)
				}
			}
		}
	}
	return strings.Join(strings.Fields(output.String()), " "), nil
}

func isTextElement(local string) bool {
	switch strings.ToLower(local) {
	case "t", "v", "f", "text":
		return true
	default:
		return false
	}
}

func runLocalExtractor(binary string, args []string) (string, error) {
	path, err := exec.LookPath(binary)
	if err != nil {
		return "", fmt.Errorf("%s is not installed; fail-closed prevents forwarding the original document", binary)
	}
	command := exec.Command(path, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	if err := command.Run(); err != nil {
		return "", fmt.Errorf("%s failed: %s", binary, strings.TrimSpace(stderr.String()))
	}
	if stdout.Len() > maxExtractedDocumentBytes {
		return "", errors.New("extracted document exceeds 64 MiB limit")
	}
	text := strings.TrimSpace(stdout.String())
	if text == "" {
		return "", fmt.Errorf("%s produced no text; OCR or another local extractor is required", binary)
	}
	if !utf8.ValidString(text) {
		return "", fmt.Errorf("%s output is not valid UTF-8", binary)
	}
	return text, nil
}

func inferTextMIME(extension string) string {
	switch extension {
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".csv":
		return "text/csv"
	case ".md":
		return "text/markdown"
	case ".html", ".htm":
		return "text/html"
	default:
		return "text/plain"
	}
}
