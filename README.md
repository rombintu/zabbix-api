# Zabbix-api

```sh
go get -u "github.com/rombintu/zabbix-api"
```

```go
package main

import (
	zabbixapi "github.com/rombintu/zabbix-api"
)

func main() {
	z := zabbixapi.NewZabbix("host", "user", "pass")
	z.GetToken()
}
```