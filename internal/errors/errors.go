package errors

import (
	"fmt"
)

type RecordNotFoundError struct {
	msg string
}

func (err RecordNotFoundError) Error() string {
	return err.msg
}

func NewRecordNotFoundError(id int64) error {
	return &RecordNotFoundError{fmt.Sprintf("record{ID:%d} was not found", id)}
}
