package auth

import (
	"backend/internal/common/server"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Uuid string `json:"uuid"`
}

type CustomClaims struct {
	Claims
	jwt.RegisteredClaims
}

func (c CustomClaims) Validate() error {
	if c.Uuid == "" {
		return errors.New("uuid claim not provided")
	}
	return nil
}

type jwtAuth struct {
	privateKey []byte
	publicKey  []byte
}

func NewJwtAuth(privateKey []byte, publicKey []byte) *jwtAuth {
	return &jwtAuth{privateKey, publicKey}
}

func (j *jwtAuth) Sign(userInfo server.UserInfo) (string, error) {
	customClaims := CustomClaims{
		Claims{Uuid: userInfo.Uuid},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "users",
		},
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, customClaims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *jwtAuth) Verify(token string) (server.UserInfo, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return server.UserInfo{}, err
	}
	parser := jwt.NewParser()
	parsedToken, err := parser.ParseWithClaims(token, CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return server.UserInfo{}, err
	}
	tokenClaims, ok := parsedToken.Claims.(CustomClaims)
	if !ok {
		return server.UserInfo{}, errors.New("custom claims are not associated with token")
	}
	return server.NewUserInfo(tokenClaims.Uuid), nil
}
