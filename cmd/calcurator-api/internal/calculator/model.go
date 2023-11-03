package calculator

import (
	"fmt"
)

type ExpressionRequest struct {
	Expression string `json:"expression"`
}

type ExpressionResponse struct {
	Result int `json:"result"`
}

type ValidateResponse struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

type InvalidData struct {
	Frequency int
	Type      string
}

type InvalidKey struct {
	Endpoint   string
	Expression string
}

type InvalidExpression map[InvalidKey]InvalidData

type GetErrorResponse struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	Type       string `json:"type"`
}

func (ie InvalidExpression) ToGetErrorsResponse() []GetErrorResponse {
	errResponses := make([]GetErrorResponse, 0)
	for key, val := range ie {
		errResp := GetErrorResponse{
			Expression: key.Expression,
			Endpoint:   key.Endpoint,
			Frequency:  val.Frequency,
			Type:       val.Type,
		}
		errResponses = append(errResponses, errResp)
	}

	return errResponses
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
