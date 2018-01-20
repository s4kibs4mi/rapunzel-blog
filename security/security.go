package security

import "golang.org/x/crypto/bcrypt"

/**
 * := Coded with love by Sakib Sami on 20/1/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func NewBCryptPassword(plain string) string {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(encrypted)
}
