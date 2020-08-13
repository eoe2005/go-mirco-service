package log

import (
	"fmt"
	"time"
)

// Debug 打印日志
func Debug(msg string, args ...interface{}) {
	star := fmt.Sprintf(msg, args...)
	fmt.Println(fmt.Sprintf("[%v] DEBUG %s", getTime(), star))

}

// Error 打印日志
func Error(msg string, args ...interface{}) {
	star := fmt.Sprintf(msg, args...)
	fmt.Println(fmt.Sprintf("[%v] Error %s", getTime(), star))

}

func getTime() string {
	return fmt.Sprintf("%v", time.Now())
}
