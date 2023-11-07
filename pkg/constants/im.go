package constants

type IMMsgType int

const (
	IMMsgTypeText IMMsgType = iota + 1
	IMMsgTypeImage
	IMMsgTypeVideo
	IMMsgTypeRetainConsultCard      IMMsgType = iota + 5 // 留资卡片
	IMMsgTypeGroupInvitation                             // 群聊邀请卡片
	IMMsgTypeAppletCard                                  // 小程序引导卡片
	IMMsgTypeAppletCoupon                                // 小程序券
	IMMsgTypeAuthPrivateMessageCard                      // 服务私信授权卡片
)
