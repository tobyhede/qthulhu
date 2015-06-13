package qthulhu

type NotFoundError struct {
	msg string
	Key uint64
}

func (e *NotFoundError) Error() string { return e.msg }

func NewNotFoundError(key uint64) *NotFoundError {
	return &NotFoundError{msg: "not found", Key: key}
}
