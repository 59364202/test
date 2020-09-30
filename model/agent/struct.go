package agent

import ()

// KeyAccess is a data from api.agent
type KeyAccess struct {
	ID              int64         `json:"id"`               // example:`1` หมายเลข agent
	UserAccount     string        `json:"account"`          // example:`dataimport-bma` ชื่อผู้ใช้ agent
	AgentType       *SelectOption `json:"agent_type"`       // example:`dataimport,www` ประเภทของ agent
	PermissionRealm *SelectOption `json:"permission_realm"` // example:`haii-thaiwater30` realm ของ agent
	CallbackURL     string        `json:"callback_url"`     // example:`http://web.thaiwater.net/thaiwater30/apicb` callback url ของ web server
	RequestOrigin   string        `json:"request_origin"`   // example:`http://web.thaiwater.net` request origin ของ web server
	KeyAccess       string        `json:"key_access"`       // example:`117ad549211be34bbe5f2b10d58c16b` key สำหรับยืนยันระหว่าง web server กับ api server
	FullName        string        `json:"full_name"`        // example:`BMA` Full user description
}

// SelectOption is a data for select option
type SelectOption struct {
	Text  string `json:"text"`  // example:`ข้อความ` ข้อความอธิบายข้อมูล
	Value int64  `json:"value"` // example:`1` รหัสข้อมูล
}

// NewAgent is a data from new agent convert
type NewAgent struct {
	Agency    string `json:"agency"`     // example:`bma` ชื่อหน่วยงาน
	AgentName string `json:"agent_name"` // example:`dataimport_bma` username ของ agent
	AgentKey  string `json:"agent_key"`  // example:`117ad549211be34bbe5f2b10d58c16b` key สำหรับยืนยันระหว่าง web server กับ api server
}
