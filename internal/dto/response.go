package dto

type Response struct {
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message, data interface{}) Response {

	return Response{
		Message: message,
		Data:    data,
		Error:   nil,
	}
}

func ErrorResponse(err, data interface{}) Response {

	return Response{
		Message: nil,
		Data:    data,
		Error:   err,
	}
}
