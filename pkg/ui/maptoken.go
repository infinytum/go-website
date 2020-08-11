package ui

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func maptoken(w http.ResponseWriter, req *http.Request) {
	var block *pem.Block
	var x509Encoded []byte
	var err error
	var privateKeyI interface{}
	var privateKey *ecdsa.PrivateKey
	var ok bool

	//decode the pem format
	block, _ = pem.Decode([]byte(ApplicationContext.Config().Get("maps", "key").String("")))
	//check if its is private key
	if block == nil || block.Type != "PRIVATE KEY" {
		return
	}

	//get the encoded bytes
	x509Encoded = block.Bytes

	//generate the private key object
	privateKeyI, err = x509.ParsePKCS8PrivateKey(x509Encoded)
	if err != nil {
		return
	}
	//cast into ecdsa.PrivateKey object
	privateKey, ok = privateKeyI.(*ecdsa.PrivateKey)
	if !ok {
		return
	}

	token := jwt.New(jwt.SigningMethodES256)
	token.Header["kid"] = ApplicationContext.Config().Get("maps", "kid").String("")
	token.Header["typ"] = "JWT"
	token.Claims = jwt.MapClaims{
		"iss": ApplicationContext.Config().Get("maps", "iss").String(""),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * 60).Unix(),
	}

	tokenString, err := token.SignedString(privateKey)
	w.Write([]byte(tokenString))
}
