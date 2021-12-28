# Zabbix-api

### Install
```sh
go get -u "github.com/rombintu/zabbix-api"
```

### Example
```go
package main

import (
	"fmt"
	"log"

	zabbixapi "github.com/rombintu/zabbix-api"
)

func main() {
	z := zabbixapi.NewZabbix(
		"192.168.213.127",
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
```