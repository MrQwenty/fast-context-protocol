package privacy

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

func sealVault(mapping map[string]string, keyID string, key []byte) (*protocol.SealedPrivacyVault, error) {
	if len(key) != 32 {
		return nil, errors.New("reversible pseudonymization requires a 32-byte AES-256 vault key")
	}
	plaintext, err := json.Marshal(mapping)
	if err != nil {
		return nil, fmt.Errorf("encode privacy vault: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("create vault cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("create vault AEAD: %w", err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("create vault nonce: %w", err)
	}
	ciphertext := gcm.Seal(nil, nonce, plaintext, []byte(keyID))
	return &protocol.SealedPrivacyVault{
		Algorithm:  "AES-256-GCM",
		KeyID:      keyID,
		Nonce:      base64.RawURLEncoding.EncodeToString(nonce),
		Ciphertext: base64.RawURLEncoding.EncodeToString(ciphertext),
	}, nil
}

func OpenVault(vault protocol.SealedPrivacyVault, key []byte) (map[string]string, error) {
	if vault.Algorithm != "AES-256-GCM" {
		return nil, fmt.Errorf("unsupported vault algorithm %q", vault.Algorithm)
	}
	if len(key) != 32 {
		return nil, errors.New("vault key must be 32 bytes")
	}
	nonce, err := base64.RawURLEncoding.DecodeString(vault.Nonce)
	if err != nil {
		return nil, fmt.Errorf("decode vault nonce: %w", err)
	}
	ciphertext, err := base64.RawURLEncoding.DecodeString(vault.Ciphertext)
	if err != nil {
		return nil, fmt.Errorf("decode vault ciphertext: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("create vault cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("create vault AEAD: %w", err)
	}
	plaintext, err := gcm.Open(nil, nonce, ciphertext, []byte(vault.KeyID))
	if err != nil {
		return nil, errors.New("privacy vault authentication failed")
	}
	var mapping map[string]string
	if err := json.Unmarshal(plaintext, &mapping); err != nil {
		return nil, fmt.Errorf("decode privacy vault: %w", err)
	}
	return mapping, nil
}
