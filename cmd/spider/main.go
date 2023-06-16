package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ptsgr/ibis_spider/internal/storage"
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
		time.Sleep(time.Second * 1)
	}
	close(urlsCh)
}

func main() {
	urlsCh := make(chan string)
	go urlGenerator(urlsCh)

	dsn := os.Getenv("STORAGE_DSN")
	if len(dsn) == 0 {
		log.Fatalf("Storage DSN not set")
	}

	s, err := storage.NewGormStorage(dsn)
	if err != nil {
		log.Fatalf("Cannot init Storage DB, err: %v", err)
		return
	}

	runID, err := s.CreateRun()
	if err != nil {
		log.Fatalf("Cannot create Storage Run entity, err: %v", err)
		return
	}

	var wg sync.WaitGroup
	for url := range urlsCh {
		wg.Add(1)
		go getHttpStatusCode(url, s, runID, &wg)
	}

	wg.Wait()
}

func getHttpStatusCode(url string, s storage.Storage, runID int, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil || !(isStatusCodeSuccess(resp.StatusCode)) {
		err = s.SetURLStatus(runID, url, storage.StatusNOK)
		if err != nil {
			log.Fatalf("Cannot set URL status in Storage, err: %v", err)
		}
		fmt.Println(url, " - not OK")
		return
	}
	err = s.SetURLStatus(runID, url, storage.StatusOK)
	if err != nil {
		log.Fatalf("Cannot set URL status in Storage, err: %v", err)
	}
	fmt.Println(url, " - OK")

}

func isStatusCodeSuccess(code int) bool {
	return code >= 200 && code <= 299
}
