package zabbix

type Creds struct {
	User string
	Password string
}

type Zabbix struct {
	apiUrl string
	Auth string
}

func NewZabbix(url string, creds Creds) (Zabbix,error) {
	//Zabbix type consctructor with authentication
	auth, err := login(url, creds)
	return Zabbix{url,auth}, err
}

func login(url string, creds Creds) (string,error) {
	request := `{
			"jsonrpc":"2.0",
			"method":"user.login",
			"id":1,
			"auth":null,
			"params": {
				"user": "` + creds.User +`",
        		"password": "`+creds.Password+`"
			}
		}`
	data, err := zApiRequest(url,request)
	return data["result"].(string),err
}

func (z *Zabbix) ZRequest(body string) (map[string]interface{}, error) {
	//Wrapper to call request as a class method
	return zApiRequest(z.apiUrl , body)
}
