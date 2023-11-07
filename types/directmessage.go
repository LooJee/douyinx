package types

import "github.com/loojee/douyinx/pkg/constants"

type DirectMessageReq struct {
	MsgId          string    `json:"msg_id"`
	ConversationId string    `json:"conversation_id"`
	ToUserId       string    `json:"to_user_id"`
	Scene          string    `json:"scene"`
	Content        IMContent `json:"content"`
}

type IMContent struct {
	MsgType constants.IMMsgType `json:"msg_type"`

	Text  *IMContentText  `json:"text,omitempty"`
	Image *IMContentImage `json:"image,omitempty"`
}

type IMContentText struct {
	Text string `json:"text"`
}

type IMContentImage struct {
	MediaId string `json:"media_id"`
}

type DirectMessageResp struct {
	MsgId string     `json:"msg_id"`
	Data  CommonData `json:"data"`
	Extra ExtraData  `json:"extra"`
}
