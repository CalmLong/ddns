package main

import (
	"ddns/dnspod"
	"flag"
	"time"
)

func main() {
	server := flag.String("s", "", "custom server")
	flag.Parse()
	for {
		dnspod.Run(*server)
		time.Sleep(1 * time.Minute)
	}
}
