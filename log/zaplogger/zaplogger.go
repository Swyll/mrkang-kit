package zaplogger

import (
	"os"
	"strings"
	"time"

	"github.com/Swyll/mrkang-kit/log"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// OptFunc 可选日志配置
type OptFunc func(*conf)

type conf struct {
	filePath      string        // 日志文件路径
	level         zapcore.Level // 日志级别
	maxSize       int           // 每个日志文件保存的最大尺寸 单位：M
	maxBackups    int           // 日志文件最多保存多少个备份
	maxAge        int           // 文件最多保存多少天
	compress      bool          // 是否压缩
	enableConsole bool          // 是否打印到控制台
}

type LoggerAgent struct {
	level *zap.AtomicLevel
	*zap.SugaredLogger
}

func (la *LoggerAgent) SetLogLevel(level string) {
	la.level.SetLevel(GetLogLevel(level))
}

// GetLogLevel 从字符串获取日志等级
func GetLogLevel(level string) zapcore.Level {
	v := strings.ToLower(level)
	switch v {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "faltal":
		return zapcore.FatalLevel
	default:
		return zapcore.WarnLevel
	}
}

// Logger 日志配置
func Logger(opts ...OptFunc) log.Logger {
	c := &conf{
		filePath:      "log/app.log",
		level:         1,
		maxSize:       5,
		maxBackups:    5,
		maxAge:        365,
		compress:      false,
		enableConsole: false,
	}
	for _, opt := range opts {
		opt(c)
	}

	hook := &lumberjack.Logger{
		Filename:   c.filePath, //filePath
		MaxSize:    c.maxSize,  // megabytes
		MaxBackups: c.maxBackups,
		MaxAge:     c.maxAge,   //days
		Compress:   c.compress, // disabled by default
	}
	defer hook.Close()
	/*zap 的 Config 非常的繁琐也非常强大，可以控制打印 log 的所有细节，因此对于我们开发者是友好的，有利于二次封装。
	  但是对于初学者则是噩梦。因此 zap 提供了一整套的易用配置，大部分的姿势都可以通过一句代码生成需要的配置。
	*/
	enConfig := zap.NewProductionEncoderConfig() //生成配置

	// 时间格式
	//enConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	enConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	enConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	w := zapcore.AddSync(hook)
	level := new(zap.AtomicLevel)
	level.SetLevel(c.level)

	allCore := []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(enConfig), //编码器配置
			w,                                   //打印到文件
			level,                               //日志等级
		),
	}

	if c.enableConsole {
		allCore = append(allCore, zapcore.NewCore(
			zapcore.NewConsoleEncoder(enConfig), //编码器配置
			os.Stdout,                           //打印到控制台
			level,                               //日志等级
		))
	}

	core := zapcore.NewTee(allCore...)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &LoggerAgent{
		SugaredLogger: logger.Sugar(),
		level:         level,
	}
}

// WithLogFile 指定日志文件
func WithLogFile(filePath string) OptFunc {
	return func(c *conf) {
		c.filePath = filePath
	}
}

// WithLevel 指定日志等级
func WithLevel(l zapcore.Level) OptFunc {
	return func(c *conf) {
		c.level = l
	}
}

// WithMaxSize 指定最大日志大小
func WithMaxSize(maxSize int) OptFunc {
	return func(c *conf) {
		c.maxSize = maxSize
	}
}

// WithMaxBackups 指定最大日志备份
func WithMaxBackups(maxBackups int) OptFunc {
	return func(c *conf) {
		c.maxBackups = maxBackups
	}
}

// WithMaxAge 指定最大日志保留时间
func WithMaxAge(maxAge int) OptFunc {
	return func(c *conf) {
		c.maxAge = maxAge
	}
}

// WithCompress 指定是否压缩
func WithCompress(compress bool) OptFunc {
	return func(c *conf) {
		c.compress = compress
	}
}

// WithConsole 是否打印到控制台
func WithConsole(enable bool) OptFunc {
	return func(c *conf) {
		c.enableConsole = enable
	}
}
