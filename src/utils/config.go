package utils

import (
	"os"
)

// IsProduction return production state
func IsProduction() bool {

	if os.Getenv("NODE_ENV") == "production" {
		return true
	} else {
		return false
	}
}

// APIPort returns API_PORT
func APIPort() string {
	return os.Getenv("API_PORT")
}

// MysqlHost returns DB_HOST
func MysqlHost() string {
	return os.Getenv("DB_HOST")
}

// MysqlPort returns DB_PORT
func MysqlPort() string {
	return os.Getenv("DB_PORT")
}

// MysqlDBName returns DB_NAME
func MysqlDBName() string {
	return os.Getenv("DB_NAME")
}

// MysqlUsername returns DB_USERNAME
func MysqlUsername() string {
	return os.Getenv("DB_USERNAME")
}

// MysqlPassword returns DB_PASSWORD
func MysqlPassword() string {
	return os.Getenv("DB_PASSWORD")
}

// FirebaseDbRef returns FIREBASE_DB_REF
func FirebaseDbRef() string {
	return os.Getenv("FIREBASE_DB_REF")
}

// JwtSecret returns JWT_SECRET
func JwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

// WebtokenSecretRefresh returns WEBTOKEN_SECRETKEY_REFRESHKEY
func WebtokenSecretRefresh() string {
	return os.Getenv("WEBTOKEN_SECRETKEY_REFRESHKEY")
}

// WebtokenSecretAccess returns WEBTOKEN_SECRETKEY_ACCESSKEY
func WebtokenSecretAccess() string {
	return os.Getenv("WEBTOKEN_SECRETKEY_ACCESSKEY")
}

// AccessTokenExpire returns ACCESS_TOKEN_EXPIRE
func AccessTokenExpire() string {
	return os.Getenv("ACCESS_TOKEN_EXPIRE")
}
