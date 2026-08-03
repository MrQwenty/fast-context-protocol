package protocol

import "time"

type PrivacyMode string

const (
	PrivacyModeNone         PrivacyMode = "none"
	PrivacyModeRedact       PrivacyMode = "redact"
	PrivacyModePseudonymize PrivacyMode = "pseudonymize"
	PrivacyModeAnonymize    PrivacyMode = "anonymize"
)

type PrivacyPolicy struct {
	Mode              PrivacyMode `json:"mode"`
	LocalOnly         bool        `json:"localOnly"`
	FailClosed        bool        `json:"failClosed"`
	Reversible        bool        `json:"reversible,omitempty"`
	ScopeID           string      `json:"scopeId,omitempty"`
	EntityTypes       []string    `json:"entityTypes,omitempty"`
	CustomTerms       []string    `json:"customTerms,omitempty"`
	AllowList         []string    `json:"allowList,omitempty"`
	MinimumConfidence float64     `json:"minimumConfidence,omitempty"`
	MaxResidualRisk   float64     `json:"maxResidualRisk,omitempty"`
	PreserveLength    bool        `json:"preserveLength,omitempty"`
	VaultKeyID        string      `json:"vaultKeyId,omitempty"`
}

type PrivacyFinding struct {
	Type        string  `json:"type"`
	Start       int     `json:"start"`
	End         int     `json:"end"`
	Confidence  float64 `json:"confidence"`
	Detector    string  `json:"detector"`
	Strategy    string  `json:"strategy"`
	Replacement string  `json:"replacement,omitempty"`
}

type PrivacySignature struct {
	Algorithm string `json:"algorithm"`
	KeyID     string `json:"keyId"`
	Value     string `json:"value"`
}

type PrivacyReceipt struct {
	ReceiptID        string            `json:"receiptId"`
	DocumentID       string            `json:"documentId"`
	Mode             PrivacyMode       `json:"mode"`
	PolicyDigest     string            `json:"policyDigest"`
	InputDigest      string            `json:"inputDigest"`
	OutputDigest     string            `json:"outputDigest"`
	Processor        string            `json:"processor"`
	ProcessorVersion string            `json:"processorVersion"`
	LocalOnly        bool              `json:"localOnly"`
	FailClosed       bool              `json:"failClosed"`
	Reversible       bool              `json:"reversible"`
	Detected         int               `json:"detected"`
	Transformed      int               `json:"transformed"`
	Unresolved       int               `json:"unresolved"`
	EntityCounts     map[string]int    `json:"entityCounts,omitempty"`
	ResidualRisk     float64           `json:"residualRisk"`
	Passed           bool              `json:"passed"`
	Warnings         []string          `json:"warnings,omitempty"`
	CreatedAt        time.Time         `json:"createdAt"`
	Signature        *PrivacySignature `json:"signature,omitempty"`
}

type PrivacyMetadata struct {
	Status          string         `json:"status"`
	Mode            PrivacyMode    `json:"mode"`
	ReceiptID       string         `json:"receiptId"`
	SourceDigest    string         `json:"sourceDigest"`
	SanitizedDigest string         `json:"sanitizedDigest"`
	ResidualRisk    float64        `json:"residualRisk"`
	EntityCounts    map[string]int `json:"entityCounts,omitempty"`
	Processor       string         `json:"processor"`
}

type SealedPrivacyVault struct {
	Algorithm  string `json:"algorithm"`
	KeyID      string `json:"keyId,omitempty"`
	Nonce      string `json:"nonce"`
	Ciphertext string `json:"ciphertext"`
}

type SanitizeRequest struct {
	DocumentID  string        `json:"documentId"`
	ContentType string        `json:"contentType"`
	Content     string        `json:"content"`
	Policy      PrivacyPolicy `json:"policy"`
}

type SanitizeResponse struct {
	Content string              `json:"content"`
	Receipt PrivacyReceipt      `json:"receipt"`
	Vault   *SealedPrivacyVault `json:"vault,omitempty"`
}
