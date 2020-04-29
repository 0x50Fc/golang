package tn

import (
	"errors"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/hailongz/golang/tunnel/kk"
)

type RService struct {
	id   int64
	lock sync.Mutex
	req  map[int64]chan *kk.Response
}

func init() {
	AddDefaultService("req", NewRService())
}

func NewRService() *RService {
	return &RService{req: map[int64]chan *kk.Response{}}
}

func (S *RService) Send(sender Sender, to string, uri string, ctype string, data []byte, header map[string]string, timeout time.Duration) (*kk.Response, error) {

	var c chan *kk.Response = nil

	S.lock.Lock()
	id := S.id + 1
	S.id = id
	m := kk.Message{Type: "req", To: to}
	req := kk.Request{Uri: uri, Data: data, Header: header, Id: id, Type: ctype}
	// fmt.Println(req)
	m.Data, _ = proto.Marshal(&req)
	err := sender.Send(&m)
	if err == nil {
		c = make(chan *kk.Response)
		defer func(id int64, S *RService) {
			S.lock.Lock()
			delete(S.req, id)
			S.lock.Unlock()
		}(id, S)
		defer close(c)
		S.req[id] = c
	}
	S.lock.Unlock()

	if err != nil {
		return nil, err
	}

	if timeout > 0 {
		timer := time.NewTimer(timeout)
		defer timer.Stop()
		select {
		case <-timer.C:
			return nil, errors.New("request timeout")
		case resp := <-c:
			// fmt.Println(resp)
			return resp, nil
		}
	} else {
		resp := <-c
		return resp, nil
	}
}

func (S *RService) Match(message *kk.Message) bool {
	return message.Type == "resp"
}

func (S *RService) Handle(message *kk.Message, sender Sender) error {
	if message.Type == "resp" {
		resp := kk.Response{}
		err := proto.Unmarshal(message.Data, &resp)
		if err != nil {
			return err
		}
		// fmt.Println(resp)
		S.lock.Lock()
		c := S.req[resp.Id]
		S.lock.Unlock()
		if c != nil {
			c <- &resp
		}
	}
	return nil
}
