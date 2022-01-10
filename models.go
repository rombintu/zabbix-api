package zabbixapi

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type DefaultJsonFields struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Id      int    `json:"id"`
}

type ResponseWithAuth struct {
	DefaultJsonFields
	Params interface{} `json:"params"`
	Auth   string      `json:"auth"`
}

type ResponseWithoutAuth struct {
	DefaultJsonFields
	Params interface{} `json:"params"`
}
