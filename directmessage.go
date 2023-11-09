package douyinx

import (
	"context"
	"fmt"
	"github.com/loojee/douyinx/pkg/constants"
	"github.com/loojee/douyinx/pkg/errorx"
	"github.com/loojee/douyinx/pkg/traffic"
	"github.com/loojee/douyinx/types"
)

type DirectMessage struct {
	accessToken *AccessToken
}

func NewDirectMessage(accessToken *AccessToken) *DirectMessage {
	return &DirectMessage{
		accessToken: accessToken,
	}
}

func (im *DirectMessage) Send(ctx context.Context, openId string, msg types.DirectMessageReq) (string, error) {
	var resp types.DirectMessageResp

	token, err := im.accessToken.GetAccessToken(ctx, openId)
	if err != nil {
		return "", err
	}
	fmt.Println("token: ", token)
	fmt.Printf("body: %+v", msg)
	err = traffic.PostJSON(ctx, constants.UriDirectMessage, msg, &resp,
		traffic.WithAccessTokenHeader(token),
		traffic.WithOpenIdQueryParam(openId))

	if err != nil {
		return "", err
	}

	if err := errorx.CatchBizError(resp.Data); err != nil {
		return "", err
	}

	return resp.MsgId, nil
}
