/************************************************************
 * Author:        jackey
 * Date:        2022/10/19
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package controller

import (
	"gin/controller/v1"
	"gin/logs"
	"github.com/gin-gonic/gin"
)

type Default struct {
	apiV1.BaseController
}

func (d *Default) Get(c *gin.Context) {
	logs.Info("----", c.GetHeader("User-Agent"))
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
