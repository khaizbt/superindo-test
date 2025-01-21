package config

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type (
	AuthService interface {
		GenerateTokenUser(token string) (string, error)
		ValidateToken(encodedToken string) (*jwt.Token, error)
		RevokeToken(userID int) (string, error)
	}

	jwtService struct {
	}

	MyClaims struct {
		jwt.StandardClaims
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)

// var LOGIN_EXP = time.Duration(1) * time.Hour //Token wil exp in 1 Hours if no activity
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewServiceAuth() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateTokenUser(userID string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claim)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *jwtService) RevokeToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claim)
	token_revoke, err := token.SignedString([]byte("aiueoi21"))

	if err != nil {
		panic(err.Error())
	}

	return token_revoke, nil
}
