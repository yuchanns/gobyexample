package pkcs12

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/coreos/etcd/pkg/testutil"
	"math/big"
	"software.sslmate.com/src/go-pkcs12"
	"testing"
)

func TestKeyGen(t *testing.T) {
	testutil.AssertNil(t, KeyGen())
}

func TestSomething(t *testing.T) {
	keyBytes, err := rsa.GenerateKey(rand.Reader, 1024)

	if err != nil {
		panic(err)
	}

	if err := keyBytes.Validate(); err != nil {
		panic(err)
	}

	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:            []string{"EN"},
			Organization:       []string{"org"},
			OrganizationalUnit: []string{"org"},
			Locality:           []string{"city"},
			Province:           []string{"province"},
			CommonName:         "name",
		},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, &keyBytes.PublicKey, keyBytes)

	if err != nil {
		panic(err)
	}

	cert, err := x509.ParseCertificate(derBytes)

	if err != nil {
		panic(err)
	}

	pfxBytes, err := pkcs12.Encode(rand.Reader, keyBytes, cert, []*x509.Certificate{}, pkcs12.DefaultPassword)

	if err != nil {
		panic(err)
	}

	_, _, _, err = pkcs12.DecodeChain(pfxBytes, pkcs12.DefaultPassword)
	if err != nil {
		panic(err) // I got: pkcs12: expected exactly one certificate in the certBag
	}
}
