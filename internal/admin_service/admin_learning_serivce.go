package admin_service

import (
	"edu-imp/internal/admin_dto"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

//根据分类id获取推荐的资源
func GetLearningAll(ctx *gin.Context) ([]admin_dto.LearningRes, cerror.Cerror) {

	var res []admin_dto.LearningRes

	showNumber := 500 //最大500条记录
	learningList := dao.GetNlearning(ctx, showNumber)

	for _, val := range learningList {
		var learningRes admin_dto.LearningRes

		categoryRes, _ := dao.GetLearningCategoryByID(ctx, val.CategoryID)

		learningRes.Category = categoryRes.Title
		learningRes.ID = val.ID
		learningRes.Title = val.Title
		learningRes.Desc = val.Desc
		learningRes.Author = val.Author
		learningRes.PictureUrl = dao.GetResourceContentFromID(ctx, val.CoverPictureID)

		res = append(res, learningRes)
	}

	return res, nil
}

//获取资源详细信息
func GetLearningByID(ctx *gin.Context, id int) admin_dto.LearningDetailRes {

	var res admin_dto.LearningDetailRes

	getRes, _ := dao.GetLearning(ctx, id)
	res.ID = getRes.ID
	res.Desc = getRes.Desc
	res.Title = getRes.Title
	res.Author = getRes.Author
	res.PictureUrl = dao.GetResourceContentFromID(ctx, getRes.CoverPictureID)
	categoryRes, _ := dao.GetLearningCategoryByID(ctx, getRes.CategoryID)
	res.Category = categoryRes.Title

	learningRes, _ := dao.GetLearningResourceList(ctx, id)
	for _, val := range learningRes {
		var item admin_dto.ResourceRes
		item.ID = val.ID
		item.Title = val.Title
		item.Desc = val.Desc
		item.Index = val.Index
		item.Content = dao.GetResourceContentFromID(ctx, val.ResourceID)

		res.ResourceList = append(res.ResourceList, item)
	}

	return res
}

func GetLearningCategory(ctx *gin.Context) ([]dto.LearningCategory, cerror.Cerror) {

	return dao.GetLearningCategory(ctx)

}

func AddLearningCategory(ctx *gin.Context, param dto.LearningCategory) (int, cerror.Cerror) {

	return dao.AddLearningCategory(ctx, param)
}

func DelLearningCategory(ctx *gin.Context, id int) (int, cerror.Cerror) {

	return dao.DelLearningCategory(ctx, id)
}
