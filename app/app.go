package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type AppClient struct {
	Ch     chan bool
	Engine *gin.Engine
}

var App *AppClient

func init() {
	gin.SetMode(gin.ReleaseMode)
	App = new(AppClient)
	App.Ch = make(chan bool)
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
		fmt.Println(name, path)
		this.Engine.Static(name, path)
	}
	return App
}

func (this *AppClient) Run(port int) {
	this.WebRouter()
	portstr := fmt.Sprintf(":%d", port)
	fmt.Println("地址: http://localhost" + portstr)
	this.Engine.Run(portstr)
}
