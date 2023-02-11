package hash

import "golang.org/x/crypto/bcrypt"

const DefaultWorkFactor = bcrypt.DefaultCost

func Password(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), DefaultWorkFactor)

	return string(hash)
}
