package api

import "net/url"

func urlValid(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)
	return err == nil
}
