package admin_controller

import (
	"edu-imp/internal/admin_dto"
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

	if param.LoginID == "" {
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
	//获取参数
	var param dto.IDParam2
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := service.DelUser(ctx, param.LoginID)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

func UpdateUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin DelUser Controller")
	//获取参数
	var param dto.User
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := service.UpdateUser(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

func AllUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin DelUser Controller")
	//获取参数

	//参数校验

	//调用service
	res, err := service.AllUser(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
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

//暂时写在这个文件
func AddTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin AddTeacher Controller")
	//获取参数
	var param admin_dto.Teacher
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := admin_service.AddTeacher(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

//暂时写在这个文件
func AllTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin AllTeacher Controller")
	//获取参数

	//参数校验

	//调用service
	res, err := admin_service.AllTeacher(ctx)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

//暂时写在这个文件
func DelTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin DelTeacher Controller")
	//获取参数
	var param admin_dto.TeacherParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := admin_service.DelTeacher(ctx, param.Name)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

//暂时写在这个文件
func UpdateTeacher(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin UpdateTeacher Controller")
	//获取参数
	var param admin_dto.UpdateTeacherParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind param fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	//参数校验

	//调用service
	res, err := admin_service.UpdateTeacher(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}
