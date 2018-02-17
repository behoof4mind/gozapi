package user

import (
	"gozapi/zabbix"
)

func (u *User) Login(z *zabbix.Zabbix, creds zabbix.Creds) (map[string]interface{},error) {
	b, err := z.MakeBody("user.login", map[string]interface{}{"user":creds.User, "password":creds.Password})
	au, err := z.ZRequest(b)
	if au["result"] != nil {
		z.Auth = au["result"].(string)
	}
	return au , err
}