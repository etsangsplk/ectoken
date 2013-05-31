package ectoken

import (
	"strings"
	"testing"
)

func TestTokenNoParams(t *testing.T) {
	token, _ := Generate("secret", "test")

	if token != "646952e22c167f40ae84bb4b2558d9fb87c8" {
		t.Errorf("Invalid token generated.")
	}
}

func TestTokenTwoParams(t *testing.T) {
	token, _ := Generate("secret", "ec_expire=1333238400&ec_url_allow=test.mp3")

	if token != "646952e22c167f40ae84bb4f2b58c8fd1e24a7cd566ac603e92e68dbbae3fa5d82cc6ab1d13f0e88477d70bf6e19cf8894e1c28adfba0b5f" {
		t.Errorf("Invalid token generated.")
	}
}

func TestErrorInputMaxCharsExceeded(t *testing.T) {
	_, err := Generate("secret", strings.Repeat("a", 511))
	if err != nil {
		t.Errorf("Error returned for input of length 511")
	}

	_, err = Generate("secret", strings.Repeat("a", 512))
	if err == nil {
		t.Errorf("Error not returned for input of length 512")
	}
}
