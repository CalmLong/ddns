package dnspod_test

import (
	"ddns/dnspod"
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	fmt.Println(dnspod.GetUserInfo())
}

func TestGetRecordAll(t *testing.T) {
	rec := dnspod.GetSubRecord()
	fmt.Println("domain: ", rec.Domain)
	fmt.Println("records: ", rec.Records)
	fmt.Println("status: ", rec.Status)
}

func TestDDNS(t *testing.T) {
	dnspod.Run()
}