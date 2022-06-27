package app

import (
	"fmt"
	"log"
	"strings"

	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type AppClient struct {
	Ch     chan bool
	Engine *gin.Engine
}

var App *AppClient
var (
	//日志地址
	logFilePath = "./"
	//日志文件名
	logFileName = "system.log"
)

func init() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	App = new(AppClient)
	App.Ch = make(chan bool)
	App.Engine = gin.Default()
	// App.Engine.Use(logerMiddleware())
}

func (this *AppClient) Static(path, name string) *AppClient {
	if path != "" {
		if _, err := os.Stat(path); err != nil {
			os.MkdirAll(path, os.ModePerm)
		}
		if name == "" {
			name = "/" + strings.Trim(path, "./")
		}
		this.Engine.Static(name, path)
	}
	return App
}

func (this *AppClient) Template(str string) *AppClient {
	strtmp := strings.Trim(str, "*")
	if _, err := os.Stat(strtmp); err != nil {
		log.Println(err)
		return this
	}
	this.Engine.LoadHTMLGlob(str)
	return this
}

func (this *AppClient) Run(port int) {
	this.WebRouter()
	portstr := fmt.Sprintf(":%d", port)
	fmt.Println("地址: http://localhost" + portstr)
	this.Engine.Run(portstr)
}

func logerMiddleware() gin.HandlerFunc {
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	logger := logrus.New()
	//设置日志级别
	logger.SetLevel(logrus.InfoLevel)
	//设置输出
	logger.Out = src
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

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

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
		}).Info()

	}
}
