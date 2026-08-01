package protocol

import "time"

const Version = "0.1"

type Discovery struct {
	Name              string            `json:"name"`
	Description       string            `json:"description,omitempty"`
	Versions          []string          `json:"versions"`
	Endpoints         map[string]string `json:"endpoints"`
	DeliveryModes     []DeliveryMode    `json:"deliveryModes"`
	Features          Features          `json:"features"`
	MaxRequestBytes   int64             `json:"maxRequestBytes,omitempty"`
	DefaultCacheTTLMS int64             `json:"defaultCacheTtlMs,omitempty"`
}

type Features struct {
	KnownContext  bool `json:"knownContext"`
	Fetch         bool `json:"fetch"`
	Receipts      bool `json:"receipts"`
	Deltas        bool `json:"deltas"`
	Subscriptions bool `json:"subscriptions"`
	Streaming     bool `json:"streaming"`
	Signatures    bool `json:"signatures"`
}

type Intent struct {
	Type   string            `json:"type"`
	Target string            `json:"target,omitempty"`
	Params map[string]string `json:"params,omitempty"`
}

type Consumer struct {
	ID            string   `json:"id,omitempty"`
	ModelFamily   string   `json:"modelFamily,omitempty"`
	ContextWindow int      `json:"contextWindow,omitempty"`
	Modalities    []string `json:"modalities,omitempty"`
}

type Budget struct {
	MaxTokens    int   `json:"maxTokens"`
	MaxLatencyMS int64 `json:"maxLatencyMs,omitempty"`
	MaxBytes     int64 `json:"maxBytes"`
}

type Requirements struct {
	FreshnessMS       int64    `json:"freshnessMs,omitempty"`
	MinimumConfidence float64  `json:"minimumConfidence,omitempty"`
	IncludeProvenance bool     `json:"includeProvenance,omitempty"`
	AcceptSensitivity []string `json:"acceptSensitivity,omitempty"`
}

type ContextRequest struct {
	RequestID    string       `json:"requestId"`
	Intent       Intent       `json:"intent"`
	Consumer     Consumer     `json:"consumer,omitempty"`
	Budget       Budget       `json:"budget"`
	Requirements Requirements `json:"requirements,omitempty"`
	KnownContext []string     `json:"knownContext,omitempty"`
}

type Provenance struct {
	Provider  string `json:"provider"`
	Resource  string `json:"resource"`
	Author    string `json:"author,omitempty"`
	Retrieved string `json:"retrievedAt,omitempty"`
}

type Relationship struct {
	Type   string `json:"type"`
	Target string `json:"target"`
}

type ContextNode struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Version       string         `json:"version,omitempty"`
	ContentType   string         `json:"contentType"`
	Content       string         `json:"content,omitempty"`
	TokenEstimate int            `json:"tokenEstimate"`
	ByteLength    int64          `json:"byteLength"`
	Priority      float64        `json:"priority"`
	Confidence    float64        `json:"confidence"`
	Sensitivity   string         `json:"sensitivity,omitempty"`
	CreatedAt     time.Time      `json:"createdAt"`
	FreshUntil    *time.Time     `json:"freshUntil,omitempty"`
	Provenance    Provenance     `json:"provenance"`
	Relationships []Relationship `json:"relationships,omitempty"`
	Intents       []string       `json:"intents,omitempty"`
	Targets       []string       `json:"targets,omitempty"`
}

type DeliveryMode string

const (
	DeliveryInline    DeliveryMode = "inline"
	DeliveryReference DeliveryMode = "reference"
	DeliveryFetch     DeliveryMode = "fetch"
)

type PlannedChunk struct {
	Node      ContextNode  `json:"node"`
	Delivery  DeliveryMode `json:"delivery"`
	Rank      int          `json:"rank"`
	Score     float64      `json:"score"`
	Reason    string       `json:"reason"`
	FetchPath string       `json:"fetchPath,omitempty"`
}

type Omission struct {
	NodeID string `json:"nodeId"`
	Reason string `json:"reason"`
}

type ContextPlan struct {
	RequestID          string         `json:"requestId"`
	PlanID             string         `json:"planId"`
	ProtocolVersion    string         `json:"protocolVersion"`
	ContextRoot        string         `json:"contextRoot"`
	Complete           bool           `json:"complete"`
	EstimatedTokens    int            `json:"estimatedTokens"`
	EstimatedBytes     int64          `json:"estimatedBytes"`
	EstimatedLatencyMS int64          `json:"estimatedLatencyMs"`
	Chunks             []PlannedChunk `json:"chunks"`
	Omissions          []Omission     `json:"omissions,omitempty"`
	CreatedAt          time.Time      `json:"createdAt"`
}

type Receipt struct {
	RequestID          string   `json:"requestId"`
	PlanID             string   `json:"planId"`
	Delivered          []string `json:"delivered,omitempty"`
	Reused             []string `json:"reused,omitempty"`
	Rejected           []string `json:"rejected,omitempty"`
	ActualTokens       int      `json:"actualTokens,omitempty"`
	TimeToFirstMS      int64    `json:"timeToFirstMs,omitempty"`
	TotalDeliveryMS    int64    `json:"totalDeliveryMs,omitempty"`
	Outcome            string   `json:"outcome,omitempty"`
	ConsumerRecordedAt string   `json:"consumerRecordedAt,omitempty"`
}
