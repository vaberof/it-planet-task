package xerror

type ErrorWrapper struct {
	StatusCode int
	Message    string
	Err        error
}

func NewErrorWrapper(statusCode int, msg string, err error) error {
	return &ErrorWrapper{
		StatusCode: statusCode,
		Message:    msg,
		Err:        err,
	}
}

func (err *ErrorWrapper) Error() string {
	return err.Message
}
