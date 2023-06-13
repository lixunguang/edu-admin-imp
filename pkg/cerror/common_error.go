package cerror

var (
	InvalidParams = NewCerror(100001, "参数错误")
	//	NotFound                  = NewCerror(100002, "找不到")
	//	UnauthorizedAuthNotExist  = NewCerror(100003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	//	UnauthorizedTokenError    = NewCerror(100004, "鉴权失败，Token 错误")
	//	UnauthorizedTokenTimeout  = NewCerror(100005, "鉴权失败，Token 超时")
	//	UnauthorizedTokenGenerate = NewCerror(100006, "鉴权失败，Token 生成失败")
	//	ErrorFindUser             = NewCerror(100007, "错误，未找到用户")
	//	ErrorPassword             = NewCerror(100008, "错误，用户名/密码错误")
	//	ErrortTokenExpired        = NewCerror(100009, "错误，token过期")

	ErrorLoginSucc    = NewCerror(10000, "登录成功") //成功登录
	ErrorLoginFailed  = NewCerror(10001, "登录失败")
	ErrorLoginAgain   = NewCerror(10002, "登录成功：再次登录")
	ErrorUserNotLogin = NewCerror(10003, "用户未登录")

	ErrorTokenFormat = NewCerror(10011, "token格式错误")
	ErrorTokenExpire = NewCerror(10012, "token过期")
	ErrorTokenEmpty  = NewCerror(10013, "token为空")

	ErrorUserExist    = NewCerror(10021, "用户已经存在")
	ErrorUserNotExist = NewCerror(10022, "用户不存在")
	ErrorPassword     = NewCerror(10023, "密码错误")

	ErrorDataOperator = NewCerror(10030, "数据库操作失败")

	ErrorRecordExist          = NewCerror(100056, "操作不允许:记录已经存在")
	ErrorAddCourseFailed      = NewCerror(100057, "增加课程失败")
	ErrorAddLessonContentFind = NewCerror(100058, "增加章节内容失败：未找到章节内容信息")
)
