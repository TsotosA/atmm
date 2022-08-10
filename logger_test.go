package main

import (
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestGetLogger(t *testing.T) {
	//ConfigInit()
	logger, _ := InitLogger("testLog.log")
	undo := zap.ReplaceGlobals(logger)
	t.Cleanup(func() {
		//err := logger.Sync()
		//if err != nil {
		//	zap.S().Errorf("%s", "failed to flush logs before exiting")
		//}
		undo()
	})
	if reflect.TypeOf(logger).Kind() != reflect.TypeOf(&zap.Logger{}).Kind() {
		t.Errorf("expected a pointer")
	}
}
