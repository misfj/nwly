package core

import (
	"fmt"
	"github.com/spf13/viper"
	"nwyl/global"
	"os"
)

func Viper() {
	// 设置配置文件的名称（不带扩展名）
	viper.SetConfigName("config")
	// 添加配置文件所在的路径
	viper.AddConfigPath(".")
	// 设置配置文件类型
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		os.Exit(1)
	}
	// 解析配置到全局结构体
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Error unmarshaling config: %s\n", err)
		os.Exit(1)
	}
}
