package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/9bany/task/db/mock"
	db "github.com/9bany/task/db/sqlc"
	mockHttpclient "github.com/9bany/task/http_client/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const testUrl = "https://vnexpress.net/khu-vuc-lu-quet-o-ha-noi-co-nhieu-cong-trinh-xay-trai-phep-4638945.html"

func TestGetSite(t *testing.T) {

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockQuerier, mockHttp *mockHttpclient.MockIframelyClient)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "empty url",
			url:  "",
			buildStubs: func(store *mockdb.MockQuerier, mockHttp *mockHttpclient.MockIframelyClient) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "url error format",
			url:  "error format",
			buildStubs: func(store *mockdb.MockQuerier, mockHttp *mockHttpclient.MockIframelyClient) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "can not get site with db error",
			url:  testUrl,
			buildStubs: func(store *mockdb.MockQuerier, mockHttp *mockHttpclient.MockIframelyClient) {
				store.EXPECT().GetSiteByURL(gomock.Any(), gomock.Any()).Times(1).Return(db.Sites{}, sql.ErrConnDone)
				//
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "renew site: get api key with some error",
			url:  testUrl,
			buildStubs: func(store *mockdb.MockQuerier, mockHttp *mockHttpclient.MockIframelyClient) {
				store.EXPECT().GetSiteByURL(gomock.Any(), gomock.Any()).Times(1).Return(db.Sites{}, sql.ErrNoRows)
				store.EXPECT().GetRandomKey(gomock.Any()).Times(1).Return(db.Keys{}, fmt.Errorf("some error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "renew site: fetch site medata error",
			url:  testUrl,
			buildStubs: func(store *mockdb.MockQuerier, mockHttp *mockHttpclient.MockIframelyClient) {
				store.EXPECT().GetSiteByURL(gomock.Any(), gomock.Any()).Times(1).Return(db.Sites{}, sql.ErrNoRows)
				store.EXPECT().GetRandomKey(gomock.Any()).Times(1).Return(db.Keys{Key: "api_key"}, nil)

				mockHttp.EXPECT().FetchURL(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return([]byte{}, fmt.Errorf("some error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "renew site: save new site error",
			url:  testUrl,
			buildStubs: func(store *mockdb.MockQuerier, mockHttp *mockHttpclient.MockIframelyClient) {
				store.EXPECT().GetSiteByURL(gomock.Any(), gomock.Any()).Times(1).Return(db.Sites{}, sql.ErrNoRows)
				store.EXPECT().GetRandomKey(gomock.Any()).Times(1).Return(db.Keys{Key: "api_key"}, nil)
				
				mapData := map[string]interface{}{}
				mapData["url"] = testUrl
				mapData["html"] = "html"

				data, err := json.Marshal(mapData)
				if err != nil {
					log.Panicln("can not marshal data")
				}

				mockHttp.EXPECT().FetchURL(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(data, nil)

				store.EXPECT().IncreaseKeyUsageCount(gomock.Any(), gomock.Any()).Times(1)

				store.EXPECT().CreateSite(gomock.Any(), gomock.Any()).Times(1).Return(db.Sites{}, fmt.Errorf("some error"))

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "renew site: save new oke",
			url:  testUrl,
			buildStubs: func(store *mockdb.MockQuerier, mockHttp *mockHttpclient.MockIframelyClient) {
				store.EXPECT().GetSiteByURL(gomock.Any(), gomock.Any()).Times(1).Return(db.Sites{}, sql.ErrNoRows)
				store.EXPECT().GetRandomKey(gomock.Any()).Times(1).Return(db.Keys{Key: "api_key"}, nil)

				mapData := map[string]interface{}{}
				mapData["url"] = testUrl
				mapData["html"] = "html"

				data, err := json.Marshal(mapData)
				if err != nil {
					log.Panicln("can not marshal data")
				}

				mockHttp.EXPECT().FetchURL(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(data, nil)

				store.EXPECT().IncreaseKeyUsageCount(gomock.Any(), gomock.Any()).Times(1)

				store.EXPECT().CreateSite(gomock.Any(), gomock.Any()).Times(1).Return(db.Sites{}, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockQuerier(ctrl)

			httpClient := mockHttpclient.NewMockIframelyClient(ctrl)

			// build stubs
			tc.buildStubs(store, httpClient)
			// start test server and send request
			server := newTestServer(t, store, httpClient)
			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/api/meta", nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("url", tc.url)
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			// check response
			tc.checkResponse(t, recorder)
		})
	}
}
