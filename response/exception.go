package response

import (
	"data-pipeline/response/errcode"
)

// Exception :
type Exception struct {
	Code    errcode.Code
	Message string
}
