package main

import (
	"flag"
	"fmt"
)

var port = flag.Int("port", 18998, "开启的端口号")

func main() {
	flag.Parse()
	fmt.Println(*port)
}
