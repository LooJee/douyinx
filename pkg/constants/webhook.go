package constants

type WebHookEventType string

const (
	WebHookEventTypeVerifyWebhook         WebHookEventType = "verify_webhook"           // 认证事件
	WebHookEventTypeCreateVideo           WebHookEventType = "create_video"             // 用户使用开发者应用分享视频到抖音
	WebHookEventTypeAuthorize             WebHookEventType = "authorize"                // 抖音用户授权给开发者App
	WebHookEventTypeUnauthorize           WebHookEventType = "unauthorize"              // 抖音用户解除授权，推送事件给开发者APP
	WebHookEventTypeImReceiveMsg          WebHookEventType = "im_receive_msg"           // 接收私信，用户收到私信触发
	WebHookEventTypeImSendMsg             WebHookEventType = "im_send_msg"              // 发送私信，用户发送私信触发
	WebHookEventTypeImEnterDirectMsg      WebHookEventType = "im_enter_direct_msg"      // 接收用户进入私信会话页事件，用户主动进入私信会话页触发
	WebHookEventTypeImGroupReceiveMsg     WebHookEventType = "im_group_receive_msg"     // 接收群消息事件
	WebHookEventTypeImGroupSendMsg        WebHookEventType = "im_group_send_msg"        // 发送群消息事件
	WebHookEventTypeEnterGroupAuditChange WebHookEventType = "enter_group_audit_change" // 用户加群申请
	WebHookEventTypeGroupFansEvent        WebHookEventType = "group_fans_event"         // 用户加群成功
	WebHookEventTypeContractAuthorize     WebHookEventType = "contract_authorize"       // 用户给应用经营关系（scope名为xxx.bind，或私信相关能力）授权
	WebHookEventTypeContractUnauthorize   WebHookEventType = "contract_unauthorize"     // 用户解除应用经营关系（scope名为xxx.bind，或私信相关能力）授权
	WebHookEventTypeNewVideoDigg          WebHookEventType = "new_video_digg"           // 接收用户点赞事件
	WebHookEventTypeNewFollowAction       WebHookEventType = "new_follow_action"        // 接收用户关注事件
)

type WebHookMsgType string

const (
	WebHookMsgTypeText              WebHookMsgType = "text"                // 文本
	WebHookMsgTypeImage             WebHookMsgType = "image"               // 图片
	WebHookMsgTypeEmoji             WebHookMsgType = "emoji"               // 表情
	WebHookMsgTypeVideo             WebHookMsgType = "video"               // 视频
	WebHookMsgTypeRetainConsultCard WebHookMsgType = "retain_consult_card" // 留资卡片
	WebHookMsgTypeOther             WebHookMsgType = "other"               // 其他 不同类型消息参数见下方介绍
)

type WebHookConversationType int

const (
	WebHookConversationTypeDirect WebHookConversationType = iota + 1 // 私信
	WebHookConversationTypeGroup                                     // 群组
)

type WebHookSceneType int

const (
	WebHookSceneTypeKeyword  WebHookSceneType = iota + 1 // 关键词自动回复
	WebHookSceneTypeWelcome                              // 欢迎语
	WebHookSceneTypeVideo                                // 视频页
	WebHookSceneTypeMainPage                             // 主页
	WebHookSceneTypeMsgPage                              // 消息页
	WebHookSceneTypeLivePage                             // 直播页
	WebHookSceneTypeOther    WebHookSceneType = 100      // 其它
)
