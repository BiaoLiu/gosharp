package services

import (
	"net/url"
)

func ParseUrl(rawUrl string) string {
	if u, err := url.Parse(rawUrl); err != nil {
		return u.Path
	}
	return ""
}
