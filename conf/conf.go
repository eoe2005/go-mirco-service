package conf

import "sync"

//注册中心服务器配置
type hostConfData struct {
	lock    sync.Mutex
	hostMap map[string]map[string]int32
}
type confData struct {
	lock sync.Mutex
	conf map[string][2]string
}

var hostCache hostConfData
var confCache confData

func init() {
	hostCache = hostConfData{hostMap: make(map[string]map[string]int32, 10)}
	confCache = confData{conf: map[string][2]string{}}
}

// 增加服务器
func (d *hostConfData) AddHost(app, host string) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	v, ok := d.hostMap[app]
	if !ok {
		v = make(map[string]int32, 10)
		d.hostMap[app] = v
	}
	v[host] = 0
	return false
}

//删除服务器
func (d *hostConfData) DelHost(app, host string) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	v, ok := d.hostMap[app]
	if ok {
		delete(v, host)
	}
	return true
}

//获取服务器
func (d *hostConfData) getHost(app) string {

	if v, ok := d.hostMap[app],ok {
		for h, v := range v {
			return h
		}
	}
	return ""
}
