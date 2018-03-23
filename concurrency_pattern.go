package gopattern

import (
	"io/ioutil"
	"net/http"
	"sync"
)

func numberGenerator(n int) <-chan int {
	nState := make(chan int)
	go func() {
		for i := 1; i <= n; i = i + 1 {
			nState <- i
		}
		defer close(nState)
	}()
	return nState
}

type bodyData struct {
	Body  []byte
	Error error
}

func futureData(url string) <-chan bodyData {
	pipe := make(chan bodyData, 1)

	go func() {
		var body []byte
		var err error

		resp, err := http.Get(url)
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)
		pipe <- bodyData{Body: body, Error: err}
	}()

	return pipe
}

func wgSchedule() {
	var numArray []int

	for i := 0; i <= 100; i = i + 1 {
		numArray = append(numArray, i%10)
	}

	waitGroup := &sync.WaitGroup{}
	printNum := func(wg *sync.WaitGroup, n int) {
		print(n)
		wg.Done()
	}

	for _, v := range numArray {
		waitGroup.Add(1)
		printNum(waitGroup, v)

		if v%9 == 0 && v != 0 {
			println()
			waitGroup.Wait()
		}
	}
}
