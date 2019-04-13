package authorization

import (
	"crypto/rsa"
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"time"

	configure "github.com/henry0475/diffLove/config"
	"github.com/henry0475/diffLove/libs/foundation"
	"github.com/dgrijalva/jwt-go"
)

var (
	conf       configure.Config
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

func init() {
	conf = foundation.GetConf()

	publicKeyByte, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/github.com/henry0475/diffLove/config/diffLove.rsa.pub")
	if err != nil {
		log.Println("Error: there is no public key")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)

	privateKeyByte, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/github.com/henry0475/diffLove/config/diffLove.rsa")
	if err != nil {
		log.Println("Error: there is no private key")
	}
	privateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
}

func (auth *Authorization) Login(userName string, password string) (err error) {
	
}