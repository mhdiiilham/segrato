package token

import (
	"github.com/mhdiiilham/segrato/config"
	"github.com/o1egl/paseto"
)

type TokenPayload struct {
	ID        string
	Username  string
	IsPremium bool
}

type TokenService struct {
	Config *config.Config
}

func (t TokenService) SignPayload(payload TokenPayload) (accessToken string, err error) {
	accessToken, err = paseto.NewV2().Encrypt([]byte(t.Config.PasetoSymmetricKey), payload, "footer")
	if err != nil {
		return
	}
	return
}
