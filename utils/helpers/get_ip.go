package helpers

import (
	"net"
	"net/http"
	"strings"
)

func GetUserIP(req *http.Request) string {
	ip := req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = req.Header.Get("X-Real-Ip")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(req.RemoteAddr)
	}
	return strings.TrimSpace(ip)
}
