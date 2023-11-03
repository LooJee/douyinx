package traffic

import (
	"context"
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/constants"
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
