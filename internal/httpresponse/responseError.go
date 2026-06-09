package httpresponse

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

/*
* omitempty কী?
* field empty হলে JSON-এ দেখাবে না
 */

// func New(code int, message string, details string) *Error {
// 	return &ResponseError{
// 		Code:    code,
// 		Message: message,
// 		Details: details,
// 	}
// }

func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// httperror.New(404, "User Not Found")

func NewWithDetails(code int, message string, details string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// httperror.NewWithDetails(
// 	400,
// 	"Validation Error",
// 	"Email already exists",
// )