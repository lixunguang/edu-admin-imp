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

// 新闻相关接口

// 为首页定制接口，用于返回首页显示的消息。
func BannerNewsLatest(ctx *gin.Context) {

	// 获取参数

	// 参数校验

	// 调用service
	picNews, err := service.BannerNewsLatest(ctx)
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

//获取新闻消息
func News(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "News Controller")
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
	news, err := service.OneNews(ctx, param)

	//结果返回
	if err == nil {
		var newsRes dto.NewsRes
		newsRes.News = news
		util.SuccessJson(ctx, newsRes)

	} else {
		util.FailJson(ctx, err)
	}

}
