package log

var std Logger

// SetLogger 设置日志
func SetLogger(logger Logger) {
	std = logger
}

// Debug debug level
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Debugf debug level
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

//// Debugln debug level
//func Debugln(v ...interface{}) {
//	std.Debugln(v...)
//}

// Info info level
func Info(v ...interface{}) {
	std.Info(v...)
}

// Infof info level
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

//// Infoln info level
//func Infoln(v ...interface{}) {
//	std.Infoln(v...)
//}

// Warn warn level
func Warn(v ...interface{}) {
	std.Warn(v...)
}

// Warnf warn level
func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

//// Warnln warn level
//func Warnln(v ...interface{}) {
//	std.Warnln(v...)
//}

// Error error level
func Error(v ...interface{}) {
	std.Error(v...)
}

// Errorf error level
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

//// Errorln error level
//func Errorln(v ...interface{}) {
//	std.Errorln(v...)
//}

// Fatal fatal level
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Fatalf fatal level
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

//// Fatalln fatal level
//func Fatalln(v ...interface{}) {
//	std.Fatalln(v...)
//}
