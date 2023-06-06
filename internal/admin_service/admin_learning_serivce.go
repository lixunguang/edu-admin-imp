package admin_service

import (
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"github.com/gin-gonic/gin"
)

func GetLearningCategory(ctx *gin.Context) ([]dto.LearningCategory, cerror.Cerror) {

	return dao.GetLearningCategory(ctx)

}

func AddLearningCategory(ctx *gin.Context, param dto.LearningCategory) (int, cerror.Cerror) {

	return dao.AddLearningCategory(ctx, param)
}

func DelLearningCategory(ctx *gin.Context, id int) (int, cerror.Cerror) {

	return dao.DelLearningCategory(ctx, id)
}
