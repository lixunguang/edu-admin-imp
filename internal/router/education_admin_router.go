package router

import (
	"edu-imp/internal/admin_controller"
	"edu-imp/internal/controller"
	"github.com/gin-gonic/gin"
)

//教育管理平台后台接口
func educationAdminRouter(e *gin.Engine) {
	g := e.Group("/admin")

	//权限 perm
	g.POST("/perm/login", admin_controller.AdminLogin)
	//if !util.IsDebug() {
	//g.Use(middleware.CheckAdminAuth)
	//}

	g.POST("/perm/logout", admin_controller.AdminLogout)
	g.POST("/copyright", admin_controller.Copyright)

	// 用户 增加 删除 查找
	g.POST("/user/add", admin_controller.AddUser)
	g.POST("/user/del", admin_controller.DelUser)
	g.POST("/user", admin_controller.GetUser) //后台api与平台api共用一个

	g.POST("/teacher/add", admin_controller.AddTeacher) //增加教师
	g.POST("/teacher", admin_controller.AddTeacher)     //获取教师列表

	// 管理员 增加 删除 查找
	g.POST("/administrator/add", admin_controller.AddAdmin)
	g.POST("/administrator/del", admin_controller.DelAdmin)
	g.POST("/administrator", admin_controller.GetAdmin)

	//新闻编辑
	g.POST("/news", controller.News)
	g.POST("/news/all", admin_controller.NewsAll)
	g.POST("/news/banner/all", admin_controller.NewsBannerAll)
	g.POST("/news/add", admin_controller.NewsAdd)
	g.POST("/news/del", admin_controller.NewsDel)
	g.POST("/news/update", admin_controller.NewsUpdate)

	//学习资源
	g.POST("/learning/category/all", admin_controller.LearningCategory)
	g.POST("/learning/category/add", admin_controller.LearningCategoryAdd)
	g.POST("/learning/category/del", admin_controller.LearningCategoryDel)

	g.POST("/learning", admin_controller.Learning)
	g.POST("/learning/all", admin_controller.LearningAll)
	g.POST("/learning/add", admin_controller.LearningAdd)
	g.POST("/learning/del", admin_controller.LearningDel)
	g.POST("/learning/update", admin_controller.LearningUpdate)

	g.POST("/learning/resource", admin_controller.LearningResource)
	g.POST("/learning/resource/add", admin_controller.LearningResourceAdd)
	g.POST("/learning/resource/del", admin_controller.LearningResourceDel)

	// 课程编辑：课程概要信息
	g.POST("/course/all", admin_controller.CourseAll)       //查：所有课程
	g.POST("/course", admin_controller.Course)              //查：课程详情
	g.POST("/course/del", admin_controller.CourseDel)       //删
	g.POST("/course/add", admin_controller.CourseAdd)       //增
	g.POST("/course/update", admin_controller.CourseUpdate) //改

	// 课程编辑：章节概要信息
	g.POST("/course/lesson/all", admin_controller.CourseLessonAll)       //查：所有章节
	g.POST("/course/lesson", admin_controller.CourseLesson)              //查：章节详情
	g.POST("/course/lesson/del", admin_controller.CourseLessonDel)       //删：
	g.POST("/course/lesson/add", admin_controller.CourseLessonAdd)       //增
	g.POST("/course/lesson/update", admin_controller.CourseLessonUpdate) //改

	// 章节编辑：内容
	g.POST("/lesson/content", admin_controller.LessonContent)
	g.POST("/lesson/content/add", admin_controller.LessonContentAdd)
	g.POST("/lesson/content/del", admin_controller.LessonContentDel)
	// 章节编辑：参考
	g.POST("/lesson/refer", admin_controller.LessonRefer)
	g.POST("/lesson/refer/add", admin_controller.LessonReferAdd)
	g.POST("/lesson/refer/del", admin_controller.LessonReferDel)
	// 章节编辑：实验
	g.POST("/lesson/experiment", admin_controller.LessonExperiment)
	g.POST("/lesson/experiment/add", admin_controller.LessonExperimentAdd)
	g.POST("/lesson/experiment/del", admin_controller.LessonExperimentDel)
	// 章节编辑：作业
	g.POST("/lesson/work", admin_controller.LessonWork)
	g.POST("/lesson/work/add", admin_controller.LessonWorkAdd)
	g.POST("/lesson/work/del", admin_controller.LessonWorkDel)

	g.POST("/lesson/work/requirement", admin_controller.LessonWorkRequirement)
	g.POST("/lesson/work/requirement/update", admin_controller.LessonWorkRequirementUpdate)

	g.POST("/upload_files", controller.UploadFiles)
	g.POST("/richtext/add", admin_controller.AddRichtext)

	//以下接口不需要对外暴露
	g.POST("/richtext/del", admin_controller.DelRichtext)
	g.POST("/richtext/update", admin_controller.UpdateRichtext)
	g.POST("/file/del", admin_controller.DelFile)

}
