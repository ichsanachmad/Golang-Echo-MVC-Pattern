package utils

import (
	"Golang-Echo-MVC-Pattern/constant"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	_, errFile := os.Stat(constant.PathLog)
	if os.IsNotExist(errFile) {
		errDir := os.MkdirAll(constant.PathToLog, 0777)
		if errDir != nil {
			LogError(errFile, DetailFunction())
		}

	}
	log.Out = os.Stdout
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(constant.PathLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err == nil {
		log.Out = file
	} else {
		log.Info(constant.LogMessageErrorLoadFile)
	}
}

// LOG ERROR TO FILE LOGRUS.LOG
func LogError(err error, name string) {
	log.WithFields(logrus.Fields{
		constant.DetailLog: name,
	}).Error(err.Error())
}

// GET SPECIFIC DETAIL FUNCTION CALLED
func DetailFunction() string {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprint(constant.LogCalled, file, constant.FromLine, line, constant.FromFunction, runtime.FuncForPC(pc).Name())
	}
	return ""
}
