package apiresponse

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
