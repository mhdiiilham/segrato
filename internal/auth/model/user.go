package model

import "github.com/mhdiiilham/segrato/user"

type UserRegiserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResgisterUserResponse struct {
	ID          string    `json:"id"`
	AccessToken string    `json:"accessToken"`
	User        user.User `json:"user"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          string    `json:"id"`
	AccessToken string    `json:"accessToken"`
	User        user.User `json:"user"`
}
