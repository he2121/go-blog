package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"user/internal/svc"
)

func GetJwtToken(ctx *svc.ServiceContext, userID int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + ctx.Config.LoginAuth.AccessExpire
	claims["iat"] = ctx.Config.LoginAuth.AccessExpire
	claims["userID"] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(ctx.Config.LoginAuth.AccessSecret))
}
