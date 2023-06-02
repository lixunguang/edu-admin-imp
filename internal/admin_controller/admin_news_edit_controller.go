// Package controller is ...
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

// NewsAdd is 增加新闻.
func NewsAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "NewsAdd Controller")
	// 获取参数
	var param dto.NewsAddParam
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
	res, err := admin_service.AddNews(ctx, param)

	// 结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

// NewsDelete is 刪除新闻.
func NewsDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "NewsDelete Controller")
	// 获取参数
	var param dto.DelNewsParam
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
	res, err := admin_service.DelNews(ctx, param.ID)

	// 结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

// 编辑新闻
func NewsUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "NewsUpdate Controller")
	// 获取参数
	var param dto.NewsUpdateParam
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
	res, err := admin_service.UpdateNews(ctx, param)

	// 结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}

//前端需要个全部news的接口，此处有优化空间，考虑平台和后台一套
//这里只做适配
func NewsAll(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "NewsALL Controller")
	//获取参数
	var param dto.NewsAllParam
	param.PageSize = 10000
	param.CurrentPage = 0

	/*
		if err := ctx.ShouldBindJSON(&param); err != nil {
			logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
			util.FailJson(ctx, cerror.InvalidParams)
			return
		}
	*/

	if util.IsDebug() {
		fmt.Println("---->input param: ", param)
		logger.Infoc(ctx, "---->input param: %+v", param)
	}
	//参数校验

	//调用service
	var newsAllRes dto.NewsAllRes
	var err cerror.Cerror
	newsAllRes, err = service.NewsALL(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, newsAllRes)
	} else {
		util.FailJson(ctx, err)
	}

}

func NewsBannerAll(ctx *gin.Context) {

	// 获取参数

	// 参数校验

	// 调用service
	picNews, err := service.BannerNewsALL(ctx)
	fmt.Println(err)

	// 结果返回
	var res dto.BannerNewsLatestRes
	res.PicNews = picNews

	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}
}
