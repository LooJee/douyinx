package douyinx

import (
	"encoding/json"
	"fmt"
	"github.com/loojee/douyinx/types"
	"io"
	"net/http"
)

type Webhook struct {
}

func NewWebhook() *Webhook {
	return &Webhook{}
}

func (w *Webhook) HandleMsg(r *http.Request) (*types.WebHookEvent, error) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("receive webhook msg: ", string(body))

	data := types.WebHookEvent{}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
