package jwsutil

import "encoding/json"

// Signature represents a detached signature.
type Signature struct {
	Protected   string          `json:"protected,omitempty"`
	Unprotected json.RawMessage `json:"header,omitempty"`
	Signature   string          `json:"signature,omitempty"`
}

// CompleteSignature represents a clear signed signature.
type CompleteSignature struct {
	Payload string `json:"payload,omitempty"`
	Signature
}

// Enclose packs the signature into an envelope.
func (s CompleteSignature) Enclose() Envelope {
	return Envelope{
		Payload: s.Payload,
		Signatures: []Signature{
			s.Signature,
		},
	}
}

// Envelope contains a common payload signed by multiple signatures.
type Envelope struct {
	Payload    string      `json:"payload,omitempty"`
	Signatures []Signature `json:"signatures,omitempty"`
}

// Signature returns the first or default signature.
func (e Envelope) Signature() Signature {
	if len(e.Signatures) == 0 {
		return Signature{}
	}
	return e.Signatures[0]
}

// CompleteSignature exports the first or default complete signature.
func (e Envelope) CompleteSignature() CompleteSignature {
	return CompleteSignature{
		Payload:   e.Payload,
		Signature: e.Signature(),
	}
}

// CompleteSignatures exports complete signatures.
func (e Envelope) CompleteSignatures() []CompleteSignature {
	signatures := make([]CompleteSignature, 0, len(e.Signatures))
	for _, sig := range e.Signatures {
		signatures = append(signatures, CompleteSignature{
			Payload:   e.Payload,
			Signature: sig,
		})
	}
	return signatures
}
