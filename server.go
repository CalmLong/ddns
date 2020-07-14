package main

import (
	"net"
	"net/http"
	"strings"
)

var handle = http.NewServeMux()

func init() {
	handle.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		xForwardedFor := r.Header.Get("X-Forwarded-For")
		xs := strings.Split(xForwardedFor, ",")
		if xs == nil {
			_, _ = w.Write([]byte("Get X-Forwarded-For params failed"))
			return
		}
		ip := strings.TrimSpace(xs[0])
		if ip != "" {
			_, _ = w.Write([]byte(ip))
			return
		}
		ip = strings.TrimSpace(r.Header.Get("X-Real-IP"))
		if ip != "" {
			_, _ = w.Write([]byte(ip))
			return
		}
		if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
			_, _ = w.Write([]byte(ip))
			return
		}
		_, _ = w.Write([]byte("Get public ip failed"))
	})
}

func main() {
	if err := http.ListenAndServe(":6666", handle); err != nil {
		panic(err)
	}
}
