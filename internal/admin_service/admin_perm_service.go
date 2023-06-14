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

//检查是否和当前存储的token是否一致
func CheckCurrentToken(userID string, token string) cerror.Cerror {
	str := AdminloginMap[userID]
	if str == "" { //未有登录记录
		return cerror.ErrorUserNotLogin
	}

	if str != token { //token已经更新
		return cerror.ErrorTokenAuthFailed

	}
	//已经登录且token匹配
	return cerror.ErrorUserAuthSucc
}

//检查用户名和密码
func UserLogin(ctx *gin.Context, name string, password string) (string, cerror.Cerror) {
	var res cerror.Cerror
	err := dao.CheckAdmin(ctx, name, password)

	var newTokenStr string
	if err.Code() == cerror.ErrorUserAuthSucc.Code() { //用户名，密码正确

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

	//检查是否已经登录过，

	//如果是，则之前登录token会被失效
	//如果否，继续登录

	tokenStr, res := UserLogin(ctx, userName, password)
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
