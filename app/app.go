package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"ucenter/app/logs"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
)

type AppClient struct {
	Ch     chan bool
	Engine *gin.Engine
}

var App *AppClient

func init() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	App = new(AppClient)
	App.Ch = make(chan bool)

	// 记录日志到文件
	// f, _ := os.Create("log.log")
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	path := logs.LogFilePath + "gin"
	writer, _ := rotatelogs.New(
		path+"-%Y%m%d.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//这里设置1分钟产生一个日志文件
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	gin.DefaultWriter = io.MultiWriter(writer)

	// // 记录错误日志到文件，同时输出到控制台
	// fErr, _ := os.Create("err.log")
	// gin.DefaultErrorWriter = io.MultiWriter(fErr, os.Stdout)

	App.Engine = gin.Default()
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

	srv := &http.Server{
		Addr:    portstr,
		Handler: this.Engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Logger.Error("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logs.Logger.Error("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
	// this.Engine.Run(portstr)
}
