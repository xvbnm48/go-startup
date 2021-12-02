package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

// service ini mengcreate JWT dan Memvalidasi token JWT

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("NYABIRRA_KAWAII")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	//membuat payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	//membuat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//mentanda tangani token
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err

	}

	return token, nil

}
