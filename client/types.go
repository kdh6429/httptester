package client

type Resp struct {
	Err     error
	Status  int
	Latency int64
	Size    int
}

type Client interface {
	Request(chan<- *Resp)
}
