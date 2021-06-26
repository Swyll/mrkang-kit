package zaplogger

import (
	"os"

	"github.com/Swyll/mrkang-kit/log"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConf struct {
	FilePath      string `ini:"filepath" comment:"日志路径"`
	MaxSize       int    `ini:"maxsize" comment:"每个日志文件保存的最大尺寸单位为M"`
	MaxBackup     int    `ini:"maxbackup" comment:"日志文件最多保存的个数"`
	MaxAge        int    `ini:"maxage" comment:"日志文件最多保存的天数"`
	Compress      bool   `ini:"compress" comment:"是否压缩"`
	EnableConsole bool   `ini:"enableconsole" comment:"是否打印到控制台"`
}

func NewDefauLogger() log.Logger {
	return NewLog(&LoggerConf{
		FilePath:      "/var/log/swy.log",
		MaxSize:       20,
		MaxBackup:     5,
		MaxAge:        100,
		Compress:      false,
		EnableConsole: true,
	})
}

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func NewLog(conf *LoggerConf) log.Logger {
	//now := time.Now()
	/**
	 *lumberjack.Logger实现了io.Writer接口,而zap和zapcore中的很多结构体都实现了io.Writer,所以
	 *实际上zap中的io调用的是lumberjack.Logger实现的方法
	 */
	hook := &lumberjack.Logger{
		//Filename:   fmt.Sprintf("log/%04d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), //filePath
		Filename:   conf.FilePath,
		MaxSize:    conf.MaxSize, // megabytes
		MaxBackups: conf.MaxBackup,
		MaxAge:     conf.MaxAge,   //days
		Compress:   conf.Compress, // disabled by default
	}
	defer hook.Close()
	/*zap 的 Config 非常的繁琐也非常强大，可以控制打印 log 的所有细节，因此对于我们开发者是友好的，有利于二次封装。
	  但是对于初学者则是噩梦。因此 zap 提供了一整套的易用配置，大部分的姿势都可以通过一句代码生成需要的配置。
	*/
	enConfig := zap.NewProductionEncoderConfig() //生成配置

	// 时间格式
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	level := zap.InfoLevel
	w := zapcore.AddSync(hook)
	allCore := make([]zapcore.Core, 0, 2)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig), //编码器配置
		w,                                   //打印到控制台和文件
		level,                               //日志等级
	)
	allCore = append(allCore, fileCore)

	if conf.EnableConsole {
		//打印到控制台
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(enConfig),
			os.Stdout,
			level,
		)
		allCore = append(allCore, consoleCore)
	}

	core := zapcore.NewTee(allCore...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	//_log := log.New(hook, "", log.LstdFlags)
	//logger.Sugar().Info()
	return logger.Sugar()
	//_log.Println("Start...")
}
