package jwsutil

import "encoding/json"

// ConvertCompactToJSON converts compact serialized JWS to flattened JSON form.
func ConvertCompactToJSON(serialized string) (string, error) {
	sig, err := ParseCompact(serialized)
	if err != nil {
		return "", err
	}
	return sig.SerializeFlattenedJSON(), nil
}

// ConvertJSONToCompact converts JSON serialized JWS to compact form.
func ConvertJSONToCompact(serialized string, unprotected interface{}) (string, error) {
	envelope, err := ParseJSON(serialized)
	if err != nil {
		return "", err
	}

	sig := envelope.CompleteSignature()
	if unprotected != nil {
		if err := json.Unmarshal(sig.Unprotected, unprotected); err != nil {
			return "", err
		}
	}
	return sig.SerializeCompact(), nil
}
