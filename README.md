# Gopattern

## 알고 넘어가야할 것들
### The Go Memory Model
<details>
<summary>:scroll: 원문 보기</summary>
The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed 
to observe values produced by writes to the same variable in a different goroutine.

iPrograms that modify data being simultaneously accessed by multiple goroutines must serialize such access.
To serialize access, protect the data with channel operations or other synchronization primitives such as those in the sync and sync/atomic packages.
If you must read the rest of this document to understand the behavior of your program, you are being too clever.

Don't be clever.
</details>

고의 메모리 모델은 하나의 변수가 다른 고루틴에 의해 변경되고 다른 고루틴이 이 변수에 접근할 때 변경되는 값들을 읽는 것을 보장하는 것에 주안점을 두고 있습니다.
변수에 대한 쓰기 작업을 하는 고루틴은 고의 동기화 로직을 통해 직렬화 되야합니다. 동기화에 필요한 모듈은 `sync` 혹은 `sync/atomic` 패키지 혹은 `gochannel` 를 통해 지원됩니다.

#### Happens Before
<details>
<summary>:scroll: 원문 보기</summary>
Within a single goroutine, reads and writes must behave as if they executed in the order specified by the program. That is, compilers and processors may reorder the reads and writes executed within a single goroutine only when the reordering does not change the behavior within that goroutine as defined by the language specification. Because of this reordering, the execution order observed by one goroutine may differ from the order perceived by another. For example, if one goroutine executes a = 1; b = 2;, another might observe the updated value of b before the updated value of a.

To specify the requirements of reads and writes, we define happens before, a partial order on the execution of memory operations in a Go program. If event e1 happens before event e2, then we say that e2 happens after e1. Also, if e1 does not happen before e2 and does not happen after e2, then we say that e1 and e2 happen concurrently.
</details>

읽기 및 쓰기 작업은 프로그램에 의해 지정된 순서에 맞게 실행됩니다. 즉 단일 고루틴이 정의된 작업을 변경하지 않는 경우에만 컴파일러와 프로세서는 작업 순서를 재정렬 할 수 있습니다.
재정렬을 통해 다른 고루틴은 변경사항을 감시하는 것이 가능합니다. 예로 하나의 고루틴이 `a = 1; b = 2` 작업을 수행하면 다른 고루틴은 `b = 2`를 먼저 관찰하고 이전에 발생한 `a = 1` 작업도 관찰합니다.

읽기 및 쓰기 작업에 대한 요구사항을 우리는 <b>Happens before</b>라고 정의합니다. 메모리 실행 순서를 예로 든다면  e<sub>1</sub> 이벤트는 e<sub>2</sub>전에 발생하면 우리는 일련의 작업이 순차적으로 발생함을 알수 있습니다. 이와 달리 e<sub>1</sub>이 e<sub>2</sub> 이전에도 발생하지 않고 이후에도 발생하지 않는다면 이는 작업이 동시에 이루어짐을 뜻합니다.

#### Synchronization

```go
var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}
```
위 코드는 즉각적으로 hello, world를 출력하지 않지만, f 고루틴은 a 변수에 쓰기작업이 있다는 것을 감시할 수 있음으로
언젠가는 hello, world를 출력하게됩니다.

```go
var a string

func hello() {
	go func() { a = "hello" }()
	print(a)
}
```
위 코드의 경우는 hello 함수는 a에 쓰기 작업이 있었는지 감시할 수가 없습니다. 따라서 별도의 동기화 구문을 추가하지 않는다면
컴파일러에 의해 go문 자체가 없어질 수 있습니다.

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
다른 언어에서 사용되는 제너레이터 혹은 이터레이터와 비슷합니다.

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
javascript에서 흔히 볼수 있는 future 기반의 비동기 로직입니다. 주로 http 요청과 같은 
응답 반응 로직에 사용됩니다. 비동기적으로 요청작업을 수행하면서 그 사이 다른 로직을 수행하고 
응답을 받는 시점에 다시 로직을 수행합니다.

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
		defer wg.Done()
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
waitGroup은 일종의 세마포어 처럼 작동합니다. 기존 세마포어에 익숙한 개발자들은
channel을 이용해 별도로 세마포어를 만들어 사용하기도 합니다.

## 기타
### Once
```go
func onceDo(n int) {
	var justOne int
	var once sync.Once
	wg := &sync.WaitGroup{}

	increaseCount := func() {
		justOne++
		print("set one")
	}

	doOne := func() {
		once.Do(increaseCount)
		println(justOne)
		wg.Done()
	}

	wg.Add(n + 1)
	for i := 0; i <= n; i++ {
		go doOne()
	}
	wg.Wait()
}
```
`sync.Once`는 고루틴에서 한번만 호출됨을 보장하는 메커니즘입니다.
위 코드에 `increaseCount`는 한번만 호출됨으로 justOne의 값은 1로 유지됩니다.
주로 고루틴 내에서의 변수를 초기화하기 위한 방법으로 사용됩니다.

### Select
```go
func selectCase() {
	nState := numberGenerator(100)
	lock := &sync.Mutex{}

	for {
		select {
		case num := <-nState:
			lock.Lock()
			print(num)
		case num := <-nState:
			fmt.Printf("%d print line 2", num)
		}
	}
}

```
select문은 언뜻 비동기적으로 이루어진 구간처럼 보이지만 동기화 구간임으로 하나의 case에 락이 걸리면 다른 case 또한 락이 걸립니다.

### Buffered, Unbuffered channel
```go
func bufferedChannel() <-chan int {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	return c
}
```
buffured channel은 버퍼가 모두 찬 후 입력시 락이 걸리고 출력시 버퍼가 비면 락이 걸립니다.
보통 channel에 쓰기작업을 하는 고루틴을 반환하여 메모리 최적화를 이룰 수 있는 부분에 사용됩니다.

```go
func unbufferedChannel() <-chan int {
	c := make(chan int)
	c <- 1
	return c
}
```
unbuffured channel은 입력시 출력을 해주는 부분이 없으면 락이 걸립니다. 보통 입력 부분에서 대기할 필요가 있을 때 사용됩니다.

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
