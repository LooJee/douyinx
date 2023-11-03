package types

import "fmt"

type CommonData struct {
	LogId       string `json:"log_id"`
	ErrorCode   int    `json:"error_code"`
	Captcha     string `json:"captcha"`
	DescUrl     string `json:"desc_url"`
	Description string `json:"description"`
}

func (d *CommonData) Error() string {
	return fmt.Sprintf("logid:%s, error_code:%d, description:%s", d.LogId, d.ErrorCode, d.Description)
}

type ExtraData struct {
	LogId string `json:"logid"`
	Now   int64  `json:"now"`
}
