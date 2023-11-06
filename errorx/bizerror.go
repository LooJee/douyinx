package errorx

import (
	"fmt"
	"github.com/loojee/douyinx/types"
)

type BizError struct {
	Code    int
	Message string
}

func (e *BizError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func CatchBizError(data types.CommonData) *BizError {
	if data.ErrorCode != 0 {
		return &BizError{
			Code:    data.ErrorCode,
			Message: data.Description,
		}
	}

	return nil
}
