package zabbixapi_test

import (
	"fmt"
	"log"
	"testing"

	zabbixapi "github.com/rombintu/zabbix-api"
)

func TestGetProblems(t *testing.T) {
	z := zabbixapi.NewZabbix(
		"zabbixtest",
		"Admin",
		"zabbix",
	)
	if _, err := z.GetToken(); err != nil {
		log.Fatal(err)
	}
	problems, err := z.GetProblems()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(problems)
}

func TestGetTaskById(t *testing.T) {
	z := zabbixapi.NewZabbix(
		"zabbixtest",
		"Admin",
		"zabbix",
	)
	if _, err := z.GetToken(); err != nil {
		log.Fatal(err)
	}
	task, err := z.GetTaskById("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(task)
}
