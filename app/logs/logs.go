package logs

import (
	"os"
	// "path"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	LogFilePath = "./logs/"
	logFileName = "latest.log"
	Logger      = logrus.New() // 初始化日志对象
	LogEntry    *logrus.Entry
)

func init() {
	if _, err := os.Stat(LogFilePath); err != nil {
		os.MkdirAll(LogFilePath, os.ModePerm)
	}

	// 日志文件
	// fileName := path.Join(LogFilePath, logFileName)
	// 写入文件
	// src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	// if err != nil {
	// 	panic(err)
	// }
	//设置输出
	// logger.Out = src
	// 实例化
	logger := logrus.New()
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		LogFilePath+"err-%Y%m%d.log",

		// 生成软链，指向最新日志文件
		// rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		panic(err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	Logger = logger
}
