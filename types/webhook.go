package types

import (
	"github.com/loojee/douyinx/pkg/constants"
	"github.com/mitchellh/mapstructure"
)

type WebHookEvent struct {
	Event      constants.WebHookEventType `json:"event"`
	ClientKey  string                     `json:"client_key"`
	FromUserId string                     `json:"from_user_id"`
	ToUserId   string                     `json:"to_user_id"`
	Content    any                        `json:"content"`
	LogId      string                     `json:"log_id"`
}

func (w WebHookEvent) Challenge() (WebHookContentChallenge, error) {
	var data WebHookContentChallenge
	err := mapstructure.Decode(w.Content, &data)

	return data, err
}

func (w WebHookEvent) ImMessage() (WebHookContentImMessage, error) {
	var data WebHookContentImMessage
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Squash: true,
		Result: &data,
	})
	if err != nil {
		return data, err
	}

	err = decoder.Decode(w.Content)

	return data, err
}

type WebHookContentChallenge struct {
	Challenge int `json:"challenge" mapstructure:"challenge"`
}

type WebHookContentCreateVideo struct {
	ItemId  string `json:"item_id"`
	ShareId string `json:"share_id"`
}

type WebHookContentAuthorize struct {
	Scopes []string `json:"scopes"`
}

type WebHookContentUnAuthorize struct {
	Scopes []string `json:"scopes"`
}

type WebHookContentContractAuthorize struct {
	Scopes    []string `json:"scopes"`
	Timestamp int64    `json:"timestamp"`
}

type WebHookContentContractUnAuthorize struct {
	Scopes    []string `json:"scopes"`
	Timestamp int64    `json:"timestamp"`
}

type WebHookContentNewVideoDigg struct {
	ActionType int    `json:"action_type"`
	ActionTime int    `json:"action_time"`
	ItemId     string `json:"item_id"`
}

type WebHookContentNewFollowAction struct {
	ActionType int `json:"action_type"`
	ActionTime int `json:"action_time"`
}

type WebHookContentImEnterDirectMsgUserInfo struct {
	OpenId   string `json:"open_id" mapstructure:"open_id"`
	NickName string `json:"nick_name" mapstructure:"nick_name"`
	Avatar   string `json:"avatar" mapstructure:"avatar"`
}

type WebHookContentImEnterDirectMsgAddInfo struct {
	AdvId           int    `json:"adv_id"`            // 广告主id
	AdvName         string `json:"adv_name"`          // 广告主名称
	AdId            int    `json:"ad_id"`             // 广告计划id
	AdName          string `json:"ad_name"`           // 广告计划名称
	CreativeId      int    `json:"creative_id"`       // 创意id
	PromotionId     int    `json:"promotion_id"`      // 广告id
	MaterialTitleId int    `json:"material_title_id"` // 广告素材-标题id
	MaterialImageId int    `json:"material_image_id"` // 广告素材-图片id
	MaterialVideoId int    `json:"material_video_id"` // 广告素材-视频id
}

// WebHookContentImEnterDirectMsg 接收用户进入私信会话页事件内容
type WebHookContentImEnterDirectMsg struct {
	ConversationShortId string                                   `json:"conversation_short_id"`
	ServerMessageId     string                                   `json:"server_message_id"`
	ConversationType    constants.WebHookConversationType        `json:"conversation_type"`
	CreateTime          int64                                    `json:"create_time"`
	SceneType           constants.WebHookSceneType               `json:"scene_type"`
	UserInfos           []WebHookContentImEnterDirectMsgUserInfo `json:"user_infos"`
	DataImExtra         string                                   `json:"data-im-extra"`
	AdInfo              WebHookContentImEnterDirectMsgAddInfo    `json:"ad_info"`
}

// WebHookContentImMessage im消息事件内容
type WebHookContentImMessage struct {
	ConversationShortId string                                   `json:"conversation_short_id" mapstructure:"conversation_short_id"` // 会话id
	ServerMessageId     string                                   `json:"server_message_id" mapstructure:"server_message_id"`         // 消息id
	ConversationType    constants.WebHookConversationType        `json:"conversation_type" mapstructure:"conversation_type"`         // 会话类型
	CreateTime          int64                                    `json:"create_time" mapstructure:"create_time"`                     // 消息创建时间，13位毫秒时间戳
	UserInfos           []WebHookContentImEnterDirectMsgUserInfo `json:"user_infos" mapstructure:"user_infos"`
	Source              string                                   `json:"source" mapstructure:"source"`

	MessageType constants.WebHookMsgType `json:"message_type" mapstructure:"message_type"` // 消息类型

	// 文本消息/其它消息字段
	WebHookContentTextMessage `mapstructure:",squash"`

	// 表情消息字段
	WebHookContentEmojiMessage `mapstructure:",squash"`

	// 视频消息字段
	WebHookContentVideoMessage `mapstructure:",squash"`

	// 留资卡片消息字段
	WebHookContentRetainCardMessage `mapstructure:",squash"`
}

func (w WebHookContentImMessage) GetUser(openId string) (WebHookContentImEnterDirectMsgUserInfo, bool) {
	for _, user := range w.UserInfos {
		if user.OpenId == openId {
			return user, true
		}
	}
	return WebHookContentImEnterDirectMsgUserInfo{}, false
}

type WebHookContentTextMessage struct {
	Text string `json:"text" mapstructure:"text"`
}

type WebHookContentEmojiMessage struct {
	ResourceType   string `json:"resource_type" mapstructure:"resource_type"`     // 资源类型
	ResourceHeight int    `json:"resource_height" mapstructure:"resource_height"` // 资源高度
	ResourceWidth  int    `json:"resource_width" mapstructure:"resource_width"`   // 资源宽度
	ResourceUrl    string `json:"resource_url" mapstructure:"resource_url"`       // 资源链接
}

type WebHookContentVideoMessage struct {
	ItemId string `json:"item_id" mapstructure:"item_id"`
}

type WebHookContentRetainCardMessage struct {
	CardId     string                            `json:"card_id" mapstructure:"card_id"`         // 卡片id
	CardStatus int                               `json:"card_status" mapstructure:"card_status"` // 卡片状态
	CardData   []WebHookContentImMessageCardData `json:"card_data" mapstructure:"card_data"`     // 卡片数据
}

type WebHookContentImMessageCardData struct {
	Label string `json:"label" mapstructure:"label"`
	Value string `json:"value" mapstructure:"value"`
}
