package service

import (
	"encoding/json"
	"github.com/ericyan/jwk"
	"github.com/gogf/gf/database/gdb"
	"jwkserver/app/model/pgdb"
	"jwkserver/library"
	"time"
)
var mydb,err=pgdb.Newdb()

func Insertjwk() (string,error){
	keyid, jwkpublicKey,jwkprivateKey, _ := library.Genjwkkey()
	pvk, _ := json.Marshal(jwkprivateKey)
	puk, _ := json.Marshal(jwkpublicKey)
	_, err = mydb.Insert("go.jwkeys", gdb.Map{
		"ulid":       keyid,
		"privatekey": pvk,
		"publickey":  puk,
		"creatstime": time.Now(),
	})
	return keyid,err
}

func Selectpukjwk(keyid string) (*jwk.RSAPublicKey,error,error) {
	r, err := mydb.GetValue("select publickey from go.jwkeys where ulid=?", keyid)
	publickey, error := jwk.ParseRSAPublicKey(r.Bytes())
	return publickey,err,error
}
func Selectpvkjwk(keyid string) (*jwk.RSAPrivateKey,error,error) {
	r,err := mydb.GetValue("select privatekey from go.jwkeys where ulid=?",keyid)
	privatekey,error:=jwk.ParseRSAPrivateKey(r.Bytes())
	return privatekey,err,error
}
