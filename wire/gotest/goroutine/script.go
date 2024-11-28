package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

const maxGoroutine int = 10

func hitApi(url string, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Hit %d: Error occured: %s\n", id, err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Hit %d: Success, response length: %d\n", id, len(body))
	} else {
		fmt.Printf("Hit %d: Failed with status %d\n", id, resp.StatusCode)
	}
}

func main() {
	fmt.Println("----- test go routine -----")

	url := "https://visitcount.itsvg.in/api?id=gpjen&icon=0&color=0"
	totalHits := 100

	sem := make(chan struct{}, maxGoroutine)

	var wg sync.WaitGroup

	for i := 1; i <= totalHits; i++ {
		wg.Add(1)

		sem <- struct{}{}

		go func(id int) {
			defer func() { <-sem }()
			hitApi(url, id, &wg)
		}(i)
	}

	wg.Wait()

	fmt.Println("All Request Completed...")

}
