package bootstrap

import (
	"github.com/feilongjump/api.howio.world/internal/config"
)

// 初始化配置文件信息
func SetupConfig() {
	config.InitConfig("local")
}
