package zabbix


type Creds struct {
	User string
	Password string
}

type Zabbix struct {
	ApiUrl string
	Auth interface{}
	AuthId int
}

type ApiBody struct {
	Jsonrpc string `json:"jsonrpc"`
	Method string	`json:"method"`
	Params map[string]interface{} `json:"params"`
	Auth interface{} `json:"auth"`
	Id int `json:"id"`
}

type Auth interface {

}