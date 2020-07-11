package main

import (
	"ddns/dnspod"
	"time"
)

func main() {
	for {
		dnspod.Run()
		time.Sleep(1 * time.Minute)
	}
}
