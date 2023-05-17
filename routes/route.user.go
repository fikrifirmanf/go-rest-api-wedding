package route

import (
	createUser "github.com/fikrifirmanf/go-rest-api-wedding/controllers/user/create"
	getUser "github.com/fikrifirmanf/go-rest-api-wedding/controllers/user/get"
	userCreateHandler "github.com/fikrifirmanf/go-rest-api-wedding/handlers/user/create"
	userGetHandler "github.com/fikrifirmanf/go-rest-api-wedding/handlers/user/get"
	middleware "github.com/fikrifirmanf/go-rest-api-wedding/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoute(db *gorm.DB, route *gin.Engine) {
	getUsersRepository := getUser.NewRepository(db)
	getUsersService := getUser.NewService(getUsersRepository)
	getUsersHandler := userGetHandler.NewHandlerGetUsers(getUsersService)

	createUserRepository := createUser.NewRepositoryCreate(db)
	createUserService := createUser.NewServiceCreate(createUserRepository)
	createUserHandler := userCreateHandler.NewHandlerCreateUser(createUserService)

	groupRoute := route.Group("/api/v1")
	usersGroupRout := groupRoute.Group("/users").Use(middleware.Authentication())
	usersGroupRout.GET("", getUsersHandler.GetUsersHandler)
	usersGroupRout.POST("", createUserHandler.CreateUserHandler)

}
