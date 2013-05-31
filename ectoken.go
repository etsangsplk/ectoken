package ectoken

import (
	"code.google.com/p/go.crypto/blowfish"
	"crypto/cipher"
	"errors"
	"fmt"
)

const (
	MAX_TOKEN_CHARS = 512
)

func Generate(secret, params string) (string, error) {
	if len(params) >= MAX_TOKEN_CHARS {
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
