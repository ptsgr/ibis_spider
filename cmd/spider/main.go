package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
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

func urlGenerator(urlsCh chan string) {
	for _, url := range urls {
		urlsCh <- url
		time.Sleep(time.Second * 5)
	}
	close(urlsCh)
}

func main() {
	urlsCh := make(chan string)
	go urlGenerator(urlsCh)

	var wg sync.WaitGroup
	for url := range urlsCh {
		wg.Add(1)
		go getHttpStatusCode(url, &wg)
	}

	wg.Wait()
}

func getHttpStatusCode(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil || !(isStatusCodeSuccess(resp.StatusCode)) {
		fmt.Println(url, " - not OK")
		return
	}
	fmt.Println(url, " - OK")

}

func isStatusCodeSuccess(code int) bool {
	return code >= 200 && code <= 299
}
