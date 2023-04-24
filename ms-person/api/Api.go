package api

import (
	"context"
	"os"
	"time"

	"br.com.charlesrodrigo/ms-person/api/controllers"
	"br.com.charlesrodrigo/ms-person/api/docs"
	"br.com.charlesrodrigo/ms-person/helper/constants"
	"br.com.charlesrodrigo/ms-person/helper/logger"
	"br.com.charlesrodrigo/ms-person/helper/tracer"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"br.com.charlesrodrigo/ms-person/api/handlers"

	"br.com.charlesrodrigo/ms-person/infra/database"
	"br.com.charlesrodrigo/ms-person/internal/service"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	appName := os.Getenv(constants.GET_SERVICE_NAME)
	ctx := context.Background()
	cleanup := tracer.Init()
	defer cleanup(ctx)

	configLogger := logger.Init(appName)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(otelgin.Middleware(appName))
	router.Use(ginzap.GinzapWithConfig(configLogger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		TraceID:    true,
		Context: ginzap.Fn(func(c *gin.Context) (fields []zapcore.Field) {
			if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
				fields = append(fields, zap.String("x-request-id", requestID))
			}
			return
		}),
	}))
	router.Use(ginzap.RecoveryWithZap(configLogger, true))

	router.Use(handlers.TimeoutMiddleware())
	router.Use(requestid.New())

	personController := initDependenciesPersonController(ctx)
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
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	logger.Info("Started Server! -> http://localhost:8080 ")
	logger.Info("Swagger! -> http://localhost:8080/swagger/index.html")

	router.Run(":8080")

}

func initDependenciesPersonController(ctx context.Context) controllers.PersonController {
	personRepository := database.NewPersonRepositoryImpl(ctx)
	personService := service.NewPersonServiceImpl(ctx, personRepository)
	personController := controllers.NewPersonController(personService)
	return personController
}
