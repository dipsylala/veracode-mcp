package hmac

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	algorithm     = "VERACODE-HMAC-SHA-256"
	requestVerStr = "vcode_request_version_1"

	dataFormat   = "id=%s&host=%s&url=%s&method=%s"
	headerFormat = "%s id=%s,ts=%s,nonce=%x,sig=%x" // lowercase hex
)

// CalculateAuthorizationHeader calculates the Veracode HMAC Authorization header value.
func CalculateAuthorizationHeader(u *url.URL, httpMethod, apiKeyID, apiKeySecret string) (string, error) {
	// ---- Input validation (fail fast, clearer errors)
	if u == nil {
		return "", errors.New("url is nil")
	}
	if apiKeyID == "" {
		return "", errors.New("apiKeyID is empty")
	}
	if apiKeySecret == "" {
		return "", errors.New("apiKeySecret is empty")
	}
	httpMethod = strings.TrimSpace(httpMethod)
	if httpMethod == "" {
		return "", errors.New("httpMethod is empty")
	}
	httpMethod = strings.ToUpper(httpMethod)

	// ---- Timestamp in milliseconds (string)
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())

	// ---- Nonce: 16 random bytes
	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("generate nonce: %w", err)
	}

	// ---- Secret is hex-encoded in Veracode examples: decode it
	secretBytes, err := hex.DecodeString(strings.TrimSpace(apiKeySecret))
	if err != nil {
		// Avoid echoing secret in error output; include only parsing failure context.
		return "", fmt.Errorf("apiKeySecret must be hex-encoded: %w", err)
	}

	// ---- Canonicalize the request URI (path + ?rawQuery)
	canonicalURI := canonicalRequestURI(u)

	// ---- Data string per spec
	data := fmt.Sprintf(dataFormat, apiKeyID, u.Hostname(), canonicalURI, httpMethod)

	// ---- Key derivation chain
	kNonce := hmacSHA256(secretBytes, nonce)
	kDate := hmacSHA256(kNonce, []byte(ts))
	kSig := hmacSHA256(kDate, []byte(requestVerStr))
	signature := hmacSHA256(kSig, []byte(data))

	// ---- Final Authorization header (lowercase hex for nonce/sig)
	return fmt.Sprintf(headerFormat, algorithm, apiKeyID, ts, nonce, signature), nil
}

func canonicalRequestURI(u *url.URL) string {
	// EscapedPath uses RawPath if set; otherwise it escapes Path.
	// Ensures we sign the encoded path representation.
	path := u.EscapedPath()
	if path == "" {
		path = "/"
	}
	if u.RawQuery != "" {
		return path + "?" + u.RawQuery
	}
	return path
}

func hmacSHA256(key, msg []byte) []byte {
	m := hmac.New(sha256.New, key)
	_, _ = m.Write(msg) // hash.Write never returns an error
	return m.Sum(nil)
}
