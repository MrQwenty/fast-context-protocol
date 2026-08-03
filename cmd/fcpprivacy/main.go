package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/MrQwenty/fast-context-protocol/internal/privacy"
	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

func main() {
	inputPath := flag.String("input", "-", "input document path or - for stdin")
	outputPath := flag.String("output", "-", "sanitized output path or - for stdout")
	reportPath := flag.String("report", "", "privacy receipt JSON path")
	vaultPath := flag.String("vault", "", "encrypted local mapping path for reversible pseudonymization")
	documentID := flag.String("document-id", "", "stable local document identifier")
	contentType := flag.String("content-type", "", "MIME type; inferred for supported files")
	mode := flag.String("mode", "anonymize", "redact, pseudonymize, or anonymize")
	scope := flag.String("scope", "", "pseudonymization scope")
	customTermsPath := flag.String("custom-terms", "", "newline-delimited sensitive terms")
	allowListPath := flag.String("allow-list", "", "newline-delimited values that must remain")
	failClosed := flag.Bool("fail-closed", true, "refuse output when residual high-risk identifiers remain")
	reversible := flag.Bool("reversible", false, "write an encrypted local re-identification vault")
	maxResidualRisk := flag.Float64("max-residual-risk", 0, "maximum accepted post-scan risk from 0 to 1")
	flag.Parse()

	content, extractedType, err := readInput(*inputPath)
	fatalIf(err)
	if *documentID == "" {
		if *inputPath == "-" {
			*documentID = "stdin"
		} else {
			*documentID = filepath.Base(*inputPath)
		}
	}
	if *contentType == "" {
		*contentType = extractedType
	}
	customTerms, err := readLines(*customTermsPath)
	fatalIf(err)
	allowList, err := readLines(*allowListPath)
	fatalIf(err)

	secret := []byte(os.Getenv("FCP_PRIVACY_SECRET"))
	vaultKey, err := decodeVaultKey(os.Getenv("FCP_PRIVACY_VAULT_KEY"))
	fatalIf(err)
	policy := protocol.PrivacyPolicy{
		Mode: protocol.PrivacyMode(*mode), LocalOnly: true, FailClosed: *failClosed,
		Reversible: *reversible, ScopeID: *scope, CustomTerms: customTerms, AllowList: allowList,
		MinimumConfidence: 0.85, MaxResidualRisk: *maxResidualRisk, VaultKeyID: "local-env-key",
	}
	response, sanitizeErr := privacy.NewEngine(secret, vaultKey).Sanitize(protocol.SanitizeRequest{
		DocumentID: *documentID, ContentType: *contentType, Content: string(content), Policy: policy,
	})
	if *reportPath != "" {
		fatalIf(writeJSON(*reportPath, response.Receipt))
	}
	if response.Vault != nil {
		if *vaultPath == "" {
			fatalIf(errors.New("-vault is required when -reversible is enabled"))
		}
		fatalIf(writeJSON(*vaultPath, response.Vault))
	}
	if sanitizeErr != nil {
		fatalIf(sanitizeErr)
	}
	fatalIf(writeOutput(*outputPath, []byte(response.Content)))
}

func readInput(path string) ([]byte, string, error) {
	if path == "-" {
		data, err := io.ReadAll(io.LimitReader(os.Stdin, 64<<20))
		return data, "text/plain", err
	}
	document, err := privacy.ExtractFile(path)
	if err != nil {
		return nil, "", err
	}
	return []byte(document.Text), document.ContentType, nil
}

func writeOutput(path string, data []byte) error {
	if path == "-" {
		_, err := os.Stdout.Write(data)
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

func writeJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o600)
}

func readLines(path string) ([]string, error) {
	if path == "" {
		return nil, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			out = append(out, line)
		}
	}
	return out, nil
}

func decodeVaultKey(value string) ([]byte, error) {
	if value == "" {
		return nil, nil
	}
	if decoded, err := base64.RawURLEncoding.DecodeString(value); err == nil && len(decoded) == 32 {
		return decoded, nil
	}
	sum := sha256.Sum256([]byte(value))
	return sum[:], nil
}

func fatalIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "fcpprivacy:", err)
	os.Exit(1)
}
