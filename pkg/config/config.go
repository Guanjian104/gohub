package config

import (
	viperlib "github.com/spf13/viper"
)

// viper 库实例
var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 初始化 Viper 库
	viper = viperlib.New()

	viper.SetConfigType("env")

	// 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(".")

	// 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")

	// 读取环境变量
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 加载环境变量
	loadEnv(env)
	
	// 注册配置信息
	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	
}