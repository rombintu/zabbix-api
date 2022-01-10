package zabbixapi

import "encoding/json"

type ProblemGetParams struct {
	Output     string   `json:"output"`
	Resent     string   `json:"resent"`
	SorteField []string `json:"sortfield"`
	Sortorder  string   `json:"sortorder"`
}

type ProblemGetResp struct {
	Result []Problems `json:"result"`
	Error  Error      `json:"error"`
}

type Problems struct {
	Clock        string `json:"clock"`
	Text         string `json:"name"`
	Acknowledged string `json:"acknowledged"`
}

func (z *Zabbix) GetProblems() (ProblemGetResp, error) {
	params := ProblemGetParams{
		Output:     "extend",
		Resent:     "true",
		SorteField: []string{"eventid"},
		Sortorder:  "DESC",
	}

	data, err := z.BuildJsonWithAuth(params, "problem.get", 1)
	if err != nil {
		return ProblemGetResp{}, err
	}
	var result ProblemGetResp
	if err := json.Unmarshal(data, &result); err != nil {
		return ProblemGetResp{}, err
	}
	return result, nil
}
