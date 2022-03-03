package password

import (
	"golang.org/x/crypto/bcrypt"
	"root/gleam/golang/tool/logging"
)


func HashAndSalt(pwd string) string{
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		logging.Error(err)
	}
	return string(hash)
}

func verifyPasswords(hashedPwd, currentPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd),[]byte(currentPwd))
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}

