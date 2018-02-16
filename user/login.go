package user

import (
	"gozapi/zabbix"
)

type User struct {

}

func (u *User) Login(z *zabbix.Zabbix, creds zabbix.Creds) (string,error) {
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
	au, err := zabbix.ZApiRequest(&z.ApiUrl,request)
	z.Auth = au["result"].(string)
	z.AuthId = "1"
	return z.Auth , err
}