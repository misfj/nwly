package initilalize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nwyl/config"
	"nwyl/global"
	"nwyl/initilalize/internal"
	"os"
)

func GormMysqlByConfig(mysqlConfig *config.MySQLConfig) {
	var err error
	dsn := mysqlConfig.Username + ":" +
		mysqlConfig.Password + "@tcp(" + mysqlConfig.Path + ":" +
		mysqlConfig.Port + ")/" + mysqlConfig.DBName + "?" + mysqlConfig.Config
	cfg := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	//以后合并Logger,临时处理
	//tmpLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second,   // Slow SQL threshold
	//		LogLevel:                  logger.Silent, // Log level
	//		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	//		ParameterizedQueries:      true,          // Don't include params in the SQL log
	//		Colorful:                  false,         // Disable color
	//	},
	//)
	if global.GobalDb, err = gorm.Open(mysql.New(cfg), internal.Gorm.Config(mysqlConfig.Prefix, mysqlConfig.Singular)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		global.GobalDb.InstanceSet("gorm:table_options", "ENGINE="+mysqlConfig.Engine)
		sqlDB, _ := global.GobalDb.DB()
		sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	}
}
