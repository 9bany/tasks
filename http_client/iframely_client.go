package httpclient

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/resty.v1"
)

type IframelyClient interface {
	FetchURL(context context.Context, apikey, url string) ([]byte, error)
}

// https://github.com/uber-go/guide/issues/25
var _ IframelyClient = (*IframelyRequestor)(nil)

type IframelyRequestor struct {
	client *resty.Client
}

func New(hostUrl string) *IframelyRequestor {
	client := resty.New()
	client.SetHostURL(hostUrl)
	client.SetTimeout(30 * time.Second)

	return &IframelyRequestor{
		client: client,
	}
}

func (r *IframelyRequestor) FetchURL(ctx context.Context, apikey, url string) ([]byte, error) {
	r.client.SetQueryParam("url", url)
	r.client.SetQueryParam("api_key", apikey)
	resp, err := r.client.R().SetContext(ctx).Get("/api/oembed")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return resp.Body(), fmt.Errorf("status code not success")
	}

	data := resp.Body()
	if data == nil {
		return nil, fmt.Errorf("data response is empty")
	}
	return data, nil
}
