package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	port = 9090
	// AppName 微服务的名字
	AppName = "app"
)

// SetAppName 设置应用的名字
func SetAppName(name string) {
	AppName = name
}

// GData http接口封装的方法
type GData struct {
	w http.ResponseWriter
	r *http.Request
}

// Success 接口成功之后的调用
func (g GData) Success(a interface{}) {
	data := map[string]interface{}{"code": 0, "msg": "", "data": a}
	r, err := json.Marshal(data)
	if err == nil {
		g.w.Write(r)
	}

}

// Fail 失败返回的数据
func (g GData) Fail(code int, msg string) {
	data := map[string]interface{}{"code": code, "msg": msg, "data": ""}
	r, err := json.Marshal(data)
	if err == nil {
		g.w.Write(r)
	}
}

// HttpFanc 注册http请求的方法
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

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}
