package service

import (
	"context"
	"fmt"
	"myapp/graph/model"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//UserClaim Token Claim
type UserClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

//TokenGenerate Generate Token
func TokenGenerate(ctx context.Context, input model.User) (*model.AuthentificationToken, error) {
	var signingMethod = jwt.SigningMethodHS256
	var expiredDate = time.Now().UTC().AddDate(0, 0, 1).UnixNano() / int64(time.Millisecond)

	customClaim := UserClaim{
		ID: input.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredDate,
		},
	}

	token := jwt.NewWithClaims(signingMethod, customClaim)

	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		panic(err)
	}

	return &model.AuthentificationToken{
        Type: "Bearer",
        Token: signedToken,
    }, nil
}

//TokenValidate Validate JWT Token
func TokenValidate(ctx context.Context, t string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(t, UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtSecret, nil
	})

	if err != nil {
		panic(err)
	}

	return token, nil
}
