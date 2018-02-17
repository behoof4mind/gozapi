package user

import (
	"gozapi/zabbix"
)

func (u *User) Login(z *zabbix.Zabbix, creds LoginParams) (map[string]interface{},error) {
	b, err := z.MakeBody("user.login", creds)
	resp, err := z.ZRequest(b)
	if resp["result"] != nil {
		if !creds.UserData {
			z.Auth = resp["result"].(string)
		}
	}
	return resp , err
}