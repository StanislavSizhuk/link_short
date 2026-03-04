package util

import (
	"crypto/sha256"
	"encoding/base64"
)

func ShortenURL(fullUrl string) string {
	hash := sha256.Sum256([]byte(fullUrl))

	short := hash[:6]
	encoded := base64.RawURLEncoding.EncodeToString(short)

	if len(encoded) > 8 {
		encoded = encoded[:8]
	}
	return encoded
}
