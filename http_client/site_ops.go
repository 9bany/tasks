package httpclient

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

func siteOperateData(urlString string, data []byte) ([]byte, error) {
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
