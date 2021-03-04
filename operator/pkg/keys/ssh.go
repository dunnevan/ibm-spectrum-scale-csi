package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

type sshKeyBuilder struct {
	bitSize int
}

// New builder that can generate private and public ssh keys
func NewSsh() sshKeyBuilder {
	return sshKeyBuilder{
		bitSize: 4096,
	}
}

// Generate private key in PEM format, and corresponding public key in OpenSSH format
func (kb sshKeyBuilder) Generate() (private []byte, public []byte, err error) {
	privKey, err := rsa.GenerateKey(rand.Reader, kb.bitSize)
	if err != nil {
		return nil, nil, err
	}

	err = privKey.Validate()
	if err != nil {
		return nil, nil, err
	}

	// Get ASN.1 DER format
	privDer := x509.MarshalPKCS1PrivateKey(privKey)

	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDer,
	}

	// Private key in PEM format
	privPem := pem.EncodeToMemory(&privBlock)

	pubKey, err := ssh.NewPublicKey(privKey.Public())
	if err != nil {
		return nil, nil, err
	}

	pubOpenSsh := ssh.MarshalAuthorizedKey(pubKey)

	return privPem, pubOpenSsh, nil
}
