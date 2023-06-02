package admin_service

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

func AddTeacher(ctx *gin.Context, name string) (dto.AdminRes, cerror.Cerror) {

	adminArray, _ := dao.GetTeacherByName(ctx, name)

	if len(adminArray) == 0 {
		return dao.AddTeacher(ctx, name)
	}

	var res dto.AdminRes
	return res, common.ErrorUserExist
}
