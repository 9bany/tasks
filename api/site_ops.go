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
		result, err = op(result)
		if err != nil {
			return nil, err
		}
	}
	return
}

func embedDataIframelyUrlOps(data []byte) ([]byte, error) {
	var mapData map[string]interface{}

	err := json.Unmarshal(data, &mapData)
	if err != nil {
		return nil, fmt.Errorf("can not unmarsahl data %s", err.Error())
	}

	htmlString, ok := mapData["html"].(string)
	if !ok {
		return nil, fmt.Errorf("can not get `html` key in data")
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
		return nil, fmt.Errorf("can not unmarsahl data %s", err.Error())
	}

	urlString, ok := mapData["url"].(string)
	if !ok {
		return nil, fmt.Errorf("can not get `url` key in data")
	}

	switch getSiteTypeFromURL(urlString) {
	case YOUTUBE_SITE_TYPE:

		url, err := url.ParseRequestURI(urlString)
		if err != nil {
			return nil, fmt.Errorf("can not parse `url` string to URL struct")
		}

		videoID := url.Query().Get("v")
		if len(videoID) == 0 {
			// skip this case
			log.Println("videoId not exist in the url")
			return data, nil
		}

		mapData["youtube_video_id"] = videoID
		return json.Marshal(mapData)

	default:
		return data, nil
	}
}
