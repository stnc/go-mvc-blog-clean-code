package security

import "golang.org/x/crypto/bcrypt"

/*
//
	hashPassword, _ := security.Hash(pass)
	fmt.Println("ŞİFELENMİŞ HALİ" +hashPassword )
	$2a$10$QPiWAgMpwHBkDjBL5pPd2.HBlfdniuGOvZd5kh.ILLjKFo67rvfsO
*/
//Hash is code hash yaparak verir
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

//VerifyPassword doğrulama
func VerifyPasswordApi(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
