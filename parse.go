package jwsutil

import (
	"encoding/json"
	"strings"
)

// Parse parses the serialized JWS smartly.
func Parse(serialized string) (Envelope, error) {
	if strings.HasPrefix(serialized, "{") {
		return ParseJSON(serialized)
	}
	sig, err := ParseCompact(serialized)
	if err != nil {
		return Envelope{}, err
	}
	return Envelope{
		Payload: sig.Payload,
		Signatures: []Signature{
			sig.Signature,
		},
	}, nil
}

// Parse parses the compact serialized JWS.
// See https://www.rfc-editor.org/rfc/rfc7515#section-7.1
func ParseCompact(serialized string) (CompleteSignature, error) {
	parts := strings.Split(serialized, ".")
	if len(parts) != 3 {
		return CompleteSignature{}, ErrInvalidCompactSerialization
	}
	return CompleteSignature{
		Payload: parts[1],
		Signature: Signature{
			Protected: parts[0],
			Signature: parts[2],
		},
	}, nil
}

// Parse parses the compact serialized JWS.
// See https://www.rfc-editor.org/rfc/rfc7515#section-7.2
func ParseJSON(serialized string) (Envelope, error) {
	var combined struct {
		Signature
		Envelope
	}
	if err := json.Unmarshal([]byte(serialized), &combined); err != nil {
		return Envelope{}, ErrInvalidJSONSerialization
	}
	if len(combined.Signatures) > 0 {
		return combined.Envelope, nil
	}
	return Envelope{
		Payload: combined.Payload,
		Signatures: []Signature{
			combined.Signature,
		},
	}, nil
}
