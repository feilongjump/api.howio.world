package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	"time"
)

// Viper 实例
var Viper *viper.Viper

// StrMap 简写 —— map[string]interface{}
type StrMap map[string]interface{}

// init 函数在 import 的时候立刻被加载
func init() {
	// 初始化 Viper 库
	Viper = viper.New()

	// 设置文件名称
	Viper.SetConfigName(".env")

	// 配置类型，支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	Viper.SetConfigType("env")

	// 环境变量配置文件查找的路径，相对于 main.go
	Viper.AddConfigPath(".")

	// 开始读根目录下的 .env 文件，读不到会报错
	err := Viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 设置环境变量前缀，用以区分 Go 的系统环境变量
	Viper.SetEnvPrefix("appenv")
	// Viper.Get() 时，优先读取环境变量
	Viper.AutomaticEnv()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}

	return Get(envName)
}

// Add 新增配置项
func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

// Get 获取配置项，允许使用点式获取，如：app.name
func Get(path string, defaultValue ...interface{}) interface{} {
	// 不存在的情况下
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	return Viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetDuration 获取 time.Duration 类型的配置信息
func GetDuration(path string, defaultValue ...interface{}) time.Duration {
	return cast.ToDuration(Get(path, defaultValue...))
}
