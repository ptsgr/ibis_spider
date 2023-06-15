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

type checkInfo struct {
	url        string
	statusCode int
}

func main() {
	wait := make(chan struct{}, len(urls))      // chennal for syncronization
	resultCh := make(chan checkInfo, len(urls)) //
	for _, url := range urls {
		go getHttpStatusCode(url, wait, resultCh)
	}
	for i := 0; i < len(urls); i++ {
		<-wait
	}
	close(wait)
	close(resultCh)

	for info := range resultCh {
		if isStatusCodeSuccess(info.statusCode) {
			fmt.Println(info.url, " - OK")
		} else {
			fmt.Println(info.url, " - not OK")
		}
	}
}

func getHttpStatusCode(url string, wait chan struct{}, ch chan checkInfo) {
	defer func() {
		wait <- struct{}{}
	}()
	resp, err := http.Get(url)
	if err != nil {
		ch <- checkInfo{url, 500}
		return
	}
	ch <- checkInfo{url, resp.StatusCode}
}

func isStatusCodeSuccess(code int) bool {
	return code >= 200 && code <= 299
}
