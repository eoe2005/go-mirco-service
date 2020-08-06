package server

import "net"

//监听的端口
const POST = 3333

//注册过来的服务器
var hostConf map[string][]string

func MainConf() {
	println("123")
}

func init() {
	hostConf = make(map[string][]string, 10)
}

func addHost(host net.Addr) {

}
func udpRun() {

}
