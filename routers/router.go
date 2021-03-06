package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/zhangyusheng/framework/docs"

	"github.com/zhangyusheng/framework/routers/api"
	"github.com/zhangyusheng/framework/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/test", v1.Test)
	}

	return r
}
