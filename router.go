/************************************************************
 * Author:        jackey
 * Date:        2022/10/19
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package main

import (
	"gin/controller"
	apiV1 "gin/controller/v1"
	"gin/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/", new(controller.Default).Get)

	// @BasePath /v1
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			con := new(apiV1.Default)
			con.Prepare(c)
			con.Get(c)
		})
	}

	// 如果是开发模式，生成文档地址
	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.BasePath = "/v1"
		urlSwagger := ginSwagger.URL("/swagger/doc.json")
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlSwagger))
	}

	return router
}
