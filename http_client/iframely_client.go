package httpclient

import (
	"context"
	"fmt"

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

	return &IframelyRequestor{
		client: client,
	}
}

func (r *IframelyRequestor) FetchURL(context context.Context, apikey, url string) ([]byte, error) {
	resp, err := r.client.R().SetPathParams(map[string]string{
		"url":     url,
		"api_key": apikey,
	}).Get("/api/oembed")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return resp.Body(), fmt.Errorf("can not fetch the url: %s", url)
	}

	data := resp.Body()
	if data == nil {
		return nil, fmt.Errorf("data response is empty")
	}

	// ops after request
	// ops data outcome with url type
	return siteOperateData(url, data)
}
