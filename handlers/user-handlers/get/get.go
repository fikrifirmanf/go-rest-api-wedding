package userGetHandler

import (
	"fmt"
	"net/http"

	getUser "github.com/fikrifirmanf/go-rest-api-wedding/controllers/user-controllers/get"
	util "github.com/fikrifirmanf/go-rest-api-wedding/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service getUser.Service
}

func NewHandlerGetUsers(service getUser.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetUsersHandler(ctx *gin.Context) {
	result, errorResult := h.service.GetServiceUsers()
	fmt.Println(result, errorResult)
	switch errorResult {
	case "USERS_NOT_FOUND":
		util.CustomAPIResponse(ctx, http.StatusNotFound, http.MethodGet, "Users not exist", nil)
	default:
		util.CustomAPIResponse(ctx, http.StatusOK, http.MethodGet, "Success get users", result)
	}

}
