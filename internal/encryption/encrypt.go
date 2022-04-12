package encryption

import (
	"github.com/Luzifer/go-openssl/v4"
)

func Encrypt(p string, k string) ([]byte, error) {
	o := openssl.New()
	return o.EncryptBytes(k, []byte(p), openssl.BytesToKeyMD5)
}
