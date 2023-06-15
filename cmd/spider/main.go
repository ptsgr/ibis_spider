package main

import (
	"fmt"
	"net/http"
)

var (
	urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}
)

func main() {
	for _, url := range urls {
		if isStatusCodeSuccess(getHttpStatusCode(url)) {
			fmt.Println(url, " - OK")

		} else {
			fmt.Println(url, " - not OK")
		}
	}
}

func getHttpStatusCode(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		return 500
	}
	return resp.StatusCode
}

func isStatusCodeSuccess(code int) bool {
	return code >= 200 && code <= 299
}
