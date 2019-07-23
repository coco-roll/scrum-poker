package routers

import (
	"github.com/gin-gonic/gin"
	"scrum-poker/app/api"
	"scrum-poker/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api")
	{
		apiv1.GET("/test", api.GetTest)
		apiv1.POST("/test", api.AddTest)
		apiv1.PUT("/test/:id", api.EditTest)
		apiv1.DELETE("/test/:id", api.DeleteTest)
		apiv1.GET("/url", api.GetUrl)
		apiv1.GET("/poker", api.Poker)
		apiv1.GET("/setck", api.SetCk)
	}

	return r
}
