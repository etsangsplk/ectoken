package ectoken

import (
	"crypto/cipher"
	"errors"
	"fmt"

	"golang.org/x/crypto/blowfish"
)

const (
	maxTokenChars = 512
)

// Generate returns a secure token for use with a request to an Edgecast cache
// host.
//
// The secret is used to encrypt the params string. The params string must not
// be longer than 512 chars.
func Generate(secret, params string) (string, error) {
	if len(params) >= maxTokenChars {
		return "", errors.New("Input too long")
	}

	plaintext := []byte(fmt.Sprintf("ec_secure=%03d&%s", len(params)+14, params))
	block, err := blowfish.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	ivec := make([]byte, blowfish.BlockSize)
	stream := cipher.NewCFBEncrypter(block, ivec)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return fmt.Sprintf("%02x", ciphertext), nil
}
