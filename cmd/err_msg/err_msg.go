package err_msg

import "errors"

// ----- custom error msg
var (
	ErrTaskNotFound = errors.New("task not found")
	ErrBadRequest   = errors.New("incorrect data format")
	ErrInternal     = errors.New("internal server error")
)
