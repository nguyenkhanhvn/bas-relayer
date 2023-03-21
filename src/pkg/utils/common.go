package utils

type ErrorType int

const (
	// Rpc Client error
	ClientError = iota

	// Other error
	OtherError
)

type Error interface {
	ErrorType() ErrorType
	ErrorString() error
}

func NewError(Type ErrorType, err error) Error {
	return &errorStruct{Type, err}
}

type errorStruct struct {
	errorType   ErrorType
	errorString error
}

func (e *errorStruct) ErrorType() ErrorType {
	return e.errorType
}

func (e *errorStruct) ErrorString() error {
	return e.errorString
}
