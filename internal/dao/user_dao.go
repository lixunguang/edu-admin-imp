package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type User struct {
	BaseModel
	LoginID string `gorm:"column:login_id"`
	// 学号 ,注意区分ID（程序内部使用）和LoginID（对外接口使用），通过GetIDByLoginID和GetLoginIDByID转换 一个为int类型，一个为string类型！！！！
	Name           string `gorm:"column:name"`            // 姓名
	OrganizationID int    `gorm:"column:organization_id"` // 所属院系
	Password       string `gorm:"column:password"`
	PhoneNumber    string `gorm:"column:phone_number"`
	CreatorID      int    `gorm:"column:creator_id"` // 创建本用户的管理员id
	IsDeleted      int8   `gorm:"column:is_deleted"`
}

func (User) TableName() string {
	return "user"
}

// 增加用户
func AddUser(ctx *gin.Context, param dto.User) (dto.AddUserRes, cerror.Cerror) {

	mysqlDB := mysql.GetDB()

	var res dto.AddUserRes

	getResult, err := GetUser(ctx, param.LoginID)
	if err != nil { // 数据库操作失败
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err")
		return res, err
	}

	if getResult.LoginId != "" { // 重复的用户
		return res, common.ErrorUserExist
	}

	user := User{LoginID: param.LoginID, Name: param.Name, Password: param.Password, OrganizationID: param.OrganizationID}
	result := mysqlDB.Create(&user)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	res.ID = user.ID
	res.Name = user.Name

	return res, nil
}

func GetUser(ctx *gin.Context, loginID string) (dto.UserRes, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var user User
	result := mysqlDB.Where("login_id = ? ", loginID).First(&user)

	var userRes dto.UserRes
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return userRes, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	userRes.Name = user.Name
	userRes.OrganizationID = user.OrganizationID
	userRes.LoginId = user.LoginID

	return userRes, nil
}

// 校验用户名及密码
func CheckUser(ctx *gin.Context, loginId string, password string) (int, cerror.Cerror) {

	db := mysql.GetDB()

	// query
	u := User{}

	result := db.Where("login_id = ? ", loginId).First(&u)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail,err=%+v, loginId=%d", result.Error, loginId)
		return common.ErrorFindUser, cerror.DbSelectError
	}

	if u.Name == "" { // 用户不存在
		return common.ErrorFindUser, nil
	}

	result = db.Where("password = ?", password).First(&u)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, loginId=%d", result.Error, loginId)
		return common.ErrorPassword, cerror.DbSelectError
	}

	if u.Name == "" { // 密码错误
		return common.ErrorPassword, nil
	}

	return common.CheckOK, nil
}

// 根据userid获取username
func GetUserNameByUserID(ctx *gin.Context, userID []int) []string {
	mysqlDB := mysql.GetDB()

	var users []User
	var userNames []string
	result := mysqlDB.Where("id IN ? ", userID).Find(&users) // 有多个的话
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, ", result.Error)
		return userNames
	}

	for _, val := range users {
		userNames = append(userNames, val.Name)
	}
	return userNames
}

// 根据id获取loginid
func GetLoginIDByID(ctx *gin.Context, userID int) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()
	var id int = 0
	var user User
	user.ID = userID

	result := mysqlDB.First(&user)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, ", result.Error)
		return id, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	return id, nil
}

// 根据loginid获取id
func GetIDByLoginID(ctx *gin.Context, loginID string) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()
	var id int = 0
	var user User

	result := mysqlDB.Where("login_id=?", loginID).First(&user)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v, ", result.Error)
		return id, cerror.NewCerror(common.FailedID, result.Error.Error())
	}

	id = user.ID
	return id, nil
}
