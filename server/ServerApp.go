package server

import (
	"net/http"
)

var port = 9090

type httpFanc func(string)

var (
	urlMethodMap = map[string]map[string]httpFanc{}
)

// SetPort 设置服务监听的端口
func SetPort(httpPort int) {
	port = httpPort
}

// addUrlMethodMap 添加路由信息
func addUrlMethodMap(method, url string, funcName httpFanc) {
	v, ok := urlMethodMap[url]
	if ok {
		v[method] = funcName
	} else {
		urlMethodMap[url] = map[string]httpFanc{method: funcName}
	}
}

// Get asd
func Get(url string, funcName httpFanc) {
	addUrlMethodMap("GET", url, funcName)
}

// Post Post数据
func Post(url string, funcName httpFanc) {
	addUrlMethodMap("POST", url, funcName)
}

// Delete 删除
func Delete(url string, funcName httpFanc) {
	addUrlMethodMap("DELETE", url, funcName)
}

// Put Put提交
func Put(url string, funcName httpFanc) {
	addUrlMethodMap("PUT", url, funcName)
}

// Option 提交
func Option(url string, funcName httpFanc) {
	addUrlMethodMap("OPTION", url, funcName)
}

// Any 提交
func Any(url string, funcName httpFanc) {
	addUrlMethodMap("ANY", url, funcName)
}

// Run aasdf
func Run() {
	for k, v := range urlMethodMap {
		var fc httpFanc = nil
		http.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			f, o := v[r.Method]
			if o {
				fc = f
			} else {
				f, o := v["ANY"]
				if o {
					fc = f
				}
			}
			if f != nil {
				f(k)
			}

		})
	}

	http.ListenAndServe("127.0.0.1:8080", nil)
}
