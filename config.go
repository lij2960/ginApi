/************************************************************
 * Author:        jackey
 * Date:        2022/10/18
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package main

import (
	"gin/logs"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
)

var (
	Config *ini.File
	Env    string
)

func configInit() {
	var err error
	Config, err = ini.Load("config/config.ini")
	if err != nil {
		logs.Error("ini.Load", err)
		return
	}
	var configFile []interface{}
	for _, val := range Config.Section("conFile").Keys() {
		fileName := Config.Section("conFile").Key(val.Name()).String()
		configFile = append(configFile, "config/"+fileName)
	}
	logs.Info("---", configFile)
	Config, err = ini.Load("config/config.ini", configFile...)
	if err != nil {
		logs.Error("ini.Load", err)
		return
	}

	Env = Config.Section("").Key("mode").String()
	// 设置gin的运行模式
	ginMode(Env)
	logs.Info("env", Env)
	logs.Info("port", Config.Section(Env).Key("port").String())
	logs.Info("test", Config.Section("").Key("test").String())
}

func ginMode(env string) {
	// 读取系统变量的运行模式，优先采用系统变量的运行模式
	sysEnv := os.Getenv("ginMode")
	logs.Info("sysEnv", sysEnv)
	if sysEnv != "" {
		env = sysEnv
	}
	if env == "test" {
		gin.SetMode(gin.TestMode)
	} else if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
