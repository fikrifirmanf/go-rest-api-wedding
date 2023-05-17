package middleware

import (
	"net/http"

	util "github.com/fikrifirmanf/go-rest-api-wedding/utils"
	"github.com/gin-gonic/gin"
)

type Unauthorized struct {
	Status  int    `json:"status_code"`
	Method  string `json:"method"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var errorResponse Unauthorized

		errorResponse.Status = http.StatusForbidden
		errorResponse.Method = ctx.Request.Method
		errorResponse.Code = "FORBIDDEN_ACCESS"
		errorResponse.Message = "Authorization is required to access this resource"

		if ctx.Request.Header.Get("Authorization") == "" {
			ctx.JSON(http.StatusForbidden, errorResponse)
			ctx.Abort()
			return
		}

		token, err := util.VerifyHeaderToken(ctx, "JWT_SECRET")

		errorResponse.Status = http.StatusUnauthorized
		errorResponse.Method = ctx.Request.Method
		errorResponse.Code = "UNAUTHORIZED_ACCESS"
		errorResponse.Message = "Invalid token or token is expired"

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			ctx.Abort()
			return
		} else {
			ctx.Set("user", token.Claims)
			ctx.Next()
		}

	}
}
