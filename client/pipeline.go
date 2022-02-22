package client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type pipelinefastHttp struct {
	uri     string
	timeout time.Duration
	client  *fasthttp.PipelineClient
}

func (h *pipelinefastHttp) Request(rchan chan<- *Resp) {
	req := fasthttp.AcquireRequest()
	req.SetBody([]byte("hello, world!"))
	req.SetRequestURI(h.uri)
	res := fasthttp.AcquireResponse()

	startTime := time.Now()
	if err := h.client.DoTimeout(req, res, h.timeout); err != nil {
		rchan <- &Resp{
			Err:     err,
			Status:  -1,
			Latency: -1,
			Size:    -1,
		}
		return
	}
	size := len(res.Body()) + 2
	res.Header.VisitAll(func(key, value []byte) {
		size += len(key) + len(value) + 2
	})
	statusCode := res.Header.StatusCode()
	fmt.Print(statusCode)
	res.Reset()

	rchan <- &Resp{
		Err:     nil,
		Status:  statusCode,
		Latency: time.Now().Sub(startTime).Milliseconds(),
		Size:    size,
	}
}
func PipelineHttpClient(uri string) Client {
	u, _ := url.Parse(uri)
	hostname := u.Hostname()
	port := u.Port()
	if port == "" {
		port = "80"
	}

	return &pipelinefastHttp{
		uri:     uri,
		timeout: time.Second * 5, // default timeout
		client: &fasthttp.PipelineClient{
			Addr:  fmt.Sprintf("%v:%v", hostname, port),
			IsTLS: u.Scheme == "https",
		},
	}

}
