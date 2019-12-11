package app

import (
	"keylime_verifier/common"
	"keylime_verifier/config"
	"keylime_verifier/crypto"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func InitTLS() {
	enable_tls := config.GetConfig("general", "enable_tls")
	if !strings.Contains("True", enable_tls) {
		log.Printf("TLS is currently disabled, keys will be sent in the clear! Should only be used for testing.")
	} // needs to return None if not enabled

	log.Printf("Setting up TLS...")
	my_cert := config.GetConfig("cloud_verifier", "my_cert")
	ca_cert := config.GetConfig("cloud_verifier", "ca_cert")
	my_priv_key := config.GetConfig("cloud_verifier", "my_priv_key")
	my_key_pw := config.GetConfig("cloud_verifier", "my_key_pw")
	tls_dir := config.GetConfig("cloud_verifier", "tls_dir")
	generatedir := "cv_ca" // maybe we should get this from config

	if strings.Contains("generate", tls_dir) {
		if !strings.Contains("default", my_cert) || !strings.Contains("default", my_priv_key) || !strings.Contains("default", ca_cert) {
			log.Fatal("To use tls_dir=generate, options ca_cert, my_cert, and private_key must all be set to 'default'")
		}

		tls_dir := filepath.Join(common.WORK_DIR, generatedir)
		ca_path := filepath.Join(tls_dir, "cacert.crt")

		if _, err := os.Stat(ca_path); err == nil {
			log.Printf("Existing CA certificate found in %v, not generating a new one", tls_dir)
		} else {
			log.Printf("Generating a new CA in %v and a client certificate for connecting", tls_dir)
			log.Printf("Use `keylime_ca -d %v` to manage this CA", tls_dir)

			if _, err := os.Stat(tls_dir); os.IsNotExist(err) {
				log.Printf("Make directory %v", tls_dir)
				os.Mkdir(tls_dir, os.FileMode(0700))
			}
			if strings.Contains("my_key_pw", my_key_pw) {
				log.Println("CAUTION: using default password for CA, please set private_key_pw to a strong password")
			}
			crypto.CmdInit(tls_dir)
			// ca_util.setpassword(my_key_pw)
			// ca_util.cmd_init(tls_dir)
			// ca_util.cmd_mkcert(tls_dir, socket.gethostname())
			// ca_util.cmd_mkcert(tls_dir, 'client')

		}
	}
}

func GetRestfulParams(myurl string) string {
	u, err := url.Parse(myurl)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(u.RequestURI())
	//values := strings.Split(u.Path, "/")
	//fmt.Println(values)
	params := u.RequestURI()
	return params
}
