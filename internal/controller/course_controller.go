package controller

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

//课程相关接口

func CourseRecommend(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseRecommend Controller")
	//获取参数

	//参数校验
	//调用service
	res := service.GetCourseRecommend(ctx)

	//结果返回

	util.SuccessJson(ctx, res)

}

func CourseAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseAll Controller")
	//获取参数
	var param dto.CourseAllParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		fmt.Println("---->input param: ", param)
		logger.Infoc(ctx, "---->input param: %+v", param)
	}
	//参数校验

	//调用service
	var courseAllRes dto.CourseAllRes
	var err cerror.Cerror
	courseAllRes, err = service.CourseALL(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, courseAllRes)
	} else {
		util.FailJson(ctx, err)
	}
}

//课程概要信息
func CourseOverview(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseOverview Controller")
	//获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		fmt.Println("---->input param: ", param)
		logger.Infoc(ctx, "---->input param: %+v", param)
	}
	//参数校验
	//调用service
	res, err := service.CourseOverview(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

//课程资源表增加
func CourseResourceAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseResourceAdd Controller")
	//获取参数
	var param admin_dto.CourseResourceParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		fmt.Println("---->input param: ", param)
		logger.Infoc(ctx, "---->input param: %+v", param)
	}
	//参数校验
	//调用service
	id, err := service.AddCourseResource(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, id)
	} else {
		util.FailJson(ctx, err)
	}
}

//课程教师表增加
func CourseTeacherAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseTeacherAdd Controller")
	var param admin_dto.CourseTeacherParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}

	if util.IsDebug() {
		fmt.Println("---->input param: ", param)
		logger.Infoc(ctx, "---->input param: %+v", param)
	}

	id, err := service.AddCourseTeacher(ctx, param)

	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, id)
	}
}
