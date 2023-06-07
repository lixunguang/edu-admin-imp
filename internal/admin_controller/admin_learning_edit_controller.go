package admin_controller

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/admin_service"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// LearningAll 获取所有学习资源
func LearningAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningAll Controller")
	//获取参数

	//参数校验

	//调用service
	res, err := admin_service.GetLearningAll(ctx)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}

}

//获取学习资源详情（根据id）
func Learning(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "Learning Controller")
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
	res := admin_service.GetLearningByID(ctx, param.ID)

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
	param.Author = "云道教育" //todo：和新闻统一考虑
	//调用service
	res, err := admin_service.AddLearning(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}

}

// LearningDel 删除学习资源
func LearningDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningDel Controller")
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
	res, err := admin_service.DelLearning(ctx, param)

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
	var param admin_dto.UpdateLearning
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
	res, err := admin_service.UpdateLearning(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, dto.IDParam{res})
	}

}

//获取所有分类
func LearningCategory(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningCategory Controller")
	//获取参数

	//参数校验
	//调用service
	res, _ := admin_service.GetLearningCategory(ctx)

	//结果返回

	util.SuccessJson(ctx, res)

}

//增加学习资源分类
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

//删除学习资源分类
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

// LearningResource 获取学习资源的视频资源.
func LearningResource(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LearningResource Controller")

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
	res, err := admin_service.GetLearningResource(ctx, param.ID)

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
	var param admin_dto.AddLearningResourceParam
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
	res, err := admin_service.AddLearningResource(ctx, param)

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
	var param admin_dto.DelLearningResourceParam
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
	res, err := admin_service.DelLearningResource(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}

}
