package client

import (
	"math/rand"
	"time"
)

type fakeHttp struct {
}

func (h *fakeHttp) Request(rchan chan<- *Resp) {
	startTime := time.Now()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	rchan <- &Resp{
		Err:     nil,
		Status:  200,
		Latency: time.Now().Sub(startTime).Milliseconds(),
		Size:    0,
	}
}
func FakeHttp() Client {
	return &fakeHttp{}
}
