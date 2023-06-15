package main

import (
	"context"
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
		"https://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}
)

func urlGenerator(urlsCh chan string) {
	for _, url := range urls {
		time.Sleep(time.Second * 1)
		urlsCh <- url
	}
	close(urlsCh)
}

func main() {
	urlsCh := make(chan string)
	go urlGenerator(urlsCh)

	successCouter := newCounter()
	var wg sync.WaitGroup
	wait := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-wait:
				return
			default:
				if successCouter.Get() > 1 {
					cancel()
					<-wait
					return
				}
			}
		}
	}()

	for url := range urlsCh {
		wg.Add(1)
		go getHttpStatusCode(ctx, url, successCouter, &wg)
	}
	wait <- struct{}{}
	wg.Wait()

}

func getHttpStatusCode(ctx context.Context, url string, successCouter *counter, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println(url, " - terminated normally")
	default:
		resp, err := http.Get(url)
		if err != nil || !(isStatusCodeSuccess(resp.StatusCode)) {
			successCouter.Reset()
			fmt.Println(url, " - not OK")
			return
		}
		successCouter.Add()
		fmt.Println(url, " - OK")
	}
}

func isStatusCodeSuccess(code int) bool {
	return code >= 200 && code <= 299
}
