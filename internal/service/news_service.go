package service

import (
	"edu-imp/config"
	"edu-imp/internal/dao"
	"edu-imp/internal/dto"
	"edu-imp/pkg/cerror"
	"edu-imp/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//获取前NewsNumber条新闻,NewsNumber为配置项
func NewsLatest(ctx *gin.Context) ([]dto.NewsResObj, cerror.Cerror) {

	newsNumberStr := fmt.Sprintf("%d", config.Vipper.Get("News.NewsNumber"))
	newsNumber, _ := strconv.ParseInt(newsNumberStr, 10, 32)
	dtoNews := dao.GetTitleNews(ctx, int(newsNumber))

	return dtoNews, nil
}

//获取前NewsNumber条新闻,NewsNumber为配置项
func BannerNewsLatest(ctx *gin.Context) ([]dto.PicNews, cerror.Cerror) {

	pictureNewsNumberStr := fmt.Sprintf("%d", config.Vipper.Get("News.PictureNewsNumber"))
	pictureNewsNumber, _ := strconv.ParseInt(pictureNewsNumberStr, 10, 32)
	dtoPicNews := dao.GetPictureNews(ctx, int(pictureNewsNumber))

	return dtoPicNews, nil
}

func BannerNewsALL(ctx *gin.Context) ([]dto.PicNews, cerror.Cerror) {

	pictureNewsNumber := -1
	dtoPicNews := dao.GetPictureNews(ctx, pictureNewsNumber)
	fmt.Println(len(dtoPicNews), cap(dtoPicNews))

	return dtoPicNews, nil
}

func OneNews(ctx *gin.Context, id dto.IDParam) (dto.NewsResObj, cerror.Cerror) {
	return dao.GetNewsById(ctx, id)
}

//获取新闻列表
func NewsALL(ctx *gin.Context, param dto.NewsAllParam) (dto.NewsAllRes, cerror.Cerror) {

	news, err := dao.GetNewsPagedData(ctx, param)
	count, _ := dao.GetNewsCount(ctx)

	//pageSize := fmt.Sprintf("%d", config.Vipper.Get("News.PageSize"))
	//ps, _ := strconv.ParseInt(pageSize, 10, 64)

	//var page common.Page
	//page = page.CreatePageInfo(param.CurrentPage, ps, count)

	var newsList []dto.NewsItemRes
	for i, val := range news {
		fmt.Println(i, val)
		var news dto.NewsItemRes
		news.NewsID = val.ID
		news.Title = val.Title
		news.Author = val.Publisher

		//news.DateStr = val.UpdatedAt.Format("2006-01-02 15:04:05") //todo:简要处理下时间格式 2023-03-15T17:39:40+08:00  --》2023-03-15 17:39:40 08:00
		news.DateStr = val.UpdatedAt.Format(util.FormatDate) //todo:简要处理下时间格式 2023-03-15T17:39:40+08:00  --》2023-03-15 17:39:40 08:00

		news.PictureID = val.PictureID

		newsList = append(newsList, news)

	}

	var newsAllRes dto.NewsAllRes
	newsAllRes.TotalNumber = count
	//newsAllRes.CurrentPage = param.CurrentPage
	//newsAllRes.Pages = page.Nums
	//newsAllRes.PageSize = ps
	newsAllRes.NewsList = newsList

	return newsAllRes, err

}
