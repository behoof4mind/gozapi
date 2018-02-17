package zabbix


type Zabbix struct {
	ApiUrl string
	Auth interface{}
	AuthId int
}

type ApiBody struct {
	Jsonrpc string `json:"jsonrpc"`
	Method string	`json:"method"`
	Params interface{} `json:"params"`
	Auth interface{} `json:"auth"`
	Id int `json:"id"`
}