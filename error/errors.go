package selferror

import "errors"

var (
	IllegalCodeError   = errors.New("illegal code")
	IllegalStateError  = errors.New("illegal state")
	NotHaveConfigError = errors.New("not have config")
)
