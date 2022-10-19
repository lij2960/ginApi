/************************************************************
 * Author:        jackey
 * Date:        2022/10/18
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package main

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	configInit()
	if gin.Mode() == gin.DebugMode {
		// 启动时，重新生成文档描述
		_ = utils.ExecNoRes("swag init")
	}
}

func main() {

	r := Router()

	r.Run(":" + Config.Section(Env).Key("port").String())
}
