package customErrors

type MyAppError struct {
	Message string
	ErrCode
	Err error // For error chains
}

// Prints the root cause
func (err *MyAppError) Error() string {
	return err.Err.Error()
}

// Unwrap prints bottom of the error chain error
func (err *MyAppError) Unwrap() error {
	return err.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{Message: message, Err: err, ErrCode: code}
}
