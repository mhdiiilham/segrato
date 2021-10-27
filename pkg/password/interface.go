package password

type Service interface {
	HashPassword(password string) (string, error)
}
