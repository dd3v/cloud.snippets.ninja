package security

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//GenerateHashFromPassword - ...
func GenerateHashFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash[:]), err
}

//CompareHashAndPassword - ...
func CompareHashAndPassword(hash string, password string) bool {

	fmt.Println(hash);
	fmt.Println(password);

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
