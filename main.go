package main

import (
	"fmt"
	"nwyl/core"
	"nwyl/global"
)

func main() {
	core.Viper()
	core.Zap()
	global.Logger.Info("MySQL configuration")
	//core.GormMysqlByConfig(&global.Config.MySQL)
	fmt.Println(global.Config)
	core.Start()
}
