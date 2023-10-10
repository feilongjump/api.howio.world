package config

import (
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {

	// 初始化 Viper 库
	config = viper.New()
	// 配置类型
	config.SetConfigType("toml")
	// 环境变量配置文件查找的路径，相对于 main.go
	config.AddConfigPath(".")
	// 设置环境变量前缀，用以区分 Go 的系统环境变量
	config.SetEnvPrefix("appenv")
	// 读取环境变量（支持 flags）
	config.AutomaticEnv()
}

// InitConfig 初始化环境配置文件
func InitConfig(envSuffix string) {

	// 默认加载 env.local.toml 文件，如果有传参 --env=name 的话，加载 env.name.toml 文件
	envPath := "env.local.toml"
	if len(envSuffix) > 0 {
		filepath := "env." + envSuffix + ".toml"
		if _, err := os.Stat(filepath); err == nil {
			// 如 env.dev.toml 或 env.prod.toml
			envPath = filepath
		}
	}

	// 加载环境配置
	config.SetConfigName(envPath)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控环境配置文件，变更时重新加载
	config.WatchConfig()

}

func GetString(key string) string {
	return config.GetString(key)
}

func GetInt(key string) int {
	return config.GetInt(key)
}

func GetBool(key string) bool {
	return config.GetBool(key)
}
