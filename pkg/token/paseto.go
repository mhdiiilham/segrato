package token

import (
	"github.com/mhdiiilham/segrato/config"
	"github.com/o1egl/paseto"
)

type TokenPayload struct {
	ID       string
	Username string
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

func (t TokenService) ExtractToken(accessToken string) (payload TokenPayload, err error) {
	if err = paseto.NewV2().Decrypt(accessToken, []byte(t.Config.PasetoSymmetricKey), &payload, "footer"); err != nil {
		return
	}
	return
}
