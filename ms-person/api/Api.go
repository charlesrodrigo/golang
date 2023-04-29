package api

import (
	"context"
	"os"

	"br.com.charlesrodrigo/ms-person/api/controllers"
	"br.com.charlesrodrigo/ms-person/api/docs"
	"br.com.charlesrodrigo/ms-person/api/handlers"
	"br.com.charlesrodrigo/ms-person/helper/constants"
	"br.com.charlesrodrigo/ms-person/helper/logger"
	"br.com.charlesrodrigo/ms-person/helper/tracer"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

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
	router.Use(requestid.New())
	router.Use(handlers.AddRequestIdInRequestContext())
	router.Use(otelgin.Middleware(appName))
	router.Use(handlers.AddLogRequestAndResponse())
	router.Use(ginzap.RecoveryWithZap(configLogger, true))
	router.Use(handlers.TimeoutMiddleware())

	docs.SwaggerInfo.BasePath = "/"
	controllers.NewApiPersonController(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	logger.Info("Started Server! -> http://localhost:8080")
	logger.Info("Swagger! -> http://localhost:8080/swagger/index.html")

	router.Run(":8080")

}
