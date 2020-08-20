package conf

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/eoe2005/go-mirco-service/log"
)

var (
	//服务器保存的服务注册内容
	serverService map[string]map[string]int64
	// 服务配置保持的额锁
	serverServiceLock sync.Mutex
	// 服务器保存的配置数据
	serverConf map[string]map[string]interface{}
	// 服务器配置保持的锁
	serverConfLock sync.Mutex

	// 客户端保存的可用服务信息
	clientService map[string]map[string]int

	//客户端服务更新锁
	clientServiceLock sync.Mutex
	// 客户端保持的配置数据
	clientConf map[string]map[string]interface{}
	// 客户端配置更新锁
	clientConfLock sync.Mutex

	// 程序运行等待退出
	mainRun sync.WaitGroup
	// 程序是否运行
	isRun = true

	// 客户端的链接
	clientConMap map[net.TCPConn]int = map[net.TCPConn]int{}
)

// RunServer 服务器运行方法
func RunServer(port int) {
	mainRun.Add(1)
	go udpRun(port)
	mainRun.Add(1)
	go tcpRun(port)
	mainRun.Wait()
}

// tcp 运行，服务发现，推送服务功能
func tcpRun(port int) {
	defer mainRun.Done()
	addrString := fmt.Sprintf("0.0.0.0:%d", port)
	addr, err := net.ResolveTCPAddr("tcp", addrString)
	if err != nil {
		log.Debug("注册中心配置启动失败 -> 绑定udp端口 %s", addrString)

	}
	server, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Debug("注册中心配置启动失败 -> 绑定端口 %s", addrString)
	}
	defer server.Close()
	for isRun {
		con, err := server.AcceptTCP()
		if err != nil {
			log.Debug(" %v", err)
		} else {
			clientConMap[*con] = 1
			go serverProcessClient(con)
		}

	}
}

// 处理客户端链接
func serverProcessClient(con *net.TCPConn) {

}

// 发送数据给客户端
func sendClient(cmd, data string) {
	line := fmt.Sprintf("%s|||%s\n", cmd, data)
	for con, _ := range clientConMap {
		con.Write([]byte(line))
	}
}

// udp运行，为了服务注册
func udpRun(port int) {
	defer mainRun.Done()
	addrString := fmt.Sprintf("0.0.0.0:%d", port)
	addr, err := net.ResolveUDPAddr("udp", addrString)
	if err != nil {
		log.Debug("注册中心配置启动失败 -> 绑定udp端口 %s", addrString)

	}
	updcon, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Debug("注册中心配置启动失败 -> 绑定端口 %s", addrString)
	}
	defer updcon.Close()
	go releaseServerHost()
	for isRun {
		buf := make([]byte, 1024)
		len, addr, _ := updcon.ReadFromUDP(buf)
		log.Debug("收到服务的注册 -> %v %s", addr, string(buf[0:len]))
		go registerServerHost(addr, buf[0:len])
		//fmt.Sprintf("%s %d -> %v\n", addr, len, addr)
	}
}

// 处理注册来的服务地址
func registerServerHost(addr *net.UDPAddr, data []byte) {
	serverServiceLock.Lock()
	defer serverServiceLock.Unlock()
	appNamePort := strings.Split(string(data), ":")
	appName := appNamePort[0]
	port := appNamePort[1]

	host := fmt.Sprintf("%s:%s", addr.IP, port)

	v, ok := serverService[appName]
	if !ok {
		v = map[string]int64{}
		serverService[appName] = v
	}
	v[host] = time.Now().Unix()
	sendClient("as", fmt.Sprintf("%s:%s", appName, host))
}

// 定期删除掉掉线的服务
func releaseServerHost() {
	timeTicker := time.NewTicker(time.Second * 10)
	for {
		timeN := time.Now().Unix()
		serverServiceLock.Lock()
		for app, v := range serverService {
			for host, t := range v {
				if timeN-t > 10 {
					sendClient("ds", fmt.Sprintf("%s:%s", app, host))
					delete(v, host)
				}
			}
		}
		serverServiceLock.Unlock()
		<-timeTicker.C
	}
}
