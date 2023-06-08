package admin_service

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/common"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

func AddTeacher(ctx *gin.Context, param admin_dto.Teacher) (dto.AdminRes, cerror.Cerror) {

	adminArray, _ := dao.GetTeacherByName(ctx, param.Name)

	if len(adminArray) == 0 {
		var teacher dao.Teacher
		teacher.Name = param.Name
		teacher.LoginID = param.Name
		teacher.Password = param.Password
		teacher.OrganizationID = param.OrganizationID
		teacher.Introduce = param.Introduce

		return dao.AddTeacher(ctx, teacher)
	}

	var res dto.AdminRes
	return res, common.ErrorUserExist
}

func DelTeacher(ctx *gin.Context, name string) (string, cerror.Cerror) {

	return dao.DelTeacher(ctx, name)
}

func AllTeacher(ctx *gin.Context) ([]admin_dto.Teacher, cerror.Cerror) {

	var res []admin_dto.Teacher

	allRes, err := dao.AllTeacher(ctx)

	for _, val := range allRes {
		var item admin_dto.Teacher
		item.Introduce = val.Introduce
		item.OrganizationID = val.OrganizationID
		item.Name = val.Name
		item.Password = val.Password

		res = append(res, item)

	}

	return res, err
}

func UpdateTeacher(ctx *gin.Context, param admin_dto.UpdateTeacherParam) (string, cerror.Cerror) {

	return dao.UpdateTeacher(ctx, param)

}
