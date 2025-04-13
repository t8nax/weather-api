package common

type AppErrorCode int

const (
	CodeInvalidLocation AppErrorCode = iota
	CodeInvalidDate
	CodeServiceError
)

type AppError struct {
	Code AppErrorCode
	Err  error
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(code AppErrorCode, err error) *AppError {
	return &AppError{
		Code: code,
		Err:  err,
	}
}
