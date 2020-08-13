package server

import (
	"net/http"
)

var port = 9090

// GData saf
type GData struct {
	w http.ResponseWriter
	r *http.Request
}

type HttpFanc func(g GData)

var (
	urlMethodMap = map[string]map[string]HttpFanc{}
)

// SetPort 设置服务监听的端口
func SetPort(httpPort int) {
	port = httpPort
}

// addUrlMethodMap 添加路由信息
func addUrlMethodMap(method, url string, funcName HttpFanc) {
	v, ok := urlMethodMap[url]
	if ok {
		v[method] = funcName
	} else {
		urlMethodMap[url] = map[string]HttpFanc{method: funcName}
	}
}

// Get asd
func Get(url string, funcName HttpFanc) {
	addUrlMethodMap("GET", url, funcName)
}

// Post Post数据
func Post(url string, funcName HttpFanc) {
	addUrlMethodMap("POST", url, funcName)
}

// Delete 删除
func Delete(url string, funcName HttpFanc) {
	addUrlMethodMap("DELETE", url, funcName)
}

// Put Put提交
func Put(url string, funcName HttpFanc) {
	addUrlMethodMap("PUT", url, funcName)
}

// Option 提交
func Option(url string, funcName HttpFanc) {
	addUrlMethodMap("OPTION", url, funcName)
}

// Any 提交
func Any(url string, funcName HttpFanc) {
	addUrlMethodMap("ANY", url, funcName)
}

// Run aasdf
func Run() {
	for k, v := range urlMethodMap {
		http.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			f, o := v[r.Method]
			if o {
				f(GData{w: w, r: r})
			} else {
				f, o := v["ANY"]
				if o {
					f(GData{w: w, r: r})
				}
			}
		})
	}

	http.ListenAndServe("127.0.0.1:8080", nil)
}
