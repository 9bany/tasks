// Code generated by MockGen. DO NOT EDIT.
// Source: http_client/iframely_client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIframelyClient is a mock of IframelyClient interface.
type MockIframelyClient struct {
	ctrl     *gomock.Controller
	recorder *MockIframelyClientMockRecorder
}

// MockIframelyClientMockRecorder is the mock recorder for MockIframelyClient.
type MockIframelyClientMockRecorder struct {
	mock *MockIframelyClient
}

// NewMockIframelyClient creates a new mock instance.
func NewMockIframelyClient(ctrl *gomock.Controller) *MockIframelyClient {
	mock := &MockIframelyClient{ctrl: ctrl}
	mock.recorder = &MockIframelyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIframelyClient) EXPECT() *MockIframelyClientMockRecorder {
	return m.recorder
}

// FetchURL mocks base method.
func (m *MockIframelyClient) FetchURL(context context.Context, apikey, url string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchURL", context, apikey, url)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchURL indicates an expected call of FetchURL.
func (mr *MockIframelyClientMockRecorder) FetchURL(context, apikey, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchURL", reflect.TypeOf((*MockIframelyClient)(nil).FetchURL), context, apikey, url)
}
