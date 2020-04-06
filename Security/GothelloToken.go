package Security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)
const (
	hmacSecret          = "qOf5Xxqhp5TnlCxULjmqgcAznxGif3dwTrNGLUUqVnzMx0cm1Ti6m6l2TB8nT44t"
	bearerSchema        = "Bearer "
	authorizationHeader = "Authorization"
)

type gothelloClaims struct {
	PlayerId int `json:"id"`
	jwt.StandardClaims
}

func CreateGothelloToken(playerId int) (string, error) {
	claims := gothelloClaims{
		PlayerId: playerId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt: time.Now().Unix(),
			NotBefore: time.Date(2020, 04, 01,00,00,00,00, time.Local).Unix(),
		},
	}
	var key = []byte(hmacSecret)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}

func GetPlayerIdFromToken(token string) int {
	var claims gothelloClaims
	jwtToken, _ := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		var key = []byte(hmacSecret)
		return key, nil
	})
	if jwtToken.Valid && claims.StandardClaims.Valid() == nil {
		return claims.PlayerId
	}
	return -1
}

func GetTokenFromRequestHeader(header http.Header) (string, error) {
	reqToken := header.Get(authorizationHeader)
	splitToken := strings.Split(reqToken, bearerSchema)
	if len(splitToken) != 2 {
		return "", errors.New("invalid header")
	}
	return splitToken[1], nil
}