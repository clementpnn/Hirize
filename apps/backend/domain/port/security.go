package port

type SecurityService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) error
}
