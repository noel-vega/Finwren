package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const tokenByteLength = 32

func generateRawToken() (string, []byte, error) {
	b := make([]byte, tokenByteLength)
	if _, err := rand.Read(b); err != nil {
		return "", nil, err
	}

	raw := base64.RawURLEncoding.EncodeToString(b)
	h := sha256.Sum256([]byte(raw))
	return raw, h[:], nil
}
