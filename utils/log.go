package utils

import (
	"go.uber.org/zap"
	"runtime"
	"reflect"
)

var (
	Log *zap.SugaredLogger
)

func NewLog() {
	logger, _ := zap.NewProduction()
	Log = logger.Sugar()
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}


func LogTemplate() string {
	return "[%s] : %s"
}

func ErrTemplate() string {
	return "[%s] with error msg : %s"
}