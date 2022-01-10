package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func (z *Zabbix) GetRequest(respBody *bytes.Buffer) ([]byte, error) {

	resp, err := http.Post(fmt.Sprintf("http://%s/zabbix/api_jsonrpc.php", z.Conf.Host), "application/json-rpc", respBody)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func (z *Zabbix) WorkJson(newStruct interface{}) ([]byte, error) {
	postBody, err := json.Marshal(newStruct)
	if err != nil {
		return []byte{}, err
	}
	responseBody := bytes.NewBuffer(postBody)
	data, err := z.GetRequest(responseBody)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (z *Zabbix) BuildJson(params interface{}, method string, id int) ([]byte, error) {
	newStruct := ResponseWithoutAuth{
		DefaultJsonFields: DefaultJsonFields{
			Jsonrpc: "2.0",
			Method:  method,
			Id:      id,
		},
		Params: params,
	}
	data, err := z.WorkJson(newStruct)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (z *Zabbix) BuildJsonWithAuth(params interface{}, method string, id int) ([]byte, error) {
	newStruct := ResponseWithAuth{
		DefaultJsonFields: DefaultJsonFields{
			Jsonrpc: "2.0",
			Method:  method,
			Id:      id,
		},
		Params: params,
		Auth:   z.Token,
	}
	data, err := z.WorkJson(newStruct)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func StrToTime(timeStr string) (string, error) {
	i, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return "Error", err
	}
	tm := time.Unix(i, 0).String()
	return tm, nil
}
