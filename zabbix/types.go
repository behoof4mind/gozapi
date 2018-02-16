package zabbix


type Creds struct {
	User string
	Password string
}

type Zabbix struct {
	ApiUrl string
	Auth string
	AuthId string
}

type ApiBody struct {
	JsonRpc string
	Method ApiMethod
	Params map[string]interface{}
}

type ApiMethod struct {
	Key string
	Action string
}