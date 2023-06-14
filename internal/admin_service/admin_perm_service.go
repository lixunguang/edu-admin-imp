package admin_service

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/middleware"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

//todo:并发锁 sync.map
var AdminloginMap = map[string]string{} /*记录登录用户信息 */

//是否已经登陆过
func HasLogin(userID string, token string) cerror.Cerror {
	str := AdminloginMap[userID]

	if str == "" { //未有登录记录
		return cerror.ErrorUserNotLogin
	}

	if str == token { //已经登录且token匹配
		return cerror.ErrorLoginAgain
	}

	//token不匹配
	return cerror.ErrorLoginFailed
}

//检查用户名和密码
func ValidateUserLogin(ctx *gin.Context, name string, password string) (string, cerror.Cerror) {
	var res cerror.Cerror
	err := dao.CheckAdmin(ctx, name, password)

	var newTokenStr string
	if err.Code() == cerror.ErrorLoginSucc.Code() { //用户名，密码正确

		newTokenStr, _ = middleware.GenerateTokenAdmin(name)

		if oldTokenStr, ok := AdminloginMap[name]; ok { //检查是否已经登录
			//如果已经登录，则更新token
			logger.Debugf("oldTokenStr:%s", oldTokenStr)
			res = cerror.ErrorLoginAgain
		} else {
			res = cerror.ErrorLoginSucc
		}

		AdminloginMap[name] = newTokenStr
		logger.Debugf("newTokenStr:%s", newTokenStr)
	} else {
		return newTokenStr, err
	}

	return newTokenStr, res
}

func Login(ctx *gin.Context, userName string, password string) (string, cerror.Cerror) {

	tokenStr, res := ValidateUserLogin(ctx, userName, password)
	logger.Debugf("tokenStr:%s,%d", tokenStr, res)

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
