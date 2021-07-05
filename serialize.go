package jwsutil

import (
	"encoding/json"
	"fmt"
)

// SerializeCompact serialize the signature in JWS Compact Serialization
// See https://www.rfc-editor.org/rfc/rfc7515#section-7.1
func (s CompleteSignature) SerializeCompact() string {
	return fmt.Sprintf("%s.%s.%s", s.Protected, s.Payload, s.Signature)
}

// SerializeJSON serialize the signature in JWS JSON Serialization
// See https://www.rfc-editor.org/rfc/rfc7515#section-7.2
func (s CompleteSignature) SerializeJSON() string {
	return s.SerializeFlattenedJSON()
}

// SerializeGeneralJSON serialize the signature in General JWS JSON Serialization
// See https://www.rfc-editor.org/rfc/rfc7515#section-7.2.1
func (s CompleteSignature) SerializeGeneralJSON() string {
	return s.Enclose().Serialize()
}

// SerializeFlattenedJSON serialize the signature in Flattened JWS JSON Serialization
// See https://www.rfc-editor.org/rfc/rfc7515#section-7.2.2
func (s CompleteSignature) SerializeFlattenedJSON() string {
	serialized, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(serialized)
}

// Serialize serialize the envelope in General JWS JSON Serialization
// See https://www.rfc-editor.org/rfc/rfc7515#section-7.2.1
func (e Envelope) Serialize() string {
	serialized, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(serialized)
}

// Flattenable checks if an envelope can be flattenned.
func (e Envelope) Flattenable() bool {
	return len(e.Signatures) < 2
}
