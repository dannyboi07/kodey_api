package schema

import (
	"fmt"
	"net/http"
)

type ExecuteCode struct {
	Code     *string `json:"code"`
	Language *string `json:"language"`
}

func (e *ExecuteCode) Validate() (int, error) {
	switch {
	case e.Code == nil:
		return http.StatusBadRequest, fmt.Errorf("Field 'code' is empty")
	case e.Language == nil:
		return http.StatusBadRequest, fmt.Errorf("Field 'Language' is empty")
	}

	return 0, nil
}

type ExecuteCodeResponse struct {
	Result   string `json:"result"`
	Language string `json:"language"`
}
