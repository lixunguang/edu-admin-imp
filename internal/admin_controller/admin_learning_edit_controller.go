package admin_controller

import (
	"edu-imp/internal/admin_service"
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// LearningAll 获取所有学习资源
func LearningAll(ctx *gin.Context) {
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

//获取学习资源详情（根据id）
func Learning(ctx *gin.Context) {
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

// LearningDel 删除学习资源
func LearningDel(ctx *gin.Context) {
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

// LearningUpdate 更新学习资源
func LearningUpdate(ctx *gin.Context) {
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

//获取所有分类
func LearningCategoryAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningCategoryAll Controller")
	//获取参数

	//参数校验
	//调用service
	res, _ := admin_service.GetLearningCategory(ctx)

	//结果返回

	util.SuccessJson(ctx, res)

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
	res, err := admin_service.AddLearningCategory(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}

}

//删除资源分类
func LearningCategoryDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningCategoryAdd Controller")

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
	res, err := admin_service.DelLearningCategory(ctx, param.ID)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDRes{res})
	}

}

// LearningResourceAdd 增加资源条目.
func LearningResourceAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningResourceAdd Controller")

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

// LearningResourceAdd 增加资源条目.
func LearningResourceAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningResourceAdd Controller")

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

// LearningResourceDel 增加资源条目.
func LearningResourceDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningResourceAdd Controller")

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