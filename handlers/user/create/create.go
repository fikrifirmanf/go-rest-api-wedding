package userCreateHandler

import (
	"net/http"

	createUser "github.com/fikrifirmanf/go-rest-api-wedding/controllers/user/create"
	util "github.com/fikrifirmanf/go-rest-api-wedding/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service createUser.Service
}

func NewHandlerCreateUser(service createUser.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUserHandler(ctx *gin.Context) {
	var input createUser.CreateUser
	ctx.ShouldBindJSON(&input)

	// config := gpc.ErrorConfig()

	_, errCreateUser, errorMessage := h.service.CreateServiceUser(&input)
	switch errCreateUser {
	case "CREATE_USER_FAILED":
		util.CustomAPIErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errCreateUser, errorMessage)
		return
	case "ACCOUNT_ALREADY_EXISTS":
		util.CustomAPIErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errCreateUser, errorMessage)
	case "VALIDATION_ERROR":
		util.CustomAPIErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errCreateUser, errorMessage)
	default:
		util.CustomAPIResponse(ctx, http.StatusCreated, http.MethodPost, "Create user success", nil)
	}
}
