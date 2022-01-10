package methods

import (
	"encoding/json"
	"errors"
)

type UserLoginParams struct {
	Username string `json:"user"`
	Password string `json:"password"`
	UserData bool   `json:"userData"`
}

type UserLoginResp struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
}

func (z *Zabbix) GetToken() (UserLoginResp, error) {

	params := UserLoginParams{
		Username: z.Conf.User,
		Password: z.Conf.Pass,
		UserData: false,
	}

	data, err := z.BuildJson(params, "user.login", 1)
	if err != nil {
		return UserLoginResp{}, err
	}
	var result UserLoginResp
	if err := json.Unmarshal(data, &result); err != nil {
		return UserLoginResp{}, err
	}
	if result.Result != "" {
		z.Token = result.Result
	} else {
		return result, errors.New(result.Error.Message)
	}
	return result, nil
}
