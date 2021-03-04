package keys

import (
	"crypto/rand"
	"io"
	"math/big"
)

type passwordBuilder struct {
	passwordLen int
	alphabet    string
	rand        io.Reader
}

// New builder that can generate a random password
func NewPassword() passwordBuilder {
	return passwordBuilder{
		passwordLen: 20,
		alphabet:    "ABCEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_",
		rand:        rand.Reader,
	}
}

// Generate random password
func (pb passwordBuilder) Generate() ([]byte, error) {

	// Allocate a buffer to generate into
	passwordBuf := make([]byte, pb.passwordLen)
	alphabetLen := big.NewInt(int64(len(pb.alphabet)))

	for i := 0; i < pb.passwordLen; i++ {
		n, err := rand.Int(pb.rand, alphabetLen)
		if err != nil {
			return nil, err
		}
		passwordBuf[i] = pb.alphabet[n.Int64()]
	}

	return passwordBuf, nil
}
