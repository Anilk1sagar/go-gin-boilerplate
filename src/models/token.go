package models

import jwt "github.com/dgrijalva/jwt-go"

// AuthToken ...// This is what is retured to the user
type AuthToken struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"accessToken"`
	ExpiresIn int64  `json:"expiresIn"`
}

// AuthTokenClaim ...
// This is the cliam object which gets parsed from the authorization header
type AuthTokenClaim struct {
	*jwt.StandardClaims
	UserID uint
}
