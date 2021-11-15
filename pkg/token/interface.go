package token

type Service interface {
	SignPayload(payload TokenPayload) (accessToken string, err error)
	ExtractToken(accessToken string) (payload TokenPayload, err error)
}
