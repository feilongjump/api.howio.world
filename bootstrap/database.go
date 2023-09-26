package bootstrap

import (
	"time"

	"github.com/feilongjump/api.howio.world/internal/database"
)

func SetupDatabase() {

	db := database.ConnectDB()

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大生存时间，以确保连接可以被驱动安全关闭。官方建议小于5分钟。
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	// 设置闲置连接的最大存在时间, support>=go1.15
	sqlDB.SetConnMaxIdleTime(time.Minute * 3)
}
