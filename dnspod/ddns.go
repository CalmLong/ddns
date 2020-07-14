package dnspod

import (
	"fmt"
	"os"
	"strings"
)

var pubParams string

var domain []string

func GetUserInfo() (info Result) {
	resp, err := request(userDetail, pubParams)
	if err != nil {
		fmt.Println("[E]", err.Error())
		return info
	}
	if err := unmarshal(resp.Body, &info); err != nil {
		fmt.Println("[E]", err.Error())
		return info
	}
	return info
}

func GetSubRecord() (result Result) {
	params := fmt.Sprintf("%s&domain=%s&sub_domain=%s", pubParams, domain[0], domain[1])
	resp, err := request(recordList, params)
	if err != nil {
		fmt.Println("[E]", err.Error())
		return result
	}
	if err := unmarshal(resp.Body, &result); err != nil {
		fmt.Println("[E]", err.Error())
		return result
	}
	return result
}

// 必要的环境变量
// DNSPOD_DOMAIN=ilouis.cn,www
// DNSPOD_KEY=YOUR_ID,YOUR_KEY
// 注意事项：调用 run() 之前该子域名必须存在
func Run(extServer ...string) {
	i, err := getPublicIP(extServer...)
	if err != nil {
		fmt.Println("[E]", err.Error())
		return
	}
	pubIp := i.String()
	if strings.TrimSpace(pubIp) == "" {
		fmt.Println("[E] Public ip is none")
		return
	}
	u := GetUserInfo()
	// 检测账户是否被禁用
	if !strings.EqualFold(u.Info.User.Status, enabled) {
		fmt.Println("[E] Your account has been disabled: ", u.Info.User.Status)
		return
	}
	// 状态码不等于 1 时返回错误
	if !strings.EqualFold(u.Status.Code, one) {
		fmt.Println("[E] msg: ", u.Status.Message)
		return
	}
	ret := GetSubRecord()
	// 查看是否存在 A 记录，不存在时报错
	if ret.Records == nil || ret.Status.Code != "1" {
		fmt.Println("[E] Get sub record failed: ", ret.Status)
		return
	}
	value := ret.Records[0].Value
	// 比较本机 IP 是否和云端记录一致，一致时则不更新记录
	// 注意事项：如果1小时之内，提交了超过5次没有任何变动的记录修改请求，该记录会被系统锁定1小时，不允许再次修改，所以在开发和测试的过程中，请自行处理IP变动，仅在本地IP发生变动的情况下才调用本接口。
	// 如何理解没有任何变动的记录修改请求？比如原记录值已经是 1.1.1.1，新的请求还要求修改为 1.1.1.1。
	// 信息来源：https://www.dnspod.cn/docs/records.html#dns
	if strings.EqualFold(pubIp, value) {
		fmt.Println("[I] Nothing to do. local:", pubIp, " record:", value)
		return
	}
	params := fmt.Sprintf("%s&domain=%s&record_id=%s&sub_domain=%s&record_line=默认&value=%s",
		pubParams, domain[0], ret.Records[0].Id, domain[1], pubIp)
	// 更新记录
	resp, err := request(recordDDNS, params)
	if err != nil {
		fmt.Println("[E]", err.Error())
		return
	}
	var res Result
	body := resp.Body
	if err := unmarshal(body, &res); err != nil {
		fmt.Println("[E]", err.Error())
		return
	}
	fmt.Println("[I]", res.Status.Message, ":", pubIp)
}

func init() {
	key := os.Getenv(dnpodKey)
	if key == "" {
		panic("DNSPOD_KEY not found")
	}
	domains := strings.Split(os.Getenv(dnpodDomain), ",")
	if len(domains) < 2 {
		panic("DNSPOD_DOMAIN params error")
	}
	domain = domains
	pubParams = fmt.Sprintf("login_token=%s&format=json&error_on_empty=no&lang=en", key)
}
