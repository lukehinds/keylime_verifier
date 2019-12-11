package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"keylime_verifier/config"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

//  Make a CA certificate.
//  Returns the certificate, private key and public key.
func MakeCACert(tls_dir) {
	get_bits := config.GetConfig("ca", "cert_bits")
	cert_bits, _ := strconv.Atoi(get_bits)
	cert_ca_name := config.GetConfig("ca", "cert_ca_name")
	MakeRequest(cert_bits, cert_ca_name)
}

func MakeRequest(bits int, cn string) {
	bitSize := bits
	reader := rand.Reader
	// Create a x509 certificate object
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Organization: []string{config.GetConfig("ca", "cert_organization")}, // config.get('ca','cert_organization')
			Country:      []string{config.GetConfig("ca", "cert_country")},
			Province:     []string{config.GetConfig("ca", "cert_state")},
			Locality:     []string{config.GetConfig("ca", "cert_locality")},
			// StreetAddress: []string{"ADDRESS"},
			// PostalCode:    []string{"POSTAL_CODE"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	// create private key
	priv, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		log.Fatal(err)
	}
	// derive public key from private key
	pub := &priv.PublicKey
	//  create x509 certificate using object data from `&x509.Certificate`
	ca_b, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Println("Create ca failed", err)
		return
	}

	// Public key
	certOut, err := os.Create("ca.crt")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: ca_b})
	certOut.Close()
	log.Print("Created cert.pem\n")

	// Private key
	keyOut, err := os.OpenFile("ca.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Print("Created key.pem\n")
}
