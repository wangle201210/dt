package url

import "net/url"

func Encode(query string) string {
	return url.QueryEscape(query)
}

func Decode(str string) (query string, err error) {
	return url.QueryUnescape(str)
}
