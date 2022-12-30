package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"
	config "github.com/fikrifirmanf/go-rest-api-wedding/configs"
	route "github.com/fikrifirmanf/go-rest-api-wedding/routes"
	util "github.com/fikrifirmanf/go-rest-api-wedding/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	log.Fatal(router.Run(":" + util.GoDotEnv("PORT")))
}

func SetupRouter() *gin.Engine {
	db := config.DBConnection()
	router := gin.Default()

	if util.GoDotEnv("ENVIRONMENT") != "production" && util.GoDotEnv("ENVIRONMENT") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GoDotEnv("ENVIRONMENT") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	// init Routes
	route.InitUserRoute(db, router)

	return router
}
