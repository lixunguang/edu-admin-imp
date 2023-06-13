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

func IsLogin(userID string) bool {
	str := AdminloginMap[userID]
	if str == "" {
		return false
	}
	return true
}

//检查用户名和密码
func ValidateUserLogin(ctx *gin.Context, name string, password string) (string, cerror.Cerror) {
	var res cerror.Cerror
	err := dao.CheckAdmin(ctx, name, password)

	var newTokenStr string
	if err.Code() == common.ErrorOK.Code() { //用户名，密码正确

		newTokenStr, _ = middleware.GenerateTokenAdmin(name)

		if oldTokenStr, ok := AdminloginMap[name]; ok { //检查是否已经登录
			//如果已经登录，则更新token
			logger.Debugf("%s", oldTokenStr)
			res = common.ErrorLoginAgain
		} else {
			res = common.ErrorOK
		}

		AdminloginMap[name] = newTokenStr
	} else {
		return newTokenStr, err
	}

	return newTokenStr, res
}

func Login(ctx *gin.Context, userName string, password string) (string, cerror.Cerror) {

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
