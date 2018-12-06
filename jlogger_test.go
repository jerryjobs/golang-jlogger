package jlogger_test

import (
	"testing"

	"jerrysd.cn/jlogger"
)

func TestWriteLog(t *testing.T) {

	t.Log(" ok start ...")
	jLog := jlogger.New("t1")
	jLog.Info("ok")
	jLog.Debug("test debug", "asb")
}
