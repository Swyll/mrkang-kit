package log

var std Logger

//Init 初始化方法
func Init(logger Logger) {
	std = logger
}

//Debug debug等级写入
func Debug(v ...interface{}) {
	std.Debug(v...)
}

//Debugf debug拼接写入
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

//Info info等级写入
func Info(v ...interface{}) {
	std.Info(v...)
}

//Infof info拼接写入
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

//Warn warn等级写入
func Warn(v ...interface{}) {
	std.Warn(v...)
}

//Warnf warn拼接写入
func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

//Error error等级写入
func Error(v ...interface{}) {
	std.Error(v...)
}

//Errorf error拼接写入
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

//Fatal fatal等级写入
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

//Fatalf fatal拼接写入
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}
