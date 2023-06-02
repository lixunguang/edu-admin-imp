package controller

import (
	"edu-imp/internal/dto"
	"edu-imp/internal/service"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

//提交作业接口

//注：布置的作业接口属于lesson_controller(课堂接口)大类

//获取已经提交作业
func LessonWorkCommit(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LessonWorkCommit Controller")
	//获取参数
	var param dto.WorkCommitParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	//参数校验

	//调用service
	res, err := service.WorkCommit(ctx, param)
	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res) //这里返回有问题。。。。。。如果返回一个空对象。。。。。。。。。。。。。。。。。。前端怎么处理？？
	} else {
		util.FailJson(ctx, err)
	}
}

//提交作业
func LessonWorkCommitAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LessonWorkCommitAdd Controller")
	//获取参数
	var param dto.WorkCommitAddParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	//参数校验

	//调用service
	id, err2 := service.WorkCommitAdd(ctx, param)
	//结果返回
	if err2 == nil {
		util.SuccessJson(ctx, id)
	} else {
		util.FailJson(ctx, err2)
	}
}

/*
//删除作业
func LessonWorkCommitDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LessonWorkCommitDel Controller")
	//获取参数
	var param dto.WorkCommitDelParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	//参数校验

	//调用service
	//删除时，我们只需要资源的id，是否还校验lessonid 和userid呢
	id, err2 := service.WorkCommitDel(ctx, param)

	var res dto.IDRes = dto.IDRes{ID: id}

	//结果返回
	if err2 == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, nil)
	}
}

//更新留言信息
func LessonWorkCommitMessageUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "LessonWorkCommitMessageUpdate Controller")
	//获取参数
	var param dto.WorkCommitDelParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		logger.Errorc(ctx, "[%s] bind params fail,err=%+v", "method", err)
		util.FailJson(ctx, cerror.InvalidParams)
		return
	}
	//参数校验

	//调用service
	//删除时，我们只需要资源的id，是否还校验lessonid 和userid呢
	id, err2 := service.WorkCommitDel(ctx, param)

	var res dto.IDRes = dto.IDRes{ID: id}

	//结果返回
	if err2 == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, nil)
	}
}
*/
