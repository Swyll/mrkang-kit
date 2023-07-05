package log

import "fmt"

func init() {
	SetLogger(&ConsoleLogger{})
}

// ConsoleLogger 控制台日志
type ConsoleLogger struct{}

// Debug debug level
func (l *ConsoleLogger) Debug(v ...interface{}) {
	fmt.Println("[\x1b[32mDEBU\x1b[0m]", v)
}

// Debugf debug level
func (l *ConsoleLogger) Debugf(format string, v ...interface{}) {
	fmt.Println("[\x1b[32mDEBU\x1b[0m]", fmt.Sprintf(format, v...))
}

// Debugln debug level
func (l *ConsoleLogger) Debugln(v ...interface{}) {
	fmt.Println("[\x1b[32mDEBU\x1b[0m]", v)
}

// Info info level
func (l *ConsoleLogger) Info(v ...interface{}) {
	fmt.Println("[\x1b[36mINFO\x1b[0m]", v)
}

// Infof info level
func (l *ConsoleLogger) Infof(format string, v ...interface{}) {
	fmt.Println("[\x1b[36mINFO\x1b[0m]", fmt.Sprintf(format, v...))
}

// Infoln info level
func (l *ConsoleLogger) Infoln(v ...interface{}) {
	fmt.Println("[\x1b[36mINFO\x1b[0m]", v)
}

// Warn warn level
func (l *ConsoleLogger) Warn(v ...interface{}) {
	fmt.Println("[\x1b[33mWARN\x1b[0m]", v)
}

// Warnf warn level
func (l *ConsoleLogger) Warnf(format string, v ...interface{}) {
	fmt.Println("[\x1b[33mWARN\x1b[0m]", fmt.Sprintf(format, v...))
}

// Warnln warn level
func (l *ConsoleLogger) Warnln(v ...interface{}) {
	fmt.Println("[\x1b[33mWARN\x1b[0m]", v)
}

// Error error level
func (l *ConsoleLogger) Error(v ...interface{}) {
	fmt.Println("[\x1b[31mERRO\x1b[0m]", v)
}

// Errorf error level
func (l *ConsoleLogger) Errorf(format string, v ...interface{}) {
	fmt.Println("[\x1b[31mERRO\x1b[0m]", fmt.Sprintf(format, v...))
}

// Errorln error level
func (l *ConsoleLogger) Errorln(v ...interface{}) {
	fmt.Println("[\x1b[31mERRO\x1b[0m]", v)
}

// Fatal fatal level
func (l *ConsoleLogger) Fatal(v ...interface{}) {
	fmt.Println("[\x1b[31mFATA\x1b[0m]", v)
}

// Fatalf fatal level
func (l *ConsoleLogger) Fatalf(format string, v ...interface{}) {
	fmt.Println("[\x1b[31mFATA\x1b[0m]", fmt.Sprintf(format, v...))
}

// Fatalln fatal level
func (l *ConsoleLogger) Fatalln(v ...interface{}) {
	fmt.Println("[\x1b[31mFATA\x1b[0m]", v)
}

func (l *ConsoleLogger) SetLogLevel(level string) {
}
