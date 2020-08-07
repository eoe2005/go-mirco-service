package server

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

// udp 提供服务注册发现功能
func udpRun() {

}

// Http 接口
func httpRun() {

}

// 配置推送功能
func tcpRun() {

}
