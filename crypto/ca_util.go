package crypto

import (
	"github.com/lukehinds/keylime_verifier/config"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// func CmdMkcert(workingdir,name) {
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(cwd)
// }

func SetPassword() {

}

// python function cmd_init:ca_util.py
func CmdInit(tls_dir string) {
	// cwd := common.GetCWD()
	// clear out old crypto files
	log.Println(tls_dir)
	rmfiles := []string{"*.pem", "*.crt", "*.zip", "*.der", "private.yml"}
	for index, itemCopy := range rmfiles {
		_ = index
		full_rmi := tls_dir + "/" + itemCopy
		files, err := filepath.Glob(full_rmi)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			if err := os.Remove(f); err != nil {
				panic(err)
			}
		}
	}

	ca_impl := config.GetConfig("general", "ca_implementation")
	if strings.Contains("cfssl", ca_impl) {
		println("cfssl not implemented yet")
	} else {
		strings.Contains("openssl", ca_impl)
		// cacert, ca_pk, _ = ca_impl.MakeCACert()
		MakeCACert(tls_dir)
	}
}

//     if common.CA_IMPL=='cfssl':
//         pk_str, cacert, ca_pk, _ = ca_impl.mk_cacert()
//     elif common.CA_IMPL=='openssl':
//         cacert, ca_pk, _ = ca_impl.mk_cacert()
//     else:
//         raise Exception("Unknown CA implementation: %s"%common.CA_IMPL)

//     priv=read_private()

//     # write out keys
//     with open('cacert.crt', 'wb') as f:
//         f.write(cacert.as_pem())

//     f = BIO.MemoryBuffer()
//     ca_pk.save_key_bio(f,None)
//     priv[0]['ca']=f.getvalue()
//     f.close()

//     # store the last serial number created.
//     # the CA is always serial # 1
//     priv[0]['lastserial'] = 1

//     write_private(priv)

//     ca_pk.get_rsa().save_pub_key('ca-public.pem')

//     # generate an empty crl
//     if common.CA_IMPL=='cfssl':
//         crl = ca_impl.gencrl([],cacert.as_pem(), pk_str)
//     elif common.CA_IMPL=='openssl':
//         crl = ca_impl.gencrl([],cacert.as_pem(),str(priv[0]['ca']))
//     else:
//         raise Exception("Unknown CA implementation: %s"%common.CA_IMPL)

//     if isinstance(crl, str):
//         crl = crl.encode('utf-8')

//     with open('cacrl.der','wb') as f:
//         f.write(crl)
//     convert_crl_to_pem("cacrl.der","cacrl.pem")

//     # Sanity checks...
//     cac = X509.load_cert('cacert.crt')
//     if cac.verify():
//         logger.info("CA certificate created successfully in %s"%workingdir)
//     else:
//         logger.error("ERROR: Cert does not self validate")
// finally:
//     os.chdir(cwd)

// setpassword

// cmd_mkcert

// cmd_init

// cmd_certpkg

// convert_crl_to_pem

// get_crl_distpoint

// cmd_revoke

// cmd_regencrl

// cmd_listen

// CRLHandler(BaseHTTPRequestHandler):

// rmfiles

// write_private

// read_private
