package password

type Service interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}
