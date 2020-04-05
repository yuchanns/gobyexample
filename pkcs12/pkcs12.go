package pkcs12

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path"
)

func KeyGen() error {
	keyBytes, err := rsa.GenerateKey(rand.Reader, 1024)

	if err != nil {
		return err
	}

	if err := keyBytes.Validate(); err != nil {
		return err
	}

	dn := pkix.Name{
		Country:            []string{"EN"},
		Organization:       []string{"org"},
		OrganizationalUnit: []string{"org"},
		Locality:           []string{"city"},
		Province:           []string{"province"},
		CommonName:         "name",
	}
	//asn1Dn, err := asn1.Marshal(dn)
	//if err != nil {
	//	return err
	//}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)

	if err != nil {
		return err
	}
	template := x509.Certificate{
		Subject:            dn,
		SignatureAlgorithm: x509.SHA256WithRSA,
		SerialNumber:       serialNumber,
	}

	//pfxBytes, err := pkcs12.Encode(rand.Reader, keyBytes, &template, []*x509.Certificate{}, pkcs12.DefaultPassword)
	//
	//if err != nil {
	//	return err
	//}
	//
	//if err := ioutil.WriteFile(
	//	"/Users/yuchanns/Coding/golang/gobyexample/pkcs12/pri.pfx",
	//	pfxBytes,
	//	os.ModePerm,
	//); err != nil {
	//	return err
	//}
	//
	//_, _, err = pkcs12.Decode(pfxBytes, pkcs12.DefaultPassword)
	//
	//return err

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &keyBytes.PublicKey, keyBytes)

	if err != nil {
		return err
	}

	savePath := path.Join(os.TempDir(), "/pub.pem")
	certOut, err := os.Create(savePath)

	fmt.Println(os.TempDir())

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
