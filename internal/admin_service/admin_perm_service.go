package admin_service

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/middleware"
	"edu-imp/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

//todo:并发锁 sync.map
var AdminloginMap = map[string]string{} /*记录登录用户信息 */

//检查登录是否过期
func CheckExpired(tokenStr string) int {
	token, err := jwt.ParseWithClaims(tokenStr, &middleware.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(middleware.SecretKey), nil
	})

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		//fmt.Println(claims)
		fmt.Printf("%+v", claims)
		fmt.Println(claims.ExpiresAt, time.Now().Unix())

		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			fmt.Println("未过期")
			return common.NotExpired
		} else {
			fmt.Println("已过期。。。")
			return common.TokenExpired
		}

	} else {
		fmt.Println(err)
	}

	return common.NotExpired
}

//检查用户名和密码
func ValidateUserLogin(ctx *gin.Context, userName string, password string) (string, int) {

	res, _ := dao.CheckAdmin(ctx, userName, password)

	var newTokenStr string
	if res == common.CheckOK { //用户名，密码正确

		newTokenStr, _ = middleware.GenerateTokenAdmin(userName)

		if oldTokenStr, ok := AdminloginMap[userName]; ok { //检查是否已经登录
			//如果已经登录，则更新token
			logger.Debugf("%s", oldTokenStr)
			res = common.SuccessLoginAgain
		} else {
			res = common.SuccessLogin
		}

		AdminloginMap[userName] = newTokenStr
	}

	return newTokenStr, res
}

func Login(ctx *gin.Context, userName string, password string) (string, int) {

	tokenStr, res := ValidateUserLogin(ctx, userName, password)
	logger.Debugf("%s,%d", tokenStr, res)

	return tokenStr, res
}

func Logout(ctx *gin.Context, tokenStr string) int {

	for k, v := range AdminloginMap {
		if v == tokenStr {
			delete(AdminloginMap, k)
		}
	}

	return common.Success

}
