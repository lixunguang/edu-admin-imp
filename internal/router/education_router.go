package router

import (
	"edu-imp/internal/controller"
	"edu-imp/internal/middleware"
	"edu-imp/pkg/util"
	"github.com/gin-gonic/gin"
)

//教育管理平台后台接口

func educationRouter(e *gin.Engine) {
	g := e.Group("/edu/v1")

	//权限 perm
	// /edu/v1/perm/login
	g.POST("/perm/login", controller.Login)
	if !util.IsDebug() {
		g.Use(middleware.CheckAuth)
	}
	// /edu/v1/perm/logout
	g.POST("/perm/logout", controller.Logout)

	//资源 resource
	g.POST("/resource/add", controller.ResourceAdd)

	//新闻栏目 news

	g.POST("/news/latest", controller.NewsLatest)
	g.POST("/news/banner/latest", controller.BannerNewsLatest)
	g.POST("/news", controller.News)
	g.POST("/news/all", controller.NewsALL)

	//学习栏目 learning
	g.POST("/learning/add", controller.LearningAdd)
	g.POST("/learning/category/add", controller.LearningCategoryAdd)
	g.POST("/learning/item/add", controller.LearningItemAdd)

	g.POST("/learning/resource/recommend", controller.LearningResourceRecommend)
	g.POST("/learning/resource", controller.LearningResource)
	g.POST("/learning/category", controller.LearningCategory)
	g.POST("/learning/category/resource", controller.LearningCategoryResource)
	g.POST("/learning/category/resource/recommend", controller.LearningCategoryResourceRecommend)

	//课程栏目 course
	g.POST("/course_resource/add", controller.CourseResourceAdd)
	g.POST("/course_teacher/add", controller.CourseTeacherAdd)

	g.POST("/course/overview", controller.CourseOverview)
	g.POST("/course/recommend", controller.CourseRecommend)
	g.POST("/course/all", controller.CourseAll)

	//章节栏目 Lesson
	//g.POST("/course/lesson/add", controller.CourseLessonAdd)
	//g.POST("/course/lesson/section/add", controller.CourseLessonSectionAdd)
	//g.POST("/course/lesson_resource/add", controller.CourseLessonResourceAdd)
	//g.POST("/course/lesson/work/add", controller.LessonWorkAdd)

	g.POST("/course/lesson/overview", controller.CourseLessonOverview)
	g.POST("/course/lesson/content", controller.CourseLessonContent)
	g.POST("/course/lesson/refer", controller.CourseLessonRefer)
	g.POST("/course/lesson/experiment", controller.CourseLessonExperiment)
	g.POST("/course/lesson/work", controller.CourseLessonWork)

	//提交作业 commit work
	g.POST("/course/lesson/work/commit/add", controller.LessonWorkCommitAdd)
	g.POST("/course/lesson/work/commit", controller.LessonWorkCommit)
	//g.POST("/course/lesson/work/commit/file/delete", controller.LessonWorkCommitDel)
	//g.POST("/course/lesson/work/commit/message/update", controller.LessonWorkCommitMessageUpdate)

	g.POST("/upload_file", controller.UploadFile) //未使用
	g.POST("/upload_files", controller.UploadFiles)

	//打分score
	// /edu/v1/course/lesson/work/scroe_table
	g.POST("/course/lesson/work/scroe_table", controller.CourseLessonWorkScroeTable)
	// /edu/v1/course/lesson/work/mark
	g.POST("/course/lesson/work/mark", controller.CourseLessonWorkMark)

	//个人 user
	// /edu/v1/user/change_password
	g.POST("/user/change_password", controller.UserChangePassword)
	// /edu/v1/user/register
	g.POST("/user/register", controller.UserRegister)
	// /edu/v1/user/info
	g.POST("/user/info", controller.UserInfo)
	// /edu/v1/user/history
	g.POST("/user/history", controller.UserHistory)
	// /edu/v1/user/favorite
	g.POST("/user/favorite", controller.UserFavorite)

	g.POST("/user", controller.GetUser) //后台api与平台api共用一个
	//school
	// /edu/v1/school/dept
	g.POST("/school/dept", controller.SchoolDept)
	// /edu/v1/school/dept/class
	g.POST("/school/dept/class", controller.SchoolDeptClass)
	// /edu/v1/school/dept/class/student
	g.POST("/school/dept/class/student", controller.SchoolDeptClassStudent)

}
