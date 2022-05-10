package auth

import (
	"time"

	"github.com/marlonmp/go-basic-library/pkg/jwt"
)

type AccessTokenResponse struct {
	AccesToken   string `json:"acces_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

func NewAccesTokenResponse(subject string) AccessTokenResponse {

	expiresIn := time.Hour * 2

	accessToken := jwt.New(expiresIn, subject)

	refreshToken := jwt.New(time.Hour*168, subject)

	return AccessTokenResponse{
		AccesToken:   accessToken,
		ExpiresIn:    int64(expiresIn),
		RefreshToken: refreshToken,
		Scope:        "",
		TokenType:    jwt.TokenType,
	}
}

type UserCredentials struct {
	// Reprecents the subject: username, email, identification, etc.
	Subject  string `json:"subject"`
	Password string `json:"password"`
}
