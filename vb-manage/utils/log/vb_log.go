package VBLog

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"vb-server/utils/config"
)

var Print *zap.Logger


func logCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	//日志文件路径配置2

	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)
}

func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	core := logCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
}

func InitLog(){

	//config
	logInfo:=config.VBGetConfigLog()

	//print
	fmt.Printf("日志参数:\n")
	fmt.Printf("服务名称 %v\n",logInfo.Name)
	fmt.Printf("日志路径 %v\n",logInfo.Path)
	fmt.Printf("日志文件最大MB %v\n",logInfo.MaxSize)
	fmt.Printf("日志文件最多备份 %v\n",logInfo.MaxBackups)
	fmt.Printf("日志文件最长保存时间 %v\n",logInfo.MaxAge)
	fmt.Printf("是否压缩 %v\n",logInfo.Compress)

	Print=NewLogger(logInfo.Path, zapcore.InfoLevel, logInfo.MaxSize, logInfo.MaxBackups, logInfo.MaxAge, logInfo.Compress,logInfo.Name)
}
