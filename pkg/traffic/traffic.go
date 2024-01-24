package traffic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/loojee/douyinx/config"
	"net/http"
	"sync"
)

var (
	c *client
	o sync.Once
)

type client struct {
	client *resty.Client
	conf   *config.Config

	httpClient *http.Client
}

type Option func(c *client)

func WithClient(httpClient *http.Client) Option {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

func MustInit(conf *config.Config, options ...Option) {
	if c == nil {
		o.Do(func() {
			c = &client{
				httpClient: http.DefaultClient,
				conf:       conf,
			}

			for _, option := range options {
				option(c)
			}

			c.client = resty.NewWithClient(c.httpClient)
		})
	}
}

func genUrl(path string) string {
	return c.conf.DouYinHost + path
}

type RequestOption func(*resty.Request)

func WithHeaders(headers map[string]string) func(*resty.Request) {
	return func(r *resty.Request) {
		r.SetHeaders(headers)
	}
}

func WithQueryParams(params map[string]string) func(*resty.Request) {
	return func(r *resty.Request) {
		r.SetQueryParams(params)
	}
}

func WithAccessTokenHeader(token string) func(*resty.Request) {
	return func(r *resty.Request) {
		r.SetHeader("access-token", token)
	}
}

func WithOpenIdQueryParam(openId string) func(*resty.Request) {
	return func(r *resty.Request) {
		r.SetQueryParam("open_id", openId)
	}
}

func PostJSON(ctx context.Context, path string, req any, resp any, options ...RequestOption) error {
	r := c.client.R()

	for _, option := range options {
		option(r)
	}

	rsp, err := r.SetContext(ctx).SetBody(req).Post(genUrl(path))
	if err != nil {
		return err
	}

	fmt.Println(rsp.Request.URL)
	fmt.Println(rsp.String())

	return json.Unmarshal(rsp.Body(), &resp)
}

func PostUrlEncodeForm(ctx context.Context, path string, req Former, resp any) error {
	rsp, err := c.client.R().SetContext(ctx).SetFormData(req.IntoForm()).Post(genUrl(path))
	if err != nil {
		return err
	}

	fmt.Println(rsp.Request.URL)
	fmt.Println(rsp.String())

	return json.Unmarshal(rsp.Body(), &resp)
}

func Get(ctx context.Context, path string, resp any, options ...RequestOption) error {
	r := c.client.R()

	for _, option := range options {
		option(r)
	}

	rsp, err := r.SetContext(ctx).Get(genUrl(path))
	if err != nil {
		return err
	}

	fmt.Println(rsp.Request.URL)
	fmt.Println(rsp.String())

	return json.Unmarshal(rsp.Body(), &resp)
}

func UploadImage(ctx context.Context, path string, form MultiPartForm, resp any, options ...RequestOption) error {
	r := c.client.R()

	for _, option := range options {
		option(r)
	}

	rsp, err := r.SetContext(ctx).
		SetFileReader(form.Param, form.Filename, form.Reader).
		Post(genUrl(path))
	if err != nil {
		return err
	}

	fmt.Println(rsp.Request.URL)

	return json.Unmarshal(rsp.Body(), &resp)
}
