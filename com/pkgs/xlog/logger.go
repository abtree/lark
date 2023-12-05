package xlog

import (
	"fmt"
	"lark/com/files"
	"lark/pb"
	"os"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	FileDebug = "/debug.log"
	FileInfo  = "/info.log"
	FileWarn  = "/warn.log"
	FileError = "/error.log"
	FilePanic = "/panic.log"
)

const (
	// CallerTypeNormal   = 0 //自定义
	CallerTypeSeparate = 1 //相对路径
)

const (
	//记录8层堆载
	CallerDepth = 8
)

const (
	//日志编码级别
	// 小写编码器(默认)
	LowercaseLevelEncoder = "Lowercase"
	// 小写编码器带颜色
	LowercaseColorLevelEncoder = "LowercaseColor"
	// 大写编码器
	CapitalLevelEncoder = "Capital"
	// 大写编码器带颜色
	CapitalColorLevelEncoder = "CapitalColor"
)

var xLog *zap.SugaredLogger

//创建一个默认的日志记录
func Shared(cfg *pb.Jlogger, directory string) {
	xLog = newLogger(cfg)
}

//如果需要创建一个自己特有的日志记录
func NewLog(cfg *pb.Jlogger, directory string) *zap.SugaredLogger {
	cfg.CallerType = CallerTypeSeparate
	return newLogger(cfg)
}

func newLogger(cfg *pb.Jlogger) *zap.SugaredLogger {
	//根据不同的级别 输出到不同的日志文件
	//调试级别
	debugLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.DebugLevel
	})
	//日志级别
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.InfoLevel
	})
	//警告级别
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.WarnLevel
	})
	//错误级别
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.ErrorLevel
	})
	//灾难级别
	panicLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.PanicLevel
	})

	//创建日志存放路径
	path := cfg.Path + cfg.Directory
	if !files.IsDir(path) {
		err := files.MKDir(path)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}
	//创建日志core
	cores := [...]zapcore.Core{
		getEncoderCore(path+FileDebug, debugLevel, cfg),
		getEncoderCore(path+FileInfo, infoLevel, cfg),
		getEncoderCore(path+FileWarn, warnLevel, cfg),
		getEncoderCore(path+FileError, errorLevel, cfg),
		getEncoderCore(path+FilePanic, panicLevel, cfg),
	}

	//NewTee创建一个Core，将日志条目复制到两个或更多的底层Core中
	logger := zap.New(zapcore.NewTee(cores[:]...))
	//用文件名、行号和zap调用者的函数名注释每条消息
	if cfg.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	sl := logger.Sugar()
	sl.Sync()
	return sl
}

func getEncoderCore(filename string, level zapcore.LevelEnabler, cfg *pb.Jlogger) zapcore.Core {
	writer := getWriteSyncer(filename, cfg)
	return zapcore.NewCore(getEncoder(cfg), writer, level)
}

// 使用lumberjack进行日志分割
func getWriteSyncer(filename string, cfg *pb.Jlogger) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   filename, //日志文件的位置
		MaxSize:    int(cfg.Segment.MaxSize),
		MaxBackups: int(cfg.Segment.MaxBackups),
		MaxAge:     int(cfg.Segment.MaxAge),
		Compress:   cfg.Segment.Compress,
	}
	if cfg.LogStdout == true {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
	}
	return zapcore.AddSync(hook)
}

func getEncoder(cfg *pb.Jlogger) zapcore.Encoder {
	switch cfg.Encoder {
	case "json":
		return zapcore.NewJSONEncoder(*getEncoderConfig(cfg))
	default:
		return zapcore.NewConsoleEncoder(*getEncoderConfig(cfg))
	}
}

func getEncoderConfig(cfg *pb.Jlogger) *zapcore.EncoderConfig {
	config := &zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  cfg.StacktraceKey,              //链路追踪名
		LineEnding:     zapcore.DefaultLineEnding,      //结尾符 '\n'
		EncodeTime:     customTimeEncoder,              // 时间格式 zapcore.ISO8601TimeEncoder
		EncodeDuration: zapcore.SecondsDurationEncoder, //编码间隔
		EncodeCaller:   shortCallerEncoder,             //默认采用自定义路径
		EncodeName:     zapcore.FullNameEncoder,        //
	}
	//如果是设置的绝对路径
	if cfg.CallerType == CallerTypeSeparate {
		config.EncodeCaller = zapcore.ShortCallerEncoder //绝对路径:zapcore.FullCallerEncoder,相对路径:zapcore.ShortCallerEncoder
	}
	//设置日志级别
	switch cfg.EncodeLevel {
	case LowercaseLevelEncoder:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case LowercaseColorLevelEncoder:
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case CapitalLevelEncoder:
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case CapitalColorLevelEncoder:
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}
	return config
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000") + "]")
}

//自定义的路径结构
func shortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(getCaller(CallerDepth))
}

func getCaller(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return ""
	}
	return prettyCaller(file, line)
}

func prettyCaller(file string, line int) string {
	idx := strings.LastIndexByte(file, '/')
	if idx < 0 {
		return fmt.Sprintf("%s:%d", file, line)
	}

	idx = strings.LastIndexByte(file[:idx], '/')
	if idx < 0 {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return fmt.Sprintf("%s:%d", file[idx+1:], line)
}
