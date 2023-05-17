package util

import "github.com/gin-gonic/gin"

type Responses struct {
	StatusCode int         `json:"status_code"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponses struct {
	StatusCode int         `json:"status_code"`
	Method     string      `json:"method"`
	Error      string      `json:"code"`
	Message    interface{} `json:"message"`
}

func CustomAPIResponse(ctx *gin.Context, statusCode int, method string, message string, data interface{}) {
	jsonResponse := Responses{
		StatusCode: statusCode,
		Method:     method,
		Message:    message,
		Data:       data,
	}

	if statusCode >= 400 {
		ctx.JSON(statusCode, jsonResponse)
		defer ctx.AbortWithStatus(statusCode)
	} else {
		ctx.JSON(statusCode, jsonResponse)
	}
}

func CustomAPIErrorResponse(ctx *gin.Context, statusCode int, method string, errorCode string, message []string) {
	jsonResponse := ErrorResponses{
		StatusCode: statusCode,
		Method:     method,
		Error:      errorCode,
		Message:    message,
	}

	ctx.JSON(statusCode, jsonResponse)
	defer ctx.AbortWithStatus(statusCode)
}
