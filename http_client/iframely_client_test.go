package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestSucesssWithNilBody(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://iframe.ly/api/oembed",
		func(req *http.Request) (*http.Response, error) {
			return nil, nil
		})

	data, err := testRequest.FetchURL(context.Background(), "apikey", "https://vnexpress.net")
	require.Nil(t, data)
	require.Error(t, err)

}

func TestRequestFail(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://iframe.ly/api/oembed",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(500, nil)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		})

	_, err := testRequest.FetchURL(context.Background(), "apikey", "https://vnexpress.net")
	require.Error(t, err)

}

func TestSucesssRequest(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	articles := map[string]interface{}{
		"key0": "value0",
	}

	httpmock.RegisterResponder("GET", "https://iframe.ly/api/oembed",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, articles)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		})

	data, err := testRequest.FetchURL(context.Background(), "apikey", "https://vnexpress.net")
	require.NoError(t, err)
	require.NotEmpty(t, data)

	var data0 map[string]interface{}
	json.Unmarshal(data, &data0)
	require.Equal(t, articles, data0)
}
