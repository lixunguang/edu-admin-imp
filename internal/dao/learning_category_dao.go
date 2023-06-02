package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type LearningCategory struct {
	BaseModel
	Title string `gorm:"column:title"`
	Desc  string `gorm:"column:desc"`
}

func (LearningCategory) TableName() string {
	return "learning_category"
}

func AddLearningCategory(ctx *gin.Context, param dto.LearningCategory) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	learningCategory := LearningCategory{Title: param.Title, Desc: param.Desc}
	result := mysqlDB.Create(&learningCategory)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return learningCategory.ID, nil
}

func GetLearningCategory(ctx *gin.Context) ([]dto.LearningCategory, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var res []dto.LearningCategory

	var learningCategory []LearningCategory
	result := mysqlDB.Select("id", "title", "desc").Find(&learningCategory)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return res, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	for _, val := range learningCategory {
		var item dto.LearningCategory
		item.Title = val.Title
		item.Desc = val.Desc
		item.ID = val.ID
		res = append(res, item)
	}

	return res, nil
}
