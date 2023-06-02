package controller

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 权限相关接口

// gwt，token串
// token存储，压缩，内存，数据库
// 重复登录，存储新的token，之前token失效。
// 过期检查

//  /edu/v1/perm/login 登录
func Login(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "Login Controller")
	// 获取参数
	var param dto.LoginParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		fmt.Println("---->input param: ", param)
		logger.Infoc(ctx, "---->input param: %+v", param)
	}
	// 参数校验

	// 调用service
	tokenStr, res := service.Login(ctx, param.Name, param.LoginId, param.Password)

	// 结果返回
	//todo:处理下结果返回，从service层返回err
	if res == common.SuccessLoginAgain || res == common.SuccessLogin {
		var loginRes dto.LoginRes
		loginRes.Token = tokenStr
		loginRes.FirstLogin = firstLogin()

		util.SuccessJson(ctx, loginRes)

	} else if res == common.ErrorFindUser {
		util.FailJson(ctx, cerror.ErrorFindUser)
	} else if res == common.ErrorPassword {
		util.FailJson(ctx, cerror.ErrorPassword)
	}

}

//是否首次登录
func firstLogin() bool {
	return true
}

// /edu/v1/perm/logout
func Logout(ctx *gin.Context) {

	//获取参数
	tokenStr := ctx.GetHeader("Authorization")

	//参数校验

	//调用service
	res := service.Logout(ctx, tokenStr)

	//结果返回

	if res == common.Success {
		util.SuccessJson(ctx, nil)
	} else {
		util.FailJson(ctx, cerror.InvalidParams)
	}

}
