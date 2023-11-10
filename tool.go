package douyinx

import (
	"context"
	"github.com/loojee/douyinx/pkg/constants"
	"github.com/loojee/douyinx/pkg/errorx"
	"github.com/loojee/douyinx/pkg/traffic"
	"github.com/loojee/douyinx/types"
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
