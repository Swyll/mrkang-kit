package log

// Logger 日志接口
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	//Debugln(v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})
	//Infoln(v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	//Warnln(v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	//Errorln(v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	//Fatalln(v ...interface{})

	SetLogLevel(string)
}
