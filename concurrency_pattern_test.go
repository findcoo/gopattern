package gopattern

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"io"

	"github.com/stretchr/testify/assert"
)

var testServer *httptest.Server

func setUpTestServer() {
	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Test Server")
	}))
}

func TestNumberGenerator(t *testing.T) {
	state := numberGenerator(100)

	for {
		select {
		case recieved := <-state:
			if recieved >= 100 {
				t.Log("NumberGenerator test success")
				return
			}
		}
	}
}

func TestFutureData(t *testing.T) {
	setUpTestServer()
	data := futureData(testServer.URL)
	expected := bodyData{Body: []byte("Test Server"), Error: nil}
	assert.Equal(t, expected, <-data)
}

func TestWGScheduleTest(t *testing.T) {
	wgSchedule()
}

func TestOnceDo(t *testing.T) {
	onceDo(50)
}

func TestSelectCase(t *testing.T) {
	selectCase()
}

func TestBufferedChannel(t *testing.T) {
	c := bufferedChannel()
	for {
		select {
		case num := <-c:
			print(num)
		}
	}
}

func TestUnbufferedChannel(t *testing.T) {
	c := unbufferedChannel()
	for {
		select {
		case num := <-c:
			print(num)
		}
	}
}
