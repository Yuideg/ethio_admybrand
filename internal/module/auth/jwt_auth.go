package auth

import (
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"os"
	"time"
)

var secret = os.Getenv("JWT_SECRET")

func CreateToken(id uuid.UUID, username string) (string, error) {
	claims := &model.Claim{
		UserID:    id,
		Username:  username,
		Authorize: true,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(), //Token expires after 1 hour
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))

}
