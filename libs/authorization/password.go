package authorization

import (
	"crypto/sha1"
	"encoding/hex"
)

func getCryptedPassword(pw string) string {
	r := sha1.Sum([]byte(pw + conf.Security.Salt))
	return hex.EncodeToString(r[:])
}
