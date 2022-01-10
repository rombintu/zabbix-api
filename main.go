package zabbixapi

func NewZabbix(host, user, pass string) Zabbix {
	c := ZabbixConf{
		Host: host,
		User: user,
		Pass: pass,
	}
	return Zabbix{
		Conf: c,
	}
}
