package security

import "golang.org/x/crypto/bcrypt"

// Hash receives a string and adds a hash
func Hash(passwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
}

// VerifyPassword compares a password and a hash and returns if they are equal
func VerifyPassword(hashPasswd, strPasswd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPasswd), []byte(strPasswd))
}