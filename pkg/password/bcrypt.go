package password

import "golang.org/x/crypto/bcrypt"

func Hash(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(b), err
}

func Check(p, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p)) == nil
}
