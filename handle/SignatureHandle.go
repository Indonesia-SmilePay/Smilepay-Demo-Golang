package handle

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"strings"
)

func Sha256RshSignature(stringToSign string, privateKeyString string) (string, bool) {
	// get from step1;  Base64-encoded private key

	// Decode the base64 string to obtain the PEM encoded bytes
	pemBytes, err := base64.StdEncoding.DecodeString(privateKeyString)
	if err != nil {
		fmt.Println("Error decoding private key:", err)
		return "", true
	}

	// Parse the DER encoded bytes to obtain the private key
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return "", true
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("Invalid private key type")
		return "", true
	}

	// Data to be signed
	message := []byte(stringToSign)

	// Compute the SHA-256 hash of the data
	hashed := sha256.Sum256(message)

	// Sign the hashed message using the private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("Error signing the message:", err)
		return "", true
	}

	// Convert the signature to base64 string
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	// Display the base64-encoded signature
	fmt.Println("Base64-encoded Signature:")
	fmt.Println(signatureBase64)
	return signatureBase64, false
}

// SHA256 computes SHA-256 hash of the input string
func SHA256(requestBody string) ([]byte, error) {
	hash := sha256.New()
	hash.Write([]byte(requestBody))
	return hash.Sum(nil), nil
}

// byte2Hex converts bytes to hexadecimal string
func byte2Hex(bytes []byte) string {
	var hexString strings.Builder
	for _, b := range bytes {
		fmt.Fprintf(&hexString, "%02x", b)
	}
	return hexString.String()
}

// hmacSHA512 computes HMAC-SHA512 hash of the input string
func hmacSHA512(signData, secret string) (string, error) {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(signData))
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

// minify removes whitespace and comments from JSON string
func minify(jsonString string) string {
	var inString, inMultiLineComment, inSingleLineComment bool
	var stringOpener rune
	var out strings.Builder

	for i := 0; i < len(jsonString); i++ {
		c := rune(jsonString[i])
		var cc string
		if i+1 < len(jsonString) {
			cc = jsonString[i : i+2]
		} else {
			cc = jsonString[i:]
		}

		if inString {
			if c == stringOpener {
				inString = false
				out.WriteRune(c)
			} else if c == '\\' {
				if i+1 < len(jsonString) {
					out.WriteString(jsonString[i : i+2])
					i++
				}
			} else {
				out.WriteRune(c)
			}
		} else if inSingleLineComment {
			if c == '\r' || c == '\n' {
				inSingleLineComment = false
			}
		} else if inMultiLineComment {
			if cc == "*/" {
				inMultiLineComment = false
				i++
			}
		} else {
			if cc == "/*" {
				inMultiLineComment = true
				i++
			} else if cc == "//" {
				inSingleLineComment = true
				i++
			} else if c == '"' || c == '\'' {
				inString = true
				stringOpener = c
				out.WriteRune(c)
			} else if !isWhitespace(c) {
				out.WriteRune(c)
			}
		}
	}
	return out.String()
}

// isWhitespace checks if a rune is a whitespace character
func isWhitespace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}
