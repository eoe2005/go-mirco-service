package main

import (
	"flag"

	g "./server"
)

var (
	app      string
	confhost string
)

func init() {
	flag.StringVar(&app, "app", "", "选择输入的应用启动类型：Conf,Gw")
	flag.StringVar(&confhost, "host", "", "配置中心的地址：127.0.0.1:8080")
	flag.Parse()
}
func main() {
	switch app {
	case "Conf":
		g.MainConf()
	case "Gw":
		g.MainGw()
	}
}
