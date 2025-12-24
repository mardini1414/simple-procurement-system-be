package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	b := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
