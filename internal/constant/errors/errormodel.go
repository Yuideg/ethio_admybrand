package errors

import "errors"
type ErrorModel struct {
	ErrorCode        int      `json:"error_code,omitempty"`
	ErrorDescription string   `json:"error_description,omitempty"`
	ErrorMessage     string    `json:"error_message,omitempty"`
	ErrorDetail      []string `json:"field_errors,omitempty"`
	ErrorStack       string   `json:"error_details,omitempty"`
}

func ErrorService(err error, e ...error) *ErrorModel {
	if err == nil {
		return nil
	}
	if len(e) != 0 {
		return &ErrorModel{
			ErrorMessage:     err.Error(),
			ErrorDescription: Descriptions[err],
			ErrorCode:        ErrCodes[err],
			ErrorStack:       e[0].Error(),
		}
	}
	return &ErrorModel{
		ErrorMessage:     err.Error(),
		ErrorDescription: Descriptions[err],
		ErrorCode:        ErrCodes[err],
	}
}
func NewError(txt string) error {
	return errors.New(txt)
}

func HttpStatus(er *ErrorModel) int {
	return StatusCodes[NewError(er.ErrorMessage)]
}
