package douyinx

import (
	"context"
	"github.com/loojee/douyinx/pkg/constants"
	"github.com/loojee/douyinx/pkg/errorx"
	"github.com/loojee/douyinx/pkg/traffic"
	"github.com/loojee/douyinx/types"
	"github.com/mitchellh/mapstructure"
)

type Tool struct {
	clientToken *ClientToken
}

func NewTool(clientToken *ClientToken) *Tool {
	return &Tool{clientToken: clientToken}
}

func (t *Tool) UploadImage(ctx context.Context, form traffic.MultiPartForm) (types.ImageUploadRspData, error) {
	rsp := types.ImageUploadRsp{}
	clientToken, err := t.clientToken.GetToken(ctx)
	if err != nil {
		return types.ImageUploadRspData{}, err
	}
	if err := traffic.UploadImage(ctx, constants.UriImageUpload, form, &rsp, traffic.WithAccessTokenHeader(clientToken)); err != nil {
		return types.ImageUploadRspData{}, err
	}

	if err := errorx.CatchBizError(rsp.Data); err != nil {
		return types.ImageUploadRspData{}, err
	}

	return rsp.ImageUploadRspData, nil
}

func (t *Tool) GetJsbTicket(ctx context.Context) (types.JSBTicket, error) {
	rsp := types.JSBTicketResp{}
	clientToken, err := t.clientToken.GetToken(ctx)
	if err != nil {
		return types.JSBTicket{}, err
	}
	if err := traffic.Get(ctx, constants.UriJsbTicket, &rsp, traffic.WithAccessTokenHeader(clientToken)); err != nil {
		return types.JSBTicket{}, err
	}

	if rsp.Data.ErrorCode != 0 {
		return types.JSBTicket{}, &errorx.BizError{
			Code:    rsp.Data.ErrorCode,
			Message: rsp.Data.Description,
		}
	}

	return rsp.Data, nil
}

func (t *Tool) GetOauthQrcode(ctx context.Context, req types.GetOauthQrcodeReq) (types.OauthQrcodeData, error) {
	rsp := types.GetOauthQrcodeResp{}

	params := make(map[string]string)

	req.ClientKey = t.clientToken.config.ClientKey

	if len(req.JumpType) == 0 {
		req.JumpType = "native"
	}

	if len(req.CustomizeParams) == 0 {
		req.CustomizeParams = "{\"comment_id\":\"\",\"source\":\"pc_auth\",\"not_skip_confirm\":\"true\",\"enter_from\":\"\"}"
	}

	if err := mapstructure.Decode(&req, &params); err != nil {
		return types.OauthQrcodeData{}, err
	}

	if err := traffic.Get(ctx, constants.UriOauthQrcode, &rsp, traffic.WithQueryParams(params)); err != nil {
		return types.OauthQrcodeData{}, err
	}

	if err := errorx.CatchBizError(rsp.Data.CommonData); err != nil {
		return types.OauthQrcodeData{}, err
	}

	return rsp.Data, nil
}

func (t *Tool) CheckOauthQrcode(ctx context.Context, req types.CheckOauthQrcodeReq) (types.CheckOauthQrcodeData, error) {
	rsp := types.CheckOauthQrcodeResp{}

	params := make(map[string]string)

	req.ClientKey = t.clientToken.config.ClientKey

	if err := mapstructure.Decode(&req, &params); err != nil {
		return types.CheckOauthQrcodeData{}, err
	}

	if err := traffic.Get(ctx, constants.UriCheckOauthQrcode, &rsp, traffic.WithQueryParams(params)); err != nil {
		return types.CheckOauthQrcodeData{}, err
	}

	if err := errorx.CatchBizError(rsp.Data.CommonData); err != nil {
		return types.CheckOauthQrcodeData{}, err
	}

	return rsp.Data, nil
}
