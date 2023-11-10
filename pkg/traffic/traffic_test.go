package traffic

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/pkg/constants"
	"github.com/loojee/douyinx/types"
	"os"
	"testing"
)

func TestPostJSON(t *testing.T) {
	rsp := types.GetClientTokenResp{}
	MustInit(config.NewConfig("", ""))

	if err := PostJSON(context.Background(), constants.UriClientToken, types.GetClientTokenReq{
		ClientKey:    os.Getenv("client_key"),
		ClientSecret: os.Getenv("client_secret"),
		GrantType:    constants.GrantTypeClientCredential,
	}, &rsp); err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.Data)
}

func TestUploadImage(t *testing.T) {
	rsp := types.ImageUploadRsp{}
	MustInit(config.NewConfig("", ""))

	imageRsp, err := resty.New().R().SetDoNotParseResponse(true).Get(os.Getenv("url"))
	if err != nil {
		t.Fatal(err)
	}

	defer imageRsp.RawBody().Close()

	if err := UploadImage(context.Background(), constants.UriImageUpload, MultiPartForm{
		Param:    "image",
		Filename: "image.JPEG",
		Reader:   imageRsp.RawBody(),
	}, &rsp, WithAccessTokenHeader(os.Getenv("client_token"))); err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.ImageUploadRspData)
}
