package cerror

var (
	InvalidParams             = NewCerror(100001, "参数错误")
	NotFound                  = NewCerror(100002, "找不到")
	UnauthorizedAuthNotExist  = NewCerror(100003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewCerror(100004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewCerror(100005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewCerror(100006, "鉴权失败，Token 生成失败")
	ErrorFindUser             = NewCerror(100007, "错误，未找到用户")
	ErrorPassword             = NewCerror(100008, "错误，用户名/密码错误")
	ErrortTokenExpired        = NewCerror(100009, "错误，token过期")
)
