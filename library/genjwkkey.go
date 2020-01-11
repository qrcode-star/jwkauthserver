package library

import (
	"crypto/rsa"
	"github.com/gogf/gf/util/guuid"
	cryptorand "crypto/rand"
	"github.com/ericyan/jwk"
)


func Genjwkkey() (string, jwk.Key, jwk.Key, error) {
	keyid ,_:= guuid.NewRandom()
	privateKey, error := rsa.GenerateKey(cryptorand.Reader, 2048)
	publicKey := privateKey.PublicKey
	param := jwk.Params{KeyID: keyid.String()}
	jwkpublicKey, _ := jwk.New(&publicKey, &param)
	jwkprivateKey, _ := jwk.New(privateKey, &param)
	return keyid.String(), jwkpublicKey, jwkprivateKey, error
}