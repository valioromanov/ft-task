package calculator

import (
	"encoding/json"
	"fmt"
)

type ExpressionRequest struct {
	Expression string `json:"expression"`
}

type ExpressionResponse struct {
	Result int `json:"result"`
}

func (er ExpressionResponse) ToExpressionResponse(i interface{}) (ExpressionResponse, error) {
	var resp ExpressionResponse
	switch t := i.(type) {
	case int:
		resp.Result = t
	default:
		return ExpressionResponse{}, fmt.Errorf("not implemented type for casting ToExpressionResponse")
	}

	return resp, nil
}

func (er ExpressionResponse) ToString() string {
	str, err := json.Marshal(er)
	if err != nil {
		fmt.Println(err.Error()) // TODO add to logger
	}
	return string(str)
}
