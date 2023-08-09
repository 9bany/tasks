package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

type SiteType string

const (
	YOUTUBE_SITE_TYPE SiteType = "youtube"
	DEFAULT_SITE_TYPE SiteType = "default"
)

func getSiteTypeFromURL(urlString string) SiteType {
	if strings.HasPrefix(urlString, "https://www.youtube.com") {
		return YOUTUBE_SITE_TYPE
	}
	return DEFAULT_SITE_TYPE
}

type siteDataFactoryOperation func(data []byte) ([]byte, error)

func siteDataFactory(data []byte, options ...siteDataFactoryOperation) (result []byte, err error) {
	result = data
	for _, op := range options {
		result, err = op(data)
		if err != nil {
			return
		}
	}
	return
}

func embedDataIframelyUrlOps(data []byte) ([]byte, error) {
	var mapData map[string]interface{}
	err := json.Unmarshal(data, &mapData)
	if err != nil {
		return nil, fmt.Errorf("embedDataIframelyUrlOps: can not unmarsahl data %s", err.Error())
	}

	htmlString, ok := mapData["html"].(string)
	if !ok {
		return nil, fmt.Errorf("embedYoutubeVideoIDOps: can not get html in data %s", err.Error())
	}

	if strings.Contains(htmlString, "data-iframely-url") {
		mapData["data-iframely-url"] = true
	}
	return json.Marshal(mapData)
}

func embedYoutubeVideoIDOps(data []byte) ([]byte, error) {
	var mapData map[string]interface{}
	err := json.Unmarshal(data, &mapData)
	if err != nil {
		return nil, fmt.Errorf("embedYoutubeVideoIDOps: can not unmarsahl data %s", err.Error())
	}
	urlString, ok := mapData["url"].(string)
	if !ok {
		return nil, fmt.Errorf("embedYoutubeVideoIDOps: can not get url in data %s", err.Error())
	}
	switch getSiteTypeFromURL(urlString) {
	case YOUTUBE_SITE_TYPE:
		url, err := url.ParseRequestURI(urlString)
		if err != nil {
			return nil, fmt.Errorf("can not operate data before return: %s", err.Error())
		}
		videoID := url.Query().Get("v")
		if len(videoID) == 0 {
			log.Println("videoId not exist in the url")
			return data, nil
		}
		var mapData map[string]interface{}

		err = json.Unmarshal(data, &mapData)
		if err != nil {
			return nil, fmt.Errorf("can not operate data before return: %s", err.Error())
		}

		mapData["youtube_video_id"] = videoID
		return json.Marshal(mapData)

	default:
		return data, nil
	}
}
