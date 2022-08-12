package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

// NOT CURRENTLY IN USE
func ReplacePassword(hashedPassword, password, newPassword string) ([]byte, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, err
	}
	return bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)

}
