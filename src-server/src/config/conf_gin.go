package config

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var ginEngine *gin.Engine
var ginSyncOnce sync.Once

type ginFnOption func(*gin.Engine)

func GinInstance() *gin.Engine {
	return ginEngine
}

// 最后调用，run会阻塞
func GinInitializeFinal(opts ...ginFnOption) {
	ginSyncOnce.Do(func() {
		// 设置gin
		ginEngine = gin.Default()
		// 服务配置
		ginConfigLogger()
		for _, opt := range opts {
			opt(ginEngine)
		}
		// 启动服务
		ginEngine.Run(":8080")
	})
}

func ginConfigLogger() {
	if ginEngine == nil {
		return
	}
	gin.DefaultWriter = Logger().Out
	switch devmode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	case InforMode:
		gin.SetMode(gin.ReleaseMode)
	}
	// 日志中间件
	ginEngine.Use(func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next() // 处理请求(关键)
		endTime := time.Now()
		runTime := endTime.Sub(startTime)

		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		Logger().WithFields(logrus.Fields{
			"status_code": statusCode,
			"run_time":    runTime,
			"client_ip":   clientIP,
			"req_method":  reqMethod,
			"req_uri":     reqUri,
		}).Info()
	})
}
