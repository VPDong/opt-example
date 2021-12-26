package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	DebugMode = 0
	InforMode = 1
)

var basedir string = ""
var devmode int = DebugMode
var logger = logrus.New()

func FirstToInitialize(baseDir string, devMode int) {
	basedir = strings.Trim(baseDir, " ")
	if !strings.HasSuffix(basedir, "/") {
		basedir += "/"
	}
	devmode = devMode
	if err := comConfigLogger(); err != nil {
		panic(err)
	}
}

func comConfigLogger() (err error) {
	file, err := os.OpenFile(path.Join(basedir, "goweb.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	switch devmode { // 设置默认输出和日志级别
	case DebugMode:
		logger.Out = io.MultiWriter(file, os.Stdout)
		logger.SetLevel(logrus.DebugLevel)
	case InforMode:
		logger.Out = file
		logger.SetLevel(logrus.InfoLevel)
	}
	return
}

func LastToFinalize() {
	if len(closers) > 0 {
		for _, closer := range closers {
			closer()
		}
	}
}

func Logger() *logrus.Logger {
	return logger
}

var closers []func()

// 注册当服务关闭时需要关闭的资源
func registCloser(closer func()) {
	closers = append(closers, closer)
}
