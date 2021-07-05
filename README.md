# go-jwsutil
Golang utility package for [JSON Web Signature (JWS)](https://www.rfc-editor.org/rfc/rfc7515).

## Example

```go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/shizhMSFT/go-jwsutil"
)

func main() {
	// Generate a RSA key pair for this demo
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	panicOnError(err)

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodPS512, jwt.MapClaims{
		"sub": "demo",
	})
	serialized, err := token.SignedString(key)
	panicOnError(err)

	// Convert compact serialization to JSON serialization
	serialized, err = jwsutil.ConvertCompactToJSON(serialized, jwt.MapClaims{
		"foo": "bar",
	})
	panicOnError(err)

	// Print the JSON serialized token
	fmt.Println(serialized)

	// Convert it back to compact
	var unprotectedClaims jwt.MapClaims
	serialized, err = jwsutil.ConvertJSONToCompact(serialized, &unprotectedClaims)
	panicOnError(err)

	// Print out the extracted unprotected claims
	fmt.Println(unprotectedClaims)

	// Parse and verify the converted token
	token, err = jwt.Parse(serialized, func(token *jwt.Token) (interface{}, error) {
		if alg := token.Method.Alg(); alg != jwt.SigningMethodPS512.Alg() {
			return nil, fmt.Errorf("Unexpected signing method: %v", alg)
		}
		return &key.PublicKey, nil
	})
	panicOnError(err)
	fmt.Println(token.Valid)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
```

The above code outputs:

```
{"payload":"eyJzdWIiOiJkZW1vIn0","protected":"eyJhbGciOiJQUzUxMiIsInR5cCI6IkpXVCJ9","header":{"foo":"bar"},"signature":"s9tGI6169wK1BJyUrZvAN-1PA1IK_sxADMmMI-tAnnkRFdM_gAscBhSWmRY7dJGAjhkuK6itQC_NUWnPYp9GD7YNSig8dcdBhCIxYhfDUDbaGEz8SDVijuJ_oZpBySGBF9Y_01v5ESHd_x8j70kZcsf5JjYah1D5DHz76D8atLbf4nn84koy6-Tc6wbBpSZLyj0A-rdNcPGk_iMBxFbhSAmsIMZEUc6frJpPwp-5uoUnrHuPwWlOpo1gQox0t8x3Wkz6ebi2RdWhJW-s_kfV72DExzNT_aDTNxX5OtyfQ7QSMdc-wBgHU1l_fvsLSylE26dey_YhOBT9jAywnF7n3g"}
map[foo:bar]
true
```

The flattened JWS JSON object in the output can be pretty printed as

```json
{
    "payload": "eyJzdWIiOiJkZW1vIn0",
    "protected": "eyJhbGciOiJQUzUxMiIsInR5cCI6IkpXVCJ9",
    "header": {
        "foo": "bar"
    },
    "signature": "s9tGI6169wK1BJyUrZvAN-1PA1IK_sxADMmMI-tAnnkRFdM_gAscBhSWmRY7dJGAjhkuK6itQC_NUWnPYp9GD7YNSig8dcdBhCIxYhfDUDbaGEz8SDVijuJ_oZpBySGBF9Y_01v5ESHd_x8j70kZcsf5JjYah1D5DHz76D8atLbf4nn84koy6-Tc6wbBpSZLyj0A-rdNcPGk_iMBxFbhSAmsIMZEUc6frJpPwp-5uoUnrHuPwWlOpo1gQox0t8x3Wkz6ebi2RdWhJW-s_kfV72DExzNT_aDTNxX5OtyfQ7QSMdc-wBgHU1l_fvsLSylE26dey_YhOBT9jAywnF7n3g"
}
```