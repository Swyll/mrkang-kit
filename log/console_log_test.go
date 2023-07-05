package log

import (
	"testing"
)

func TestConsoleLog(t *testing.T) {
	Debug("test")
	Debugf("%v", "test")
	Debugln("test")

	Info("test")
	Infof("%v", "test")
	Infoln("test")

	Warn("test")
	Warnf("%v", "test")
	Warnln("test")

	Error("test")
	Errorf("%v", "test")
	Errorln("test")

	Fatal("test")
	Fatalf("%v", "test")
	Fatalln("test")

}
