package route

import (
	createUser "github.com/fikrifirmanf/go-rest-api-wedding/controllers/user-controllers/create"
	getUser "github.com/fikrifirmanf/go-rest-api-wedding/controllers/user-controllers/get"
	userCreateHandler "github.com/fikrifirmanf/go-rest-api-wedding/handlers/user-handlers/create"
	userGetHandler "github.com/fikrifirmanf/go-rest-api-wedding/handlers/user-handlers/get"
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
	usersGroupRout := groupRoute.Group("/users")
	usersGroupRout.GET("", getUsersHandler.GetUsersHandler)
	usersGroupRout.POST("", createUserHandler.CreateUserHandler)

}
