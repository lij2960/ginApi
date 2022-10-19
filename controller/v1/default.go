/************************************************************
 * Author:        jackey
 * Date:        2022/10/19
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package apiV1

import (
	"gin/logs"
	"github.com/gin-gonic/gin"
)

type Default struct {
	BaseController
}

// Get default
// @Summary get 测试
// @Schemes
// @Description just for test 11111
// @Tags        example
// @Accept      json
// @Produce     json
// @Success     200 {string} Helloworld
// @Router      / [get]
func (d *Default) Get(c *gin.Context) {
	logs.Info("Test", d.Test)
	c.JSON(200, gin.H{
		"version": d.Test,
	})
}
