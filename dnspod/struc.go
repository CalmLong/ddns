package dnspod

type Records struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	Line          string      `json:"line"`
	LineId        string      `json:"line_id"`
	Type          string      `json:"type"`
	TTL           string      `json:"ttl"`
	Value         string      `json:"value"`
	Weight        interface{} `json:"weight"`
	MX            string      `json:"mx"`
	Enabled       string      `json:"enabled"`
	Status        string      `json:"status"`
	MonitorStatus string      `json:"monitor_status"`
	Remark        string      `json:"remark"`
	UpdatedOn     string      `json:"updated_on"`
	UseAqb        string      `json:"use_aqb"`
}

type User struct {
	RealName          string `json:"real_name"`
	UserType          string `json:"user_type"`
	Telephone         string `json:"telephone"`
	Im                string `json:"im"`
	Nick              string `json:"nick"`
	Id                string `json:"id"`
	Email             string `json:"email"`
	QQ                string `json:"qq"`
	Status            string `json:"status"`
	EmailVerified     string `json:"email_verified"`
	TelephoneVerified string `json:"telephone_verified"`
	WeixinBinded      string `json:"weixin_binded"`
	AgentPending      bool   `json:"agent_pending"`
	Balance           int64  `json:"balance"`
	Smsbalance        int64  `json:"smsbalance"`
	UserGrade         string `json:"user_grade"`
	IsDtokenOn        bool   `json:"is_dtoken_on"`
}

type Info struct {
	User        User   `json:"user"`
	UserId      string `json:"user_id"`
	Uin         string `json:"uin"`
	Avatar      string `json:"avatar"`
	AvatarId    string `json:"avatar_id"`
	SubDomains  string `json:"sub_domains"`
	RecordTotal string `json:"record_total"`
	RecordsNum  string `json:"records_num"`
}

type Domain struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Punycode  string   `json:"punycode"`
	Grade     string   `json:"grade"`
	Owner     string   `json:"owner"`
	ExtStatus string   `json:"ext_status"`
	TTL       int64    `json:"ttl"`
	DnspodNS  []string `json:"dnspod_ns"`
}

type Status struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type Result struct {
	Status  Status    `json:"status"`
	Domain  Domain    `json:"domain"`
	Info    Info      `json:"info"`
	Records []Records `json:"records"`
}
