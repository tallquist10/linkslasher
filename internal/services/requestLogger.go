package services

import (
	"fmt"
	"time"

	"github.com/tallquist10/linkslasher/internal/links"
)

type RequestLogger struct {
	requestChannel chan *links.LinksApiRequest
}

func NewRequestLogger(rc chan *links.LinksApiRequest) *RequestLogger {
	return &RequestLogger{
		requestChannel: rc,
	}
}

func (rl *RequestLogger) LogRequest(req *links.LinksApiRequest) {
	rl.requestChannel <- req
}

func (rl *RequestLogger) Listen() {
	for {
		select {
		case req := <-rl.requestChannel:
			fmt.Println(req)
		default:
			time.Sleep(time.Second)
		}
	}
}
