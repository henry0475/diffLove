package authorization

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

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

// Login will verify the valid of username and password
func (auth *Authorization) Login(userName string, password string) (userInfo *UserInfo, err error) {
	stmt, err := foundation.GetMysqlClient().Prepare("SELECT `id`,`gender`,`username`,`nickname` FROM `diffLove_db`.`users` WHERE `username`=? AND `password`=? LIMIT 1")
	if err != nil {
		return
	}
	defer stmt.Close()

	userInfo = new(UserInfo)
	err = stmt.QueryRow(userName, getCryptedPassword(password)).Scan(&userInfo.ID, &userInfo.Gender, &userInfo.UserName, &userInfo.NickName)
	if err != nil {
		return
	}

	userInfo.Token, err = createToken(userInfo.UserName, conf.Validation.Token.Web)
	return
}

// Register will register a user to system
func (auth *Authorization) Register(username string, password string, gender int) (userInfo *UserInfo, err error) {
	// Check whether this username is occupied
	stmt, err := foundation.GetMysqlClient().Prepare("SELECT `id` FROM `diffLove_db`.`users` WHERE `username`=? LIMIT 1")
	if err != nil {
		return
	}
	defer stmt.Close()

	userInfo = new(UserInfo)
	_ = stmt.QueryRow(username).Scan(&userInfo.ID)
	if userInfo.ID > 0 {
		return nil, errors.New("Error: this username is not valid")
	}

	tx, err := foundation.GetMysqlClient().Begin()
	if err != nil {
		return
	}

	res, err := tx.Exec("INSERT INTO `diffLove_db`.`users`(`gender`,`username`,`add_time`,`password`) VALUES(?,?,?,?)", gender, username, time.Now().Unix(), getCryptedPassword(password))
	if err != nil {
		return
	}

	id, _ := res.LastInsertId()
	userInfo.ID = uint64(id)
	userInfo.UserName = username
	userInfo.Gender = int(gender)
	userInfo.Token, err = createToken(username, conf.Validation.Token.Web)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

// GetUID will get the username from token's string
func (auth *Authorization) GetUID() (uid string, err error) {
	if auth.Token == "" {
		return uid, errors.New("Error: There is no token which is available")
	}

	return getUIDFromToken(auth.Token)
}
