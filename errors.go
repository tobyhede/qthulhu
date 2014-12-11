package qthulhu

type NotFoundError struct {
	msg string
	Key []byte
}

func (e *NotFoundError) Error() string { return e.msg }

func NewNotFoundError(key []byte) *NotFoundError {
	return &NotFoundError{msg: "Key Not Found", Key: key}
}
