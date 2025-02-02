package routes

import (
	"car-collection-manager/controllers"
	_ "car-collection-manager/controllers"
	docs "car-collection-manager/docs"
	"car-collection-manager/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/gen"
)

var db = make(map[string]string)

func generateSwaggerDocs() {
	fmt.Println("generateSwaggerDocs Loaded")

	generator := gen.New()

	err := generator.Build(&gen.Config{
		SearchDir:  "./",
		OutputDir:  "./docs",
		ParseDepth: 2,
		Excludes:   ".git,vendor",
	})

	if err != nil {
		log.Fatalf("Failed to generate Swagger documentation: %v", err)
	}
	log.Println("Swagger docs generated successfully!")
}

func SetupRouter(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Car Collection Manager API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// generateSwaggerDocs()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	HealthRouter(router)
	AdminRouter(router)
	CarRouter(router)
	UserRouter(router)

	router.Run(":" + utils.GetEnvOrDefault("PORT", "8080"))
}

func HealthRouter(router *gin.Engine) {
	router.GET("/health", controllers.Health)
	router.GET("/ping", controllers.PingPongHandler)
}

func CarRouter(router *gin.Engine) {
	router.GET("/car", controllers.GetCar)
	router.POST("/car", controllers.PostCar)
}

func AdminRouter(router *gin.Engine) {
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(ctx *gin.Context) {
		user := ctx.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if ctx.Bind(&json) == nil {
			db[user] = json.Value
			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
}

func UserRouter(router *gin.Engine) {
	// Get user value
	router.GET("/user/:name", func(ctx *gin.Context) {
		user := ctx.Params.ByName("name")
		db["abc"] = "ok"

		value, ok := db[user]
		if ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})
}
