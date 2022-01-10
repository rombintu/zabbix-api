package zabbixapi

import (
	"strconv"
	"time"
)

type ZabbixConf struct {
	Host string `toml:"Host"`
	User string `toml:"User"`
	Pass string `toml:"Password"`
}

type Zabbix struct {
	Conf  ZabbixConf
	Token string
}

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

func StrToTime(timeStr string) (string, error) {
	i, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return "Error", err
	}
	tm := time.Unix(i, 0).String()
	return tm, nil
}
