package log

import "testing"

// Test debug
func TestDebug(t *testing.T) {
	Debug("你好 %s", "asdf")
}

func TestError(t *testing.T) {
	Error("你好 %s", "asdf")
}
