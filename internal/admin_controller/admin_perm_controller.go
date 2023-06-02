package admin_controller

import (
	"edu-imp/internal/admin_service"
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/middleware"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// 权限相关接口

// gwt，token串
// token存储，压缩，内存，数据库
// 重复登录，存储新的token，之前token失效。
// 过期检查

func AdminLogin(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "adminLogin Controller")
	// 获取参数
	var param dto.AdminLoginParam
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
	tokenStr, res := admin_service.Login(ctx, param.Name, param.Password)

	// 结果返回
	//todo:处理下结果返回，从service层返回err
	if res == common.SuccessLoginAgain || res == common.SuccessLogin {
		var loginRes dto.LoginRes
		loginRes.Token = tokenStr
		loginRes.FirstLogin = false

		util.SuccessJson(ctx, loginRes)

	} else if res == common.ErrorFindUser {
		util.FailJson(ctx, cerror.ErrorFindUser)
	} else if res == common.ErrorPassword {
		util.FailJson(ctx, cerror.ErrorPassword)
	}

}

// /edu/v1/perm/logout
func AdminLogout(ctx *gin.Context) {

	//获取参数
	tokenStr := ctx.GetHeader("Authorization")

	//参数校验

	//调用service
	res := admin_service.Logout(ctx, tokenStr)

	//结果返回

	if res == common.Success {
		util.SuccessJson(ctx, nil)
	} else {
		util.FailJson(ctx, cerror.InvalidParams)
	}

}

type UserFake struct {
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	Avatar           string `json:"avatar"`
	Email            string `json:"email"`
	Job              string `json:"job"`
	JobName          string `json:"jobName"`
	Organization     string `json:"organization"`
	OrganizationName string `json:"organizationName"`
	Location         string `json:"location"`
	LocationName     string `json:"locationName"`
	Introduction     string `json:"introduction"`
	PersonalWebsite  string `json:"personalWebsite"`
	Phone            string `json:"phone"`
	RegistrationDate string `json:"registrationDate"`
	AccountId        string `json:"accountId"`
	Certification    int    `json:"certification"`

	Role int `json:"role"`
	dto.UserRes
}

func GetIdByToken(tokenStr string) string { //
	//middleware.CustomClaims{}
	var userIDStr string
	token, err := jwt.ParseWithClaims(tokenStr, &middleware.CustomClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YundaoEdu"), nil
	})

	if claims, ok := token.Claims.(*middleware.CustomClaimsAdmin); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserName, claims.StandardClaims.ExpiresAt)
		userIDStr = claims.UserName
	} else {
		fmt.Println(err)
	}

	return userIDStr

}

func GetUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "admin GetUser Controller")
	//获取参数

	tokenStr := ctx.GetHeader("Authorization")

	loginId := GetIdByToken(tokenStr)

	//参数校验

	var param dto.AdminParam
	param.Name = loginId
	//调用service
	user, err := admin_service.GetAdmin(ctx, param)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		var fake UserFake //todo：remove THIS not used
		if len(user) > 0 {
			fake.Name = user[0].Name
			fake.LoginId = user[0].Name
		}

		fake.UserID = user[0].ID
		fake.OrganizationID = 1 //todo:for test

		fake.Role = 1 // todo:role需要填上值
		fake.Avatar = "//lf1-xgcdn-tos.pstatp.com/obj/vcloud/vadmin/start.8e0e4855ee346a46ccff8ff3e24db27b.png"
		fake.Email = "wangliqun@email.com"
		fake.Job = "前端艺术家"
		fake.Organization = "Frontend"
		fake.OrganizationName = "前端"
		fake.Location = "beijing"
		fake.LocationName = "北京"
		fake.Introduction = "人潇洒，性温存"
		fake.PersonalWebsite = "https://www.arco.design"
		fake.Phone = "150****0000"
		fake.RegistrationDate = "2013-05-10 12:10:00"
		fake.AccountId = "15012312300"
		fake.Certification = 1

		util.SuccessJson(ctx, fake)
	}

}

func Copyright(ctx *gin.Context) {
	var copyright string
	copyright = "<pre><span style=\"color: #ced4d9; font-size: 12pt;\">Copyright 2014 - 2023 IBE edu. All Rights Reserved</span></pre>"

	util.SuccessJson(ctx, copyright)
}
