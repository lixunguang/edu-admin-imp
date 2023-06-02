package util

import (
	"edu-imp/pkg/cerror"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode    = 0
	SuccessMessage = "success"
	FailCode       = -1
	HttpStatus     = http.StatusOK
	HeaderData     = "_header_bind_data"
	MaxCharNum     = 255
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewSuccessHttpResponse(data interface{}) *HttpResponse {
	return &HttpResponse{
		Code:    SuccessCode,
		Data:    data,
		Message: SuccessMessage,
	}
}

func NewFailHttpResponse(err cerror.Cerror) *HttpResponse {
	if err == nil {
		return &HttpResponse{
			Code:    FailCode,
			Data:    nil,
			Message: "error message is nil",
		}
	}
	return &HttpResponse{
		Data:    nil,
		Code:    err.Code(),
		Message: err.Error(),
	}
}

func RecordNotFoundHttpResponse() *HttpResponse {
	return &HttpResponse{
		Code:    0,
		Data:    nil,
		Message: "record not found",
	}

}

func NewFailMessageHttpResponse(err string) *HttpResponse {
	return &HttpResponse{
		Code:    FailCode,
		Data:    nil,
		Message: err,
	}
}

func SuccessJson(c *gin.Context, data interface{}) {
	if IsDebug() {
		str := fmt.Sprintf("<----return success: %+v", data)
		if len(str) >= MaxCharNum {
			str = str[:MaxCharNum]
			str = str + " ... "
		}
		str = str + "---->"
		fmt.Println(str)
		//	logger.Infoc(c,"<----return success: %+v", data)
	}

	c.JSON(
		HttpStatus,
		NewSuccessHttpResponse(data),
	)

}

func FailJson(c *gin.Context, err cerror.Cerror) {
	if IsDebug() {
		fmt.Println("<----return fail: ", err)
		//	logger.Infoc(c,"<----return success: %+v", err)
	}

	c.JSON(
		HttpStatus,
		NewFailHttpResponse(err),
	)
}

func RecordNotFoundJson(c *gin.Context) {
	if IsDebug() {
		fmt.Println("<----return Record Not Found: ")
		//	logger.Infoc(c,"<----return success: %+v", err)
	}

	c.JSON(
		HttpStatus,
		RecordNotFoundHttpResponse(),
	)
}
