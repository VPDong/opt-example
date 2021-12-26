package main

import (
	"server/src/config"
	"server/src/router"
	"server/src/structs"
)

func main() {
	basedir := "/Users/mine/work/dev/GoWeb/log"
	// 首先初始化通用配置
	config.FirstToInitialize(basedir, config.InforMode)
	// 初始化gorm数据库
	config.GormInitialize(structs.DBInitialize)
	// 最后初始化gin框架
	config.GinInitializeFinal(router.Initialize)
	// 最后的最后清理资源
	defer config.LastToFinalize()
}
