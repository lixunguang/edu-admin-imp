package service

import (
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context, user dto.User) (dto.AddUserRes, cerror.Cerror) {

	return dao.AddUser(ctx, user)

}

func UserChangePassword(ctx *gin.Context) {

}

func UserRegister(ctx *gin.Context) {

}

func GetUser(ctx *gin.Context, loginID string) (dto.UserRes, cerror.Cerror) {
	return dao.GetUser(ctx, loginID)
}
