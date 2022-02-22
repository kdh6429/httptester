package httptester

import (
	"context"
	"time"

	"github.com/kdh6429/httptester/client"
)

const (
	clientDefault    = 5
	FactorDefault    = 10
	DelaytimeDefault = time.Second * 1
)

type HttpTester interface {
	Do()
}
type HttpTestBuilder interface {
	Uri(string) HttpTestBuilder
	Clients(int) HttpTestBuilder
	Factors(int) HttpTestBuilder
	DelayTime(time.Duration) HttpTestBuilder
}
type httpTest struct {
	client    client.Client
	uri       string
	nClient   int
	nFactor   int
	delaytime time.Duration
	timeout   time.Duration
	rChan     chan *client.Resp
}

func (t *httpTest) Uri(uri string) *httpTest {
	t.uri = uri
	return t
}
func (t *httpTest) NClient(nClient int) *httpTest {
	t.nClient = nClient
	return t
}
func (t *httpTest) NFactor(NFactor int) *httpTest {
	t.nFactor = NFactor
	return t
}
func (t *httpTest) Timeout(timeout time.Duration) *httpTest {
	t.timeout = timeout
	return t
}
func (t *httpTest) DelayTime(delaytime time.Duration) *httpTest {
	t.delaytime = delaytime
	return t
}
func (t *httpTest) Build() (*httpTest, <-chan *client.Resp) {
	t.rChan = make(chan *client.Resp, 2*t.nClient*t.nFactor)
	return t, t.rChan
}
func (t *httpTest) Do() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(t.rChan)
				return
			default:
				t.request()
			}
		}
	}()
}
func (t *httpTest) request() {
	for i := 0; i < t.nClient; i++ {
		newClient := t.client
		for j := 0; j < t.nFactor; j++ {
			newClient.Request(t.rChan)
		}
	}
}
func NewHttpTest(client client.Client) *httpTest {
	return &httpTest{
		client:    client,
		nClient:   clientDefault,
		nFactor:   FactorDefault,
		delaytime: DelaytimeDefault,
	}
}

// prebuild testers
func NewFastHttpTest(uri string) *httpTest {
	return &httpTest{
		client:    client.FastHTTP(uri),
		nClient:   clientDefault,
		nFactor:   FactorDefault,
		delaytime: DelaytimeDefault,
	}
}
func NewPipelineHttpTest(uri string) *httpTest {
	return &httpTest{
		client:    client.PipelineHttpClient(uri),
		nClient:   clientDefault,
		nFactor:   FactorDefault,
		delaytime: DelaytimeDefault,
	}
}

// mock http tester for test
func NewFakeHttpTest() *httpTest {
	return &httpTest{
		client:    client.FakeHttp(),
		nClient:   clientDefault,
		nFactor:   FactorDefault,
		delaytime: DelaytimeDefault,
	}
}
