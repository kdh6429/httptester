package client

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type fastHttp struct {
	uri     string
	timeout time.Duration
	client  *fasthttp.Client
}

func (h *fastHttp) Request(rchan chan<- *Resp) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(h.uri)
	req.Header.SetMethod(fasthttp.MethodGet)
	res := fasthttp.AcquireResponse()

	startTime := time.Now()
	if err := h.client.Do(req, res); err != nil {
		rchan <- &Resp{
			Err:     err,
			Status:  -1,
			Latency: -1,
			Size:    -1,
		}
		return
	}
	fasthttp.ReleaseRequest(req)

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
func FastHTTP(uri string) Client {
	return &fastHttp{
		uri:     uri,
		timeout: time.Second * 5, // default timeout
		client: &fasthttp.Client{
			// ReadTimeout:                   readTimeout,
			// WriteTimeout:                  writeTimeout,
			// MaxIdleConnDuration:           maxIdleConnDuration,
			NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
			DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
			DisablePathNormalizing:        true,
			// increase DNS cache time to an hour instead of default minute
			Dial: (&fasthttp.TCPDialer{
				Concurrency:      4096,
				DNSCacheDuration: time.Hour,
			}).Dial,
		},
	}
}
