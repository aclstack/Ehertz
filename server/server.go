package server

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"net"
)

func NewServer(ip string, port int) *server.Hertz {
	if net.ParseIP(ip) == nil {
		panic(any("ip地址异常"))
	}

	if port< 80 || port > 65535 {
		panic(any("port端口异常"))
	}

	return server.Default(
		server.WithHostPorts(
			fmt.Sprintf("%s:%d", ip,port),
			),
		)
}