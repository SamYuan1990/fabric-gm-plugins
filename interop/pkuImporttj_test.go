package interop

import (
	"testing"

	"github.com/Hyperledger-TWGC/pku-gm/gmssl"
)

func TestPKUSM2KeyPemImport(t *testing.T) {
	var privKeyPem = "testdata/tjfoc/priv.pem"
	pemBytes := ReadFile(privKeyPem, t)
	_, err := gmssl.NewPrivateKeyFromPEM(string(pemBytes), "")
	Fatal(err, t)
}
