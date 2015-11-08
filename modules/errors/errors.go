package errors

type errorContainer struct {
	code    string
	message string
}

// New returns an error.
func New(code, message string) error {
	return &errorContainer{
		code:    code,
		message: message,
	}
}

// Error returns an error message.
func (e *errorContainer) Error() string {
	return e.message
}

// Code returns an error code.
func (e *errorContainer) Code() string {
	return e.code
}

// Message returns an error message.
func (e *errorContainer) Message() string {
	return e.message
}
