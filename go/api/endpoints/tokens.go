package endpoints

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY string = "Secret_key_a45556"

// JWTDetails
type JWTDetails struct {
	Username string
	Uid      string
	jwt.StandardClaims
}

// GenerateTokens generates both the detailed token and refresh token
func GenerateTokens(email string, firstName string, lastName string, uid string) (*string, *string, error) {
	claims := &JWTDetails{
		Username: email,
		Uid:      uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	refreshClaims := &JWTDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(2160)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return nil, nil, err
	}

	return &token, &refreshToken, err
}
