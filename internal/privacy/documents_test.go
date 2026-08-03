package privacy

import (
	"archive/zip"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractTextFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sample.md")
	if err := os.WriteFile(path, []byte("Hello Matteo"), 0o600); err != nil {
		t.Fatal(err)
	}
	document, err := ExtractFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if document.Text != "Hello Matteo" || document.Extractor != "native-text" {
		t.Fatalf("unexpected extraction: %+v", document)
	}
}

func TestExtractDOCXTextWithoutForwardingBinary(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sample.docx")
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	archive := zip.NewWriter(file)
	part, err := archive.Create("word/document.xml")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = part.Write([]byte(`<?xml version="1.0"?><w:document xmlns:w="urn:test"><w:body><w:p><w:r><w:t>Mario Rossi</w:t></w:r></w:p><w:p><w:r><w:t>mario@example.com</w:t></w:r></w:p></w:body></w:document>`))
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
	document, err := ExtractFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(document.Text, "Mario Rossi") || !strings.Contains(document.Text, "mario@example.com") {
		t.Fatalf("unexpected DOCX text: %q", document.Text)
	}
	if document.Extractor != "native-openxml" {
		t.Fatalf("unexpected extractor %q", document.Extractor)
	}
}
