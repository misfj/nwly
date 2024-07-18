package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nwyl/config"
)

var (
	//GVA_DB     *gorm.DB
	//GVA_DBList map[string]*gorm.DB
	//GVA_REDIS  redis.UniversalClient
	//GVA_MONGO  *qmgo.QmgoClient
	//GVA_CONFIG config.Server
	//GVA_VP     *viper.Viper
	//// GVA_LOG    *oplogging.Logger
	//GVA_LOG                 *zap.Logger
	//GVA_Timer               timer.Timer = timer.NewTimerTask()
	//GVA_Concurrency_Control             = &singleflight.Group{}
	//GVA_ROUTERS             gin.RoutesInfo
	//BlackCache              local_cache.Cache
	//lock                    sync.RWMutex
	GobalDb *gorm.DB
	Config  *config.Config
	Logger  *zap.Logger
)

//LogMode(LogLevel) Interface
//Info(context.Context, string, ...interface{})
//Warn(context.Context, string, ...interface{})
//Error(context.Context, string, ...interface{})
//Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
