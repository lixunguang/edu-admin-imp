package middleware

import (
	"edu-imp/internal/common"
	"edu-imp/pkg/cerror"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SecretKey = []byte("YundaoEdu")

type CustomClaims struct {
	UserName string `json:"user_name"`
	UserID   string `json:"user_id"`
	Exp      int64  `json:"exp"`
	jwt.StandardClaims
}

type CustomClaimsAdmin struct {
	UserName string `json:"user_name"`
	//Exp      int64  `json:"exp"`
	jwt.StandardClaims
}

//后台管理
func GenerateTokenAdmin(username string) (string, error) {

	//  Create the Claims
	claims := CustomClaimsAdmin{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    fmt.Sprintf("%d", time.Now().Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ParseAdminToken(tokenString string) (string, cerror.Cerror) {
	fmt.Println("token is=", tokenString)

	var claims CustomClaimsAdmin
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("YundaoEdu"), nil
	})

	if token == nil {
		return "", common.ErrorTokenFormat
	}

	if err != nil {
		return "", common.ErrorTokenExpire
	}

	if claims, ok := token.Claims.(*CustomClaimsAdmin); ok && token.Valid {
		fmt.Printf("%+v", claims)

		err1 := claims.VerifyExpiresAt(time.Now().Unix(), true)
		if err1 {
			fmt.Println("Token not expired.")
			return claims.UserName, nil
		}

		return "", common.ErrorTokenExpire
	} else {
		fmt.Println(err)
	}

	return "", common.ErrorTokenExpire
}
