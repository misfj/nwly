package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"nwyl/core/internal"
	"nwyl/global"
	"nwyl/utils"
	"os"
)

// Zap 获取 zap.Logger
func Zap() {
	if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
	}
	levels := global.Config.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	global.Logger = zap.New(zapcore.NewTee(cores...))
	if global.Config.Zap.ShowLine {
		global.Logger = global.Logger.WithOptions(zap.AddCaller())
	}
}
