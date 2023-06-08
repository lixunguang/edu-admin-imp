package controller

import (
	"edu-imp/internal/dto"
	"edu-imp/internal/middleware"
	"edu-imp/internal/service"
	"edu-imp/pkg/logger"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//用户相关接口

type UserFake struct {
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

func GetIdByToken(tokenStr string) string {

	var userIDStr string
	token, err := jwt.ParseWithClaims(tokenStr, &middleware.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YundaoEdu"), nil
	})

	if claims, ok := token.Claims.(*middleware.CustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
		userIDStr = claims.UserID
	} else {
		fmt.Println(err)
	}

	return userIDStr

}

func GetUser(ctx *gin.Context) {
	logger.Infoc(ctx, "[%s] start ...", "GetUser Controller")
	//获取参数

	tokenStr := ctx.GetHeader("Authorization")

	loginId := GetIdByToken(tokenStr)

	//参数校验

	//调用service
	user, err := service.GetUser(ctx, loginId)

	//结果返回
	if err != nil {
		util.FailJson(ctx, err)
	} else {
		var fake UserFake //todo：remove THIS not used
		fake.Name = user.Name
		fake.LoginID = user.LoginID
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

func UserChangePassword(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func UserRegister(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func UserInfo(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func UserFavorite(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}

func UserHistory(ctx *gin.Context) {
	//get 解析
	fmt.Printf("%p", ctx.Request.URL.Query())
	util.SuccessJson(ctx, nil)

}
