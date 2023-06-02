package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	BaseModel

	LoginID     int    `gorm:"column:login_id"`
	Name        string `gorm:"column:name"`
	Password    string `gorm:"column:password"`
	Role        int    `gorm:"column:role"`
	PhoneNumber string `gorm:"column:phone_number"`
	CreatorID   int    `gorm:"column:creator_id"`
	IsDeleted   int8   `gorm:"column:is_deleted"`
}

func (Admin) TableName() string {
	return "admin"
}

func AddAdmin(ctx *gin.Context, param dto.Admin) (dto.AdminRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	admin := Admin{Name: param.Name, Password: param.Password}
	result := mysqlDB.Create(&admin)

	var res dto.AdminRes

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	res.Name = admin.Name
	res.ID = admin.ID

	return res, nil
}

func GetAdmin(ctx *gin.Context, param dto.AdminParam) ([]Admin, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var admin Admin
	admin.Name = param.Name

	var adminArray []Admin
	result := mysqlDB.Where(&admin).Find(&adminArray)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return adminArray, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return adminArray, nil

}

func DelAdmin(ctx *gin.Context, param dto.AdminParam) ([]Admin, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var admin Admin
	admin.Name = param.Name

	var adminArray []Admin
	result := mysqlDB.Where(&admin).Find(&adminArray)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return adminArray, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	var ids []int
	if len(adminArray) > 0 {

		for _, v := range adminArray {
			ids = append(ids, v.ID)
		}

		result := mysqlDB.Delete(&Admin{}, ids)

		if result.Error != nil {
			logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
			return adminArray, cerror.NewCerror(common.FailedID, result.Error.Error())
		}
	}

	return adminArray, nil
}

// 校验用户名及密码
func CheckAdmin(ctx *gin.Context, userName string, password string) (int, cerror.Cerror) {

	db := mysql.GetDB()

	// query
	u := Admin{}

	result := db.Where("name = ? ", userName).First(&u)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail,err=%+v, loginId=%d", result.Error, userName)
		return common.ErrorFindUser, cerror.DbSelectError
	}

	if u.Name == "" { // 用户不存在
		return common.ErrorFindUser, nil
	}

	result = db.Where("password = ?", password).First(&u)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, loginId=%d", result.Error, userName)
		return common.ErrorPassword, cerror.DbSelectError
	}

	if u.Name == "" { // 密码错误
		return common.ErrorPassword, nil
	}

	return common.CheckOK, nil
}
