package accesslog

import ()

// ResultAccessLog is a data from api.accesslog
type ResultAccessLog struct {
	ID              int64  `json:"id"`                // example:`40` ลำดับการเข้าถึง service
	AccessTime      string `json:"access_time"`       // example:`2017-06-07 17:51:00.067474+07` เวลาในการเรียก service
	AgentUser       string `json:"agent_user"`        // examplae:`daimport-bma` ชื่อ username ของ agent
	User            string `json:"user"`              // example:`admin` ชื่อ User
	ServerAgentUser string `json:"server_agent_user"` // example:`api` ชื่อ server agent
	Service         string `json:"service"`           // example:`login`
	ServiceMethod   string `json:"service_method"`    // example:`get`
	Host            string `json:"host"`              // example:`127.0.0.1:9200` ip address server
	RequestURL      string `json:"request_url"`       // example:`/api/v1/thaiwater30/public/thailand` url ที่ส่งคำขอเข้ามา
	AccessDuration  int64  `json:"access_duration"`   // example:`1535023` เวลาในการเข้าถึง service
	ReplyCode       int64  `json:"reply_code"`        // example:`200` reply code
	ReplyReason     string `json:"reply_reason"`      // example:`OK` ข้อความตอบกลับ
	ClientIP        string `json:"client_ip"`         // example:`192.168.1.56` ip address client
}

type AgentName struct {
	ID        interface{} `json:"id"`         
	AgentName interface{} `json:"agent_name"` 
}
