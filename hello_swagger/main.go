package main

import (
	
	"github.com/gin-gonic/gin"
	docs "example/hello_swagger/docs"
	ginSwagger "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/gin-swagger"

	
	"net/http"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} HelloWorld
// @Router /example/helloworld [get]
func HelloWorld(g *gin.Context){
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", HelloWorld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}