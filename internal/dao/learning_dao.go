package dao

import (
	"edu-imp/internal/common"
	"edu-imp/internal/dto"
	"edu-imp/internal/model/mysql"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Learning struct {
	BaseModel
	Title          string `gorm:"column:title"`
	Desc           string `gorm:"column:desc"`
	Author         string `gorm:"column:author"`
	CategoryID     int    `gorm:"column:category_id"`
	CoverPictureID int    `gorm:"column:cover_picture_id"`
}

func (Learning) TableName() string {
	return "learning"
}

func AddLearning(ctx *gin.Context, param dto.Learning) (int, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	learning := Learning{Title: param.Title, Desc: param.Desc, CoverPictureID: param.CoverPictureID, CategoryID: param.CategoryID}
	result := mysqlDB.Create(&learning)

	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.CheckUser] fail 2,err=%+v", result.Error)
		return common.FailedID, cerror.NewCerror(common.Failed, result.Error.Error())
	}

	return learning.ID, nil
}

//获取n个学习资料，根据一定的推荐标准，比如：时间
func GetNlearning(ctx *gin.Context, number int) []Learning {
	mysqlDB := mysql.GetDB()

	var res []Learning

	result := mysqlDB.Order("updated_at desc").Limit(int(number)).Find(&res)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return nil
	}

	return res
}

//获取学习资料详情
func GetLearning(ctx *gin.Context, id int) dto.LearningRes {
	mysqlDB := mysql.GetDB()

	var learning Learning
	var res dto.LearningRes
	result := mysqlDB.Where("id=?", id).First(&learning)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return res
	}

	res.Desc = learning.Desc
	res.Title = learning.Title
	res.Author = learning.Author

	return res
}

//获取n个学习资料，根据一定的推荐标准，比如：时间
func GetNCategoryLearning(ctx *gin.Context, categoryID int, number int) []Learning {
	mysqlDB := mysql.GetDB()

	var res []Learning

	result := mysqlDB.Where("category_id=?", categoryID).Order("updated_at desc").Limit(int(number)).Find(&res)
	if result.Error != nil {
		logger.Warnc(ctx, "[userDao.NewsLatest] fail,err=%+v", result.Error)
		return nil
	}

	return res
}

//分页获取学习资源
func GetLearningCategoryResourcePagedData(ctx *gin.Context, pageParam dto.LearningCategoryParam) ([]dto.CategoryResourceItem, cerror.Cerror) {
	var res []dto.CategoryResourceItem

	mysqlDB := mysql.GetDB()

	ps := pageParam.PageSize
	offset := (pageParam.CurrentPage - 1) * (ps)

	var learningList []Learning
	result := mysqlDB.Where("category_id=?", pageParam.CategoryID).Limit(int(ps)).Offset(int(offset)).Find(&learningList)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsDao.NewsALL] fail,err=%+v, CurrentPage=%d", result.Error, pageParam.CurrentPage)
		return res, cerror.DbSelectError
	}

	for _, val := range learningList {
		var item dto.CategoryResourceItem

		item.Name = val.Title
		item.Desc = val.Desc
		item.LearningID = val.ID
		item.Author = val.Author
		item.PictureUrl = GetResourceContentFromID(ctx, val.CoverPictureID)

		res = append(res, item)
	}

	return res, nil
}

//获取新闻条数
func GetLearningCategoryResourceCount(ctx *gin.Context, id int) (int64, cerror.Cerror) {
	mysqlDB := mysql.GetDB()

	var learning []Learning
	var count int64

	result := mysqlDB.Where("category_id=?", id).Model(&learning).Count(&count)
	if result.Error != nil {
		logger.Warnc(ctx, "[newsDao.NewsALL] fail 2,err=%+v", result.Error)
		return 0, cerror.DbSelectError
	}

	return count, nil
}
