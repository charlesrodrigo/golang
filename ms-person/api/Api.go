package api

import (
	"time"

	"br.com.charlesrodrigo/ms-person/api/controllers"
	"br.com.charlesrodrigo/ms-person/api/docs"
	"br.com.charlesrodrigo/ms-person/helper/logger"

	"br.com.charlesrodrigo/ms-person/api/handlers"

	"br.com.charlesrodrigo/ms-person/infra/database"
	"br.com.charlesrodrigo/ms-person/internal/service"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Crud Person API
// @version         1.0
// @description     This is a crud of person.

// @contact.name   Charles Rodrigo
// @contact.email  charlesrodrigo@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          http://localhost:8080/swagger/index.html
func StartServerApi() {
	configLogger := logger.Init()

	personController := initDependenciesPersonController()

	router := gin.Default()
	router.Use(ginzap.Ginzap(configLogger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(configLogger, true))
	router.Use(handlers.TimeoutMiddleware())

	docs.SwaggerInfo.BasePath = "/"
	v1 := router.Group("/api/v1")
	{
		personGroup := v1.Group("/person")
		{
			personGroup.POST("/", personController.CreatePerson)
			personGroup.PUT("/:id", personController.UpdatePerson)
			personGroup.GET("/:id", personController.GetPerson)
			personGroup.GET("/", personController.GetAllPerson)
			personGroup.DELETE("/:id", personController.DeletePerson)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	logger.Info("Started Server! -> http://localhost:8080 ")
	logger.Info("Swagger! -> http://localhost:8080/swagger/index.html")

	router.Run(":8080")

}

func initDependenciesPersonController() controllers.PersonController {
	personRepository := database.NewPersonRepositoryImpl()
	personService := service.NewPersonServiceImpl(personRepository)
	personController := controllers.NewPersonController(personService)
	return personController
}
