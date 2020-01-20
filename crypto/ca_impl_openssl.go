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
	"path"
	"time"
)

//  Make a CA certificate.
//  Returns the certificate, private key and public key.
func MakeCACert(tls_dir string) {
	log.Println("Create certificates")
	get_bits := config.GetConfig("ca", "cert_bits")
	cert_bits, _ := strconv.Atoi(get_bits)
	cert_ca_name := config.GetConfig("ca", "cert_ca_name")
	MakeRequest(tls_dir, cert_bits, cert_ca_name)
}

func MakeRequest(tls_dir string, bits int, cn string) {
	bitSize := bits
	reader := rand.Reader
	// Create a x509 certificate object
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Organization: []string{config.GetConfig("ca", "cert_organization")},
			Country:      []string{config.GetConfig("ca", "cert_country")},
			Province:     []string{config.GetConfig("ca", "cert_state")},
			Locality:     []string{config.GetConfig("ca", "cert_locality")},

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
	ca_cert := path.Join(tls_dir, "/ca.pem")
	certOut, err := os.Create(ca_cert)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: ca_b})
	certOut.Close()
	log.Print("Created cert.pem\n")

	// Private key
	priv_key := path.Join(tls_dir, "/ca.key")
	keyOut, err := os.OpenFile(priv_key, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Print("Created key.pem\n")
}
