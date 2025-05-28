package gateway

type CryptorGateway interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
