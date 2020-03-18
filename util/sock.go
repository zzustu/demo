package util

import (
	"fmt"
	"net"
	"time"
)

// 判断到某个主机的某个端口是否连通(TCP)
func Reachable(host string, port int) bool {
	if conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", port)), 3*time.Second); err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

// 获取一个本机未被占用的端口号
func LocalFreePort(port int) int {
	if port < 1 || port > 65534 {
		port = 8080
	}

	for Reachable("", port) {
		if port > 65534 {
			break
		}
		port = port + 1
	}
	return port
}
