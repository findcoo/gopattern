package gopattern

import (
	"io/ioutil"
	"net/http"
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
