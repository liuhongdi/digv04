package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/liuhongdi/digv04/global"
	"github.com/liuhongdi/digv04/router"
	"log"
)

//init
func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = global.SetupDBLink()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	//设置运行模式
	gin.SetMode(global.ServerSetting.RunMode)
	//引入路由
	r := router.Router()
	//run
	r.Run(":"+global.ServerSetting.HttpPort)
}




