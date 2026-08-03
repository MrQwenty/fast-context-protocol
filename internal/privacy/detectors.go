package privacy

import (
	"net"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

type Candidate struct {
	Type       string
	Start      int
	End        int
	Confidence float64
	Detector   string
}

type Detector interface {
	Name() string
	Detect(text string) []Candidate
}

type patternRule struct {
	entity     string
	expression *regexp.Regexp
	confidence float64
	validate   func(string) bool
}

type PatternDetector struct {
	rules []patternRule
}

func NewPatternDetector() *PatternDetector {
	return &PatternDetector{rules: []patternRule{
		{entity: "EMAIL", expression: regexp.MustCompile(`(?i)\b[A-Z0-9._%+\-]+@[A-Z0-9.\-]+\.[A-Z]{2,}\b`), confidence: 0.99},
		{entity: "IP_ADDRESS", expression: regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`), confidence: 0.98, validate: validIP},
		{entity: "IBAN", expression: regexp.MustCompile(`(?i)\b[A-Z]{2}\d{2}(?:[ ]?[A-Z0-9]){11,30}\b`), confidence: 0.99, validate: validIBAN},
		{entity: "PAYMENT_CARD", expression: regexp.MustCompile(`\b(?:\d[ -]*?){13,19}\b`), confidence: 0.99, validate: validLuhn},
		{entity: "ITALIAN_TAX_CODE", expression: regexp.MustCompile(`(?i)\b[A-Z]{6}[0-9LMNPQRSTUV]{2}[ABCDEHLMPRST][0-9LMNPQRSTUV]{2}[A-Z][0-9LMNPQRSTUV]{3}[A-Z]\b`), confidence: 0.98},
		{entity: "US_SSN", expression: regexp.MustCompile(`\b\d{3}-\d{2}-\d{4}\b`), confidence: 0.98},
		{entity: "JWT", expression: regexp.MustCompile(`\beyJ[A-Za-z0-9_-]{8,}\.[A-Za-z0-9_-]{8,}\.[A-Za-z0-9_-]{8,}\b`), confidence: 0.995},
		{entity: "API_KEY", expression: regexp.MustCompile(`\b(?:sk-[A-Za-z0-9_-]{16,}|gh[pousr]_[A-Za-z0-9]{20,}|AKIA[0-9A-Z]{16})\b`), confidence: 0.995},
		{entity: "PHONE", expression: regexp.MustCompile(`(?:\+?\d{1,3}[ .-]?)?(?:\(?\d{2,4}\)?[ .-]?)?\d(?:[ .-]?\d){6,12}`), confidence: 0.88, validate: validPhone},
	}}
}

func (d *PatternDetector) Name() string { return "structured-patterns" }

func (d *PatternDetector) Detect(text string) []Candidate {
	var out []Candidate
	for _, rule := range d.rules {
		for _, index := range rule.expression.FindAllStringIndex(text, -1) {
			value := text[index[0]:index[1]]
			if rule.validate != nil && !rule.validate(value) {
				continue
			}
			out = append(out, Candidate{Type: rule.entity, Start: index[0], End: index[1], Confidence: rule.confidence, Detector: d.Name()})
		}
	}
	return out
}

type ContextDetector struct {
	rules []contextRule
}

type contextRule struct {
	entity     string
	expression *regexp.Regexp
	group      int
	confidence float64
}

func NewContextDetector() *ContextDetector {
	return &ContextDetector{rules: []contextRule{
		{entity: "PERSON", expression: regexp.MustCompile(`(?im)\b(?:full[ -]?name|name|nome(?: e cognome)?|patient|paziente|cliente|employee|dipendente)\s*[:=]\s*([\p{L}][\p{L}'’-]+(?:[ \t]+[\p{L}][\p{L}'’-]+){1,4})`), group: 1, confidence: 0.93},
		{entity: "ADDRESS", expression: regexp.MustCompile(`(?im)\b(?:address|indirizzo|residenza|domicilio)\s*[:=]\s*([^\n;]{5,160})`), group: 1, confidence: 0.91},
		{entity: "DATE_OF_BIRTH", expression: regexp.MustCompile(`(?im)\b(?:date of birth|birth date|dob|data di nascita)\s*[:=]\s*([0-3]?\d[./-][01]?\d[./-](?:19|20)\d{2}|(?:19|20)\d{2}-[01]\d-[0-3]\d)`), group: 1, confidence: 0.97},
		{entity: "ORGANIZATION", expression: regexp.MustCompile(`(?im)\b(?:company|azienda|employer|datore di lavoro|societ[aà])\s*[:=]\s*([^\n;,]{2,120})`), group: 1, confidence: 0.88},
		{entity: "ACCOUNT_ID", expression: regexp.MustCompile(`(?im)\b(?:account|customer|client|patient|record|pratica)[ _-]?(?:id|number|no\.?|n\.)\s*[:=]\s*([A-Z0-9][A-Z0-9_.\-/]{3,64})`), group: 1, confidence: 0.94},
	}}
}

func (d *ContextDetector) Name() string { return "label-context" }

func (d *ContextDetector) Detect(text string) []Candidate {
	var out []Candidate
	for _, rule := range d.rules {
		matches := rule.expression.FindAllStringSubmatchIndex(text, -1)
		for _, match := range matches {
			idx := rule.group * 2
			if idx+1 >= len(match) || match[idx] < 0 {
				continue
			}
			start, end := trimRange(text, match[idx], match[idx+1])
			if end > start {
				value := strings.TrimSpace(text[start:end])
				if isPrivacyPlaceholder(value) {
					continue
				}
				out = append(out, Candidate{Type: rule.entity, Start: start, End: end, Confidence: rule.confidence, Detector: d.Name()})
			}
		}
	}
	return out
}

type DictionaryDetector struct {
	terms []string
}

func NewDictionaryDetector(terms []string) *DictionaryDetector {
	cleaned := make([]string, 0, len(terms))
	seen := map[string]struct{}{}
	for _, term := range terms {
		term = strings.TrimSpace(term)
		if term == "" {
			continue
		}
		key := strings.ToLower(term)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		cleaned = append(cleaned, term)
	}
	sort.Slice(cleaned, func(i, j int) bool { return len(cleaned[i]) > len(cleaned[j]) })
	return &DictionaryDetector{terms: cleaned}
}

func (d *DictionaryDetector) Name() string { return "custom-dictionary" }

func (d *DictionaryDetector) Detect(text string) []Candidate {
	lower := strings.ToLower(text)
	var out []Candidate
	for _, term := range d.terms {
		needle := strings.ToLower(term)
		for offset := 0; offset < len(lower); {
			index := strings.Index(lower[offset:], needle)
			if index < 0 {
				break
			}
			start := offset + index
			end := start + len(needle)
			if wordBoundary(text, start, end) {
				out = append(out, Candidate{Type: "CUSTOM_IDENTIFIER", Start: start, End: end, Confidence: 1, Detector: d.Name()})
			}
			offset = end
		}
	}
	return out
}

func validIP(value string) bool {
	parsed := net.ParseIP(value)
	return parsed != nil && parsed.To4() != nil
}

func validPhone(value string) bool {
	digits := onlyDigits(value)
	return len(digits) >= 7 && len(digits) <= 15
}

func validLuhn(value string) bool {
	digits := onlyDigits(value)
	if len(digits) < 13 || len(digits) > 19 {
		return false
	}
	sum := 0
	parity := len(digits) % 2
	for i, r := range digits {
		d := int(r - '0')
		if i%2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return sum%10 == 0
}

func validIBAN(value string) bool {
	value = strings.ToUpper(strings.ReplaceAll(value, " ", ""))
	if len(value) < 15 || len(value) > 34 {
		return false
	}
	rearranged := value[4:] + value[:4]
	remainder := 0
	for _, r := range rearranged {
		if r >= '0' && r <= '9' {
			remainder = (remainder*10 + int(r-'0')) % 97
			continue
		}
		if r < 'A' || r > 'Z' {
			return false
		}
		n := int(r-'A') + 10
		remainder = (remainder*100 + n) % 97
	}
	return remainder == 1
}

func onlyDigits(value string) string {
	var b strings.Builder
	for _, r := range value {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func trimRange(text string, start, end int) (int, int) {
	for start < end {
		r, size := runeAt(text[start:end])
		if !unicode.IsSpace(r) {
			break
		}
		start += size
	}
	for start < end {
		r, size := lastRune(text[start:end])
		if !unicode.IsSpace(r) && r != ',' && r != '.' {
			break
		}
		end -= size
	}
	return start, end
}

func runeAt(value string) (rune, int) {
	for _, r := range value {
		return r, len(string(r))
	}
	return 0, 0
}

func lastRune(value string) (rune, int) {
	var last rune
	for _, r := range value {
		last = r
	}
	return last, len(string(last))
}

func wordBoundary(text string, start, end int) bool {
	leftOK := start == 0 || !isWordByte(text[start-1])
	rightOK := end == len(text) || !isWordByte(text[end])
	return leftOK && rightOK
}

func isWordByte(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_'
}

func isPrivacyPlaceholder(value string) bool {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "[REDACTED:") && strings.HasSuffix(value, "]") {
		return true
	}
	if strings.HasPrefix(value, "<") && strings.HasSuffix(value, ">") {
		inner := value[1 : len(value)-1]
		for _, r := range inner {
			if (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == ':' {
				continue
			}
			return false
		}
		return inner != ""
	}
	return false
}
