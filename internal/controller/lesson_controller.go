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

// 课堂相关接口

//课堂概要信息
func CourseLessonOverview(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseLessonOverview Controller")
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
	res, err := service.CourseLessonOverview(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

//根据lessonid 获取课程内容
func CourseLessonContent(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseLessonContent Controller")
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
	var sectionType int
	sectionType = common.LessonContent //内容类型
	res, err := service.GetCourseLessonSectionResource(ctx, param.ID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

//根据lessonid 获取课程参考资源
func CourseLessonRefer(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseLessonRefer Controller")
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
	var sectionType int
	sectionType = common.LessonRefer //参考资源类型
	res, err := service.GetCourseLessonSectionResource(ctx, param.ID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

//根据lessonid 获取实验内容
func CourseLessonExperiment(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseLessonExperiment Controller")
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
	var sectionType int
	sectionType = common.LessonExperiment //实验
	res, err := service.GetCourseLessonSectionResource(ctx, param.ID, sectionType)

	//结果返回
	if err == nil && len(res) >= 1 {
		util.SuccessJson(ctx, res[0])
	} else {
		util.FailJson(ctx, err)
	}

}

//根据lessonid 获取课程所布置的作业
func CourseLessonWork(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "CourseLessonWork Controller")
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
	var sectionType int
	sectionType = common.LessonWork //作业
	res, err := service.CourseLessonSectionWorkResource(ctx, param.ID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}
