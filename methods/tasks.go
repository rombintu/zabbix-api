package methods

import "encoding/json"

type Task struct {
	Taskid string `json:"taskid"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Clock  string `json:"clock"`
	Ttl    string `json:"ttl"`
}

func (z *Zabbix) GetTaskById(taskids string) (Task, error) {
	params := struct {
		Output  string `json:"output"`
		Taskids string `json:"taskids"`
	}{
		Output:  "extend",
		Taskids: taskids,
	}

	data, err := z.BuildJsonWithAuth(params, "task.get", 1)
	if err != nil {
		return Task{}, err
	}
	var result Task
	if err := json.Unmarshal(data, &result); err != nil {
		return Task{}, err
	}
	return result, nil
}
