package httpclient

import (
	"log"
	"os"
	"testing"

	"github.com/9bany/task/util"
)

var testRequest *IframelyRequestor

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatalln("Can not load config file", err)
	}
	testRequest = New(config.IframeURL)
	os.Exit(m.Run())
}
