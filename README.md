# Gopattern

## 일반적으로 사용되는 비동기 패턴

### generator
```go
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

func main() {
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
```

### future
```go
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

func main() {
	data := futureData(testServer.URL)
	expected := bodyData{Body: []byte("Test Server"), Error: nil}
	assert.Equal(t, expected, <-data)
}
```

### waitGroup
```go
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
		go printNum(waitGroup, v)

		if v%9 == 0 && v != 0 {
			waitGroup.Wait()
			println()
		}
	}
}
```



## 유용한 도구들
* [gore](https://github.com/motemen/gore): REPL 도구
* [gin](https://github.com/gin-gonic/gin): 경량 웹프레임워크
* [revel](https://revel.github.io/): 웹프레임워크
* [beego](https://beego.me/): 웹프레임워크
* [gorilla](http://www.gorillatoolkit.org/): 웹툴킷
* [sarama](https://github.com/Shopify/sarama): 카프카 클라이언트 
* [goad](https://github.com/goadapp/goad): AWS 람다 기반 부하분산 테스트 도구
* [logrus](https://github.com/sirupsen/logrus): 로그 라이브러리
* [libchan](https://github.com/docker/libchan): gochannel 기반 서버간 통신 라이브러리
* [etcd](https://github.com/coreos/etcd): 분산 시스템 기반 키밸류 저장소
* [squirrel](https://github.com/Masterminds/squirrel): SQL 포맷터
* [sqlx](https://github.com/jmoiron/sqlx): 기본 sql 라이브러리 확장버전
* [viper](https://github.com/spf13/viper): project configuration

## Go 기반 주목 받는 프로젝트들
* [cockroachdb](https://github.com/cockroachdb/cockroach): 분산 SQL 데이터베이스
* [NSQ](https://github.com/nsqio/nsq): 실시간 분산 메세지 큐 시스템









