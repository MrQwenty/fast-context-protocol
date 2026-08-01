package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/MrQwenty/fast-context-protocol/internal/protocol"
)

type Catalogue struct {
	Nodes []protocol.ContextNode `json:"nodes"`
	byID  map[string]protocol.ContextNode
}

func LoadCatalogue(path string) (*Catalogue, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read catalogue: %w", err)
	}
	var c Catalogue
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("decode catalogue: %w", err)
	}
	c.byID = make(map[string]protocol.ContextNode, len(c.Nodes))
	for i := range c.Nodes {
		n := &c.Nodes[i]
		if n.Content == "" {
			return nil, fmt.Errorf("node %d has empty content", i)
		}
		sum := sha256.Sum256([]byte(n.Content))
		computed := "sha256:" + hex.EncodeToString(sum[:])
		if n.ID == "" || strings.HasPrefix(n.ID, "auto:") {
			n.ID = computed
		} else if n.ID != computed {
			return nil, fmt.Errorf("node %q digest mismatch: expected %s", n.ID, computed)
		}
		if n.ByteLength == 0 {
			n.ByteLength = int64(len([]byte(n.Content)))
		}
		if n.TokenEstimate == 0 {
			n.TokenEstimate = estimateTokens(n.Content)
		}
		c.byID[n.ID] = *n
	}
	return &c, nil
}

func (c *Catalogue) Get(id string) (protocol.ContextNode, bool) {
	n, ok := c.byID[id]
	return n, ok
}

func estimateTokens(content string) int {
	// Portable baseline estimate. Production providers should negotiate a tokenizer profile.
	words := len(strings.Fields(content))
	if words == 0 {
		return 1
	}
	return (words*4 + 2) / 3
}

// NewCatalogueForTesting builds an in-memory catalogue for conformance and unit tests.
func NewCatalogueForTesting(nodes []protocol.ContextNode) *Catalogue {
	byID := make(map[string]protocol.ContextNode, len(nodes))
	for _, node := range nodes {
		byID[node.ID] = node
	}
	return &Catalogue{Nodes: nodes, byID: byID}
}
