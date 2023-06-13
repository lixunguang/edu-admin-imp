// Package common provides commonly used functions and utilities.
package common

import (
	"edu-imp/config"
	"edu-imp/pkg/cerror"
	"fmt"
)

var SigKey = []byte("signature_hmac_secret_shared_key")

const (
	Failed  = -1
	Success = 0
)

const (
	FailedID = -1 //-1为不合法的id，一般为插入失败,最初为0  //todo：重构
)

//错误处理
var (
	ErrorOK           = cerror.NewCerror(10000, "ok")
	ErrorTokenFormat  = cerror.NewCerror(10001, "token格式错误")
	ErrorTokenExpire  = cerror.NewCerror(10002, "token过期")
	ErrorTokenEmpty   = cerror.NewCerror(10003, "token为空")
	ErrorUserExist    = cerror.NewCerror(10004, "用户已经存在")
	ErrorUserNotExist = cerror.NewCerror(10005, "用户不存在")
	ErrorPassword     = cerror.NewCerror(10006, "密码错误")
	ErrorLoginAgain   = cerror.NewCerror(10007, "再次登录")
	ErrorUserNotLogin = cerror.NewCerror(10005, "用户未登录")

	ErrorRecordExist          = cerror.NewCerror(100016, "操作不允许:记录已经存在")
	ErrorAddCourseFailed      = cerror.NewCerror(100017, "增加课程失败")
	ErrorAddLessonContentFind = cerror.NewCerror(100018, "增加章节内容失败：未找到章节内容信息")
)

const (
	NotExpired       = 0 //是否过期
	TokenExpired     = 1
	TokenFormatError = 2 //
)

const (
	PictureType = 1  //图片资源类型
	WordType    = 2  //word文档资源类型
	PdfType     = 3  //pdf资源类型
	AppType     = 4  //app资源类型
	IbeType     = 5  //ibe资源类型
	MessageType = 6  //同7,未使用
	RichType    = 7  //富文本资源类型
	VedioType   = 8  //视频资源类型
	FileType    = 23 //除了MessageType和RichType之外的所有类型
)

//课程资源子类型
const (
	CoursePicture = 1 //课程缩略图
	CourseVedio   = 2 //课程介绍视频
	CourseLesson  = 3 //课程章节
)

//章节子类型
const (
	LessonContent    = 1
	LessonRefer      = 2
	LessonExperiment = 3
	LessonWork       = 4
)

//作业资源类型
const (
	WorkRequireRichText   = 41 //作业要求：富文本格式要求
	WorkRequireAttachment = 42 //作业要求：文件附件格式的要求
	WorkCommitAttachment  = 43 //作业提交的文件
	WorkCommitText        = 44 //作业提交附言
)

const (
	StudentRole = 1 //人员类型；学生1、教师2、管理员3
	TeacherRole = 2
	AdminRole   = 3
)

// GetCDNAddr is a function that returns the CDN address.
func GetCDNAddr() string {
	addr := fmt.Sprintf("%s", config.Vipper.Get("CDNServer.Addr"))

	return addr
}
