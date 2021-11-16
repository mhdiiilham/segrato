package apiresponse

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HealtCheck struct {
	Code       int        `json:"code"`
	Message    string     `json:"message"`
	AuthHealth AuthHealth `json:"authHealth"`
}

type AuthHealth struct {
	Server          bool `json:"server"`
	MongoConnection bool `json:"mongoConnection"`
}
