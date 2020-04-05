package pkcs12

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"software.sslmate.com/src/go-pkcs12"
)

func KeyGen() error {
	baseDir := os.TempDir()
	priPath := path.Join(baseDir, "pri.pfx")
	pubPath := path.Join(baseDir, "pub.pem")
	keyBytes, err := rsa.GenerateKey(rand.Reader, 1024)

	if err != nil {
		return err
	}

	if err := keyBytes.Validate(); err != nil {
		return err
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
		return err
	}

	cert, err := x509.ParseCertificate(derBytes)

	if err != nil {
		return err
	}

	pfxBytes, err := pkcs12.Encode(rand.Reader, keyBytes, cert, []*x509.Certificate{}, pkcs12.DefaultPassword)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(
		priPath,
		pfxBytes,
		os.ModePerm,
	); err != nil {
		return err
	}
	if _, _, err := pkcs12.Decode(pfxBytes, pkcs12.DefaultPassword); err != nil {
		return err
	}
	certOut, err := os.Create(pubPath)

	if err != nil {
		return err
	}

	if err := pem.Encode(certOut, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}); err != nil {
		return err
	}

	if err := certOut.Close(); err != nil {
		return err
	}

	return nil
}
