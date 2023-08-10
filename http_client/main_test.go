package httpclient

import (
	"os"
	"testing"

	"github.com/9bany/task/util"
)

var testRequest *IframelyRequestor

func TestMain(m *testing.M) {
	config := util.LoadConfig()
	testRequest = New(config.IframeURL)
	os.Exit(m.Run())
}
