package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gormEngine *gorm.DB
var gormSyncOnce sync.Once

type gormFnOption func(*gorm.DB)

type gormDBSource struct {
	typz string
	host string
	port string
	name string
	user string
	pwds string
}

func GormInstance() *gorm.DB {
	return gormEngine
}

func GormInitialize(opts ...gormFnOption) {
	gormSyncOnce.Do(func() {
		// 配置信息
		var dnSource gormDBSource
		dnSource.typz = "mysql"
		dnSource.host = "127.0.0.1"
		dnSource.port = "3306"
		dnSource.name = "test"
		dnSource.user = "root"
		dnSource.pwds = "password"
		// 设置gorm
		var err error = nil
		gormEngine, err = gorm.Open(dnSource.typz, gormGetSourceStr(&dnSource))
		if err != nil {
			gormEngine = nil
			panic("failed to connect database")
		}
		// 服务配置
		gormConfigLogger()
		gormConfigConnect()
		for _, opt := range opts {
			opt(gormEngine)
		}
		// 注册销毁器
		registCloser(func() {
			if gormEngine != nil {
				gormEngine.Close()
				gormEngine = nil
			}
		})
	})
}

func gormConfigLogger() {
	if gormEngine == nil {
		return
	}
	gormEngine.SetLogger(Logger())
	switch devmode {
	case DebugMode: // 开启以展示详细的日志
		gormEngine.LogMode(true)
	case InforMode:
		gormEngine.LogMode(false)
	}
}

func gormConfigConnect() {
	if gormEngine == nil {
		return
	}
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	gormEngine.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns 设置数据库连接最大打开数。
	gormEngine.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置可重用连接的最长时间
	gormEngine.DB().SetConnMaxLifetime(time.Hour)

	// 转义struct表名字的时候后缀不用加上s
	gormEngine.SingularTable(true)
}

func gormGetSourceStr(dnSource *gormDBSource) string {
	var result string
	switch dnSource.typz {
	case "mysql":
		result = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dnSource.user, dnSource.pwds, dnSource.host, dnSource.port, dnSource.name)
	case "sqlite":
		result = ""
	}
	return result
}
