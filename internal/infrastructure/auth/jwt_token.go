package auth

import (
	"bm-novel/internal/domain/user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	uuid "github.com/satori/go.uuid"
	"time"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func SetToken(auth *user.User) (string, error) {
	claims := jwt.MapClaims{"name": auth.UserName,
		"role": auth.RoleCode,
		"jti":  uuid.NewV4().String(),
		"exp":  time.Now().AddDate(0, 0, 1),
	}
	_, tokenString, err := TokenAuth.Encode(claims)
	fmt.Printf("%s\n", tokenString)

	return tokenString, err
}
