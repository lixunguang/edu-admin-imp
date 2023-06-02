package controller

import (
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 学习相关接口

// LearningAdd 增加学习资源
func LearningAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningAdd Controller")
	//获取参数
	var param dto.Learning
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
	res, err := service.AddLearning(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}

}

//增加资源分类
func LearningCategoryAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningCategoryAdd Controller")

	//获取参数
	var param dto.LearningCategory
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
	res, err := service.AddLearningCategory(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}

}

// LearningItemAdd 增加资源条目.
func LearningItemAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningItemAdd Controller")

	//获取参数
	var param dto.LearningItemParam
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
	res, err := service.AddLearningItem(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}

}

func LearningResourceRecommend(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "learningResourceRecommend Controller")
	//获取参数

	//参数校验
	//调用service
	res := service.GetLearningRecommend(ctx)

	//结果返回

	util.SuccessJson(ctx, res)

}

//获取学习资源详情（根据id）
func LearningResource(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "learningResource Controller")
	//获取参数
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v,param:%+v", "method", err, ctx.Request)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		fmt.Println("---->input param: ", param)
		logger.Infoc(ctx, "---->input param: %+v", param)
	}
	//参数校验
	//调用service
	res := service.GetLearningByID(ctx, param.ID)

	//结果返回

	util.SuccessJson(ctx, res)

}

func LearningCategory(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "learningCategory Controller")
	//获取参数

	//参数校验
	//调用service
	res, _ := service.GetLearningCategory(ctx)

	//结果返回

	util.SuccessJson(ctx, res)

}

func LearningCategoryResourceRecommend(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "learningCategoryResourceRecommend Controller")

	//获取参数

	//参数校验
	var param dto.IDParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v,param:%+v", "method", err, ctx.Request)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	if util.IsDebug() {
		fmt.Println("---->input param: ", param)
		logger.Infoc(ctx, "---->input param: %+v", param)
	}
	//调用service
	res, _ := service.GetLearningCategoryResourceRecommend(ctx, param)

	//结果返回

	util.SuccessJson(ctx, res)

}

func LearningCategoryResource(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "learningCategoryResource Controller")
	//获取参数
	var param dto.LearningCategoryParam
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
	var res dto.LearningCategoryResourceRes
	var err cerror.Cerror
	res, err = service.GetLearningCategoryResourcePagedData(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}
