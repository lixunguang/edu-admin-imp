package middleware

import (
	"edu-imp/internal/common"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var SecretKey = []byte("YundaoEdu")

type CustomClaims struct {
	UserName string `json:"user_name"`
	UserID   string `json:"user_id"`
	Exp      int64  `json:"exp"`
	jwt.StandardClaims
}

func GenerateToken(username string, loginId string) (string, error) {
	//  Create the token
	// token := jwt.New(jwt.SigningMethodHS256)
	//  Set some claims
	// claims := token.Claims.(jwt.MapClaims)

	//  Create the Claims
	claims := CustomClaims{username,
		loginId,
		time.Now().Add(time.Minute * 60).Unix(),
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	// claims["username"] = username
	// claims["userid"] = loginId
	// claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
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

func CheckAdminAuth(c *gin.Context) {

	token := c.GetHeader("Authorization")
	logger.Infoc(c, "checkAuth:%p\n", token)

	err := parseAdminToken(token)

	if err != nil {
		util.FailJson(c, err)
		c.Abort()
		return
	}

	c.Next()

}

func parseAdminToken(tokenString string) cerror.Cerror {
	fmt.Println("token is=", tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YundaoEdu"), nil
	})

	fmt.Printf("%+v", token)

	if claims, ok := token.Claims.(*CustomClaimsAdmin); ok && token.Valid {
		fmt.Printf("%+v", claims)

		err1 := claims.VerifyExpiresAt(time.Now().Unix(), true)
		if err1 {
			fmt.Println("Token not expired.")
			return nil
		}

		return common.ErrorTokenExpire
	} else {
		fmt.Println(err)
	}

	return common.ErrorTokenExpire
}

func CheckAuth(c *gin.Context) {

	token := c.GetHeader("Authorization")

	logger.Infoc(c, "checkAuth:%p\n", token)
	err := parseToken(token)

	if err != nil {
		util.FailJson(c, err)
		c.Abort()
		return
	}

	c.Next()

}

func parseToken(tokenString string) cerror.Cerror {

	fmt.Println("token is=", tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YundaoEdu"), nil
	})

	fmt.Printf("%+v", token)

	if token == nil {
		fmt.Printf("err:token is null")
		return cerror.NewCerror(common.TokenFormatError, err.Error())
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		// fmt.Println(claims)
		fmt.Printf("%+v", claims)

		fmt.Println(claims.ExpiresAt, time.Now().Unix())

		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil
		}

		return cerror.NewCerror(common.TokenExpired, "token过期")
	} else {
		fmt.Println(err)
	}

	return cerror.NewCerror(common.TokenExpired, err.Error())
}
