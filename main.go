package zabbixapi

import (
	"github.com/rombintu/zabbix-api/methods"
)

func NewZabbix(host, user, pass string) methods.Zabbix {
	c := methods.ZabbixConf{
		Host: host,
		User: user,
		Pass: pass,
	}
	return methods.Zabbix{
		Conf: c,
	}
}
