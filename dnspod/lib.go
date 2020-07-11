package dnspod

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

const (
	dnpodKey    = "DNSPOD_KEY"
	dnpodDomain = "DNSPOD_DOMAIN"
)

const (
	userDetail = "https://dnsapi.cn/User.Detail"
	recordList = "https://dnsapi.cn/Record.List"
	recordDDNS = "https://dnsapi.cn/Record.Ddns"
	getPublic  = "http://icanhazip.com"
)

const (
	one     = "1"
	enabled = "enabled"
)

func unmarshal(body io.Reader, v interface{}) error {
	j, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(j, &v)
}

func getPublicIP() (net.IP, error) {
	resp, err := http.Get(getPublic)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	ipStr := strings.TrimSpace(string(body))
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, errors.New("parse ip failed")
	}
	return ip, resp.Body.Close()
}

func request(uri, body string) (*http.Response, error) {
	req, err := http.NewRequest("POST", uri, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header["Accept"] = []string{"application/json"}
	header["Content-Type"] = []string{"application/x-www-form-urlencoded", "charset=UTF-8"}
	header["User-Agent"] = []string{"autoDDNS/0.1 (2509222540@qq.com)"}
	req.Header = header
	client := &http.Client{}
	return client.Do(req)
}