package admin_controller

import (
	"edu-imp/internal/admin_service"
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin AddUser Controller")
	//获取参数
	var param dto.User
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := service.AddUser(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func DelUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin DelUser Controller")

}

func AddAdmin(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin AddAdmin Controller")
	//获取参数
	var param dto.Admin
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := admin_service.AddAdmin(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func DelAdmin(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin DelAdmin Controller")

	//获取参数
	var param dto.AdminParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := admin_service.DelAdmin(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

//分页获取
func GetAdmin(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin GetAdmin Controller")

}

type teacherParam struct {
	Name string `json:"name"`
}

//暂时写在这个文件
func AddTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin AddTeacher Controller")
	//获取参数
	var param teacherParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := admin_service.AddTeacher(ctx, param.Name)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}
