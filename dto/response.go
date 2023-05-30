package dto

type ResponseMeta struct {
	Success      bool   `json:"success"`
	MessageTitle string `json:"messageTitle"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
}

type ErrorResponse struct {
	ResponseMeta
	Data   any `json:"data"`
	Errors any `json:"errors,omitempty"`
}

func DefaultErrorResponse() *ErrorResponse {
	return DefaultErrorResponseWithMessage("")
}

func DefaultErrorResponseWithMessage(msg string) *ErrorResponse {
	return &ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      msg,
			ResponseTime: "",
		},
		Data: nil,
	}
}

func DefaultErrorInvalidDataWithMessage(msg string) *ErrorResponse {
	return &ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      "Form Invalid data.",
			ResponseTime: "",
		},
		Data: msg,
	}
}

func DefaultDataInvalidResponse(validationErrors any) *ErrorResponse {
	return &ErrorResponse{
		ResponseMeta: ResponseMeta{
			MessageTitle: "Oops, something went wrong.",
			Message:      "Data invalid.",
		},
		Errors: validationErrors,
	}
}

func DefaultBadRequestResponse() *ErrorResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}

type Response struct {
	ResponseMeta
	Data any `json:"data"`
}

func DefaultInvalidInputResponse(errs map[string][]string) *Response {
	return &Response{
		ResponseMeta: ResponseMeta{
			Success: false,
			Message: "invalid data",
		},

		Data: errs,
	}
}

func NewSuccessResponse(data any, msg string, resTime string) *Response {
	return &Response{
		ResponseMeta: ResponseMeta{
			Success:      true,
			Message:      msg,
			MessageTitle: "Success",
			ResponseTime: resTime,
		},
		Data: data,
	}
}
