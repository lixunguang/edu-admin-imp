package admin_controller

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/admin_service"
	"edu-imp/internal/common"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

//获取课程详情
func CourseLesson(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin CourseLesson Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
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
	res, err := admin_service.CourseLesson(ctx, param.CourseID, param.LessonID)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {

		util.SuccessJson(ctx, res)
	}
}

func CourseLessonUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin CourseLessonUpdate Controller")
	//获取参数
	var param admin_dto.UpdateLessonParam
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
	res, err := admin_service.LessonUpdate(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		util.SuccessJson(ctx, res)
	}
}

//根据lessonid 获取课程内容
func LessonContent(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonContent Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
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
	res, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonContentAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonContentAdd Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
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
	res, err := admin_service.LessonContentAdd(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonContentDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonContentDel Controller")
	//获取参数
	var param admin_dto.LessonContentParam
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
	res, err := admin_service.LessonContentDel(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

//根据lessonid 获取课程参考资源
func LessonRefer(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonRefer Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
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
	sectionType = common.LessonRefer //内容类型
	res, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonReferAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonReferAdd Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
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
	res, err := admin_service.LessonReferAdd(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonReferDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonReferDel Controller")
	//获取参数
	var param admin_dto.LessonResourceDelParam
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
	res, err := admin_service.LessonReferDel(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonExperiment(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonExperiment Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
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
	sectionType = common.LessonExperiment //内容类型
	res, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonExperimentAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonExperimentAdd Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
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
	res, err := admin_service.LessonExperimentAdd(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonExperimentDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonExperimentDel Controller")
	//获取参数
	var param admin_dto.LessonResourceDelParam
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
	res, err := admin_service.LessonExperimentDel(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonWork(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonExperiment Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
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
	sectionType = common.LessonWork //内容类型
	ret, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	var res []admin_dto.LessonSectionContent
	for _, val := range ret {
		if val.Index != common.WorkRequireRichText {
			res = append(res, val)
		}
	}

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonWorkAdd(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonWorkAdd Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
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
	res, err := admin_service.LessonWorkAdd(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonWorkDel(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonExperimentDel Controller")
	//获取参数
	var param admin_dto.LessonResourceDelParam
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
	res, err := admin_service.LessonWorkDel(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonWorkRequirement(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonWorkRequirement Controller")
	//获取参数
	var param admin_dto.CourseLessonIDParam
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
	sectionType = common.LessonWork //内容类型

	resource, err := admin_service.GetCourseLessonSectionResource(ctx, param.LessonID, sectionType)

	var res admin_dto.LessonSectionContent
	for _, val := range resource {
		if val.Index == common.WorkRequireRichText {
			res = val
		}
	}
	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}

func LessonWorkRequirementUpdate(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin LessonWorkRequirementUpdate Controller")
	//获取参数
	var param admin_dto.LessonResourceAddParam
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
	param.ResourceIndex = common.WorkRequireRichText
	res, err := admin_service.LessonWorkRequirementUpdate(ctx, param)

	//结果返回
	if err == nil {
		util.SuccessJson(ctx, res)
	} else {
		util.FailJson(ctx, err)
	}

}
