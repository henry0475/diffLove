package authorization

import (
	"crypto/rsa"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	configure "github.com/henry0475/diffLove/config"
	"github.com/henry0475/diffLove/libs/foundation"
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

func (auth *Authorization) Login(userName string, password string) (userInfo *UserInfo, err error) {
	stmt, err := foundation.GetMysqlClient().Prepare("SELECT `id`,`gender`,`username`,`nickname` FROM `diffLove_db`.`users` WHERE `username`=? AND `password`=? LIMIT 1")
	if err != nil {
		return
	}
	defer stmt.Close()

	userInfo := UserInfo{}
	err = stmt.QueryRow(userName, getHashedPassword(password)).Scan(&userInfo.ID, &userInfo.Gender, &userInfo.UserName, &userInfo.NickName)
	if err != nil {
		return
	}

	userInfo.Token = createToken(userInfo.UserName, conf.Validation.Token.Web)
	return
}

func getHashedPassword(originalPassword string) string {
	r := sha1.Sum([]byte(originalPassword + conf.Security.Salt))
	return hex.EncodeToString(r[:])
}
