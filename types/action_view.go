package types

import (
	"encoding/base64"
	"encoding/json"
)

type ActionView map[string]interface{}
type ActionArgs string

type DeployContract struct {
	Code string `json:"code"`
}

type FunctionCall struct {
	MethodName string     `json:"method_name"`
	Args       ActionArgs `json:"args"`
	Gas        *BigInt    `json:"gas"`
	Deposit    *BigInt    `json:"deposit"`
}

type Transfer struct {
	Deposit *BigInt `json:"deposit"`
}

type Stake struct {
	Stake     *BigInt `json:"stake"`
	PublicKey string  `json:"public_key"`
}

func (actionView *ActionView) IsDeployContract() bool {
	_, ok := (*actionView)["DeployContract"]
	return ok
}

func (actionView *ActionView) IsFunctionCall() bool {
	_, ok := (*actionView)["FunctionCall"]
	return ok
}

func (actionView *ActionView) IsTransfer() bool {
	_, ok := (*actionView)["Transfer"]
	return ok
}

func (actionView *ActionView) IsStake() bool {
	_, ok := (*actionView)["Stake"]
	return ok
}

func (actionView *ActionView) GetFunctionCall() *FunctionCall {
	if actionView.IsFunctionCall() {
		data, err := json.Marshal((*actionView)["FunctionCall"])
		if err != nil {
			return nil
		}
		fc := FunctionCall{}
		err = json.Unmarshal(data, &fc)
		return &fc
	}
	return nil
}

func (args *ActionArgs) Decode() (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(*args))
	if err != nil {
		return "{}", err
	}
	return string(decodedData), nil
}

func ConvertActionView(i interface{}) *ActionView {
	data, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	action := ActionView{}
	err = json.Unmarshal(data, &action)
	if err != nil {
		return nil
	}
	return &action
}
