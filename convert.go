package jwsutil

import "encoding/json"

// ConvertCompactToJSON converts compact serialized JWS to flattened JSON form, adding unprotected headers.
func ConvertCompactToJSON(serialized string, unprotected interface{}) (string, error) {
	sig, err := ParseCompact(serialized)
	if err != nil {
		return "", err
	}
	if unprotected != nil {
		unprotectedJSON, err := json.Marshal(unprotected)
		if err != nil {
			return "", err
		}
		sig.Unprotected = unprotectedJSON
	}
	return sig.SerializeFlattenedJSON(), nil
}

// ConvertJSONToCompact converts JSON serialized JWS to compact form, extracting unprotected headers.
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
