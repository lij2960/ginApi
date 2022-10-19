/************************************************************
 * Author:        jackey
 * Date:        2022/10/19
 * Description: 基础controller
 * Version:    V1.0.0
 **********************************************************/

package apiV1

import (
	"gin/logs"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	Test *string
}

// Prepare 前置操作
func (b *BaseController) Prepare(c *gin.Context) {
	logs.Info("BaseController-Prepare", "done")
	str := "测试内容"
	b.Test = &str
}
