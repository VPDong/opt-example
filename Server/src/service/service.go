package service

import (
	"server/src/config"

	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	return config.Logger()
}
