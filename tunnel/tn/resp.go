package tn

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/hailongz/golang/tunnel/kk"
)

type RespHandler struct {
	prefix string
	handle func(req *kk.Request) (*kk.Response, error)
}

type RespService struct {
	s []*RespHandler
}

func init() {
	AddDefaultService("resp", NewRespService())
}

func NewRespService() *RespService {
	return &RespService{s: []*RespHandler{}}
}

func (S *RespService) AddHandler(prefix string, handle func(req *kk.Request) (*kk.Response, error)) {
	S.s = append(S.s, &RespHandler{prefix: prefix, handle: handle})
}

func (S *RespService) Match(message *kk.Message) bool {
	return message.Type == "req"
}

func (S *RespService) Handle(message *kk.Message, sender Sender) error {
	if message.Type == "req" {
		req := kk.Request{}
		err := proto.Unmarshal(message.Data, &req)
		fmt.Println(req)
		if err != nil {
			return err
		}
		for _, s := range S.s {
			if strings.HasPrefix(req.Uri, s.prefix) {
				go func(req *kk.Request, s *RespHandler, to string, sender Sender) {
					resp, err := s.handle(req)
					if err != nil {
						m := kk.Message{Type: "resp", To: to}
						resp = &kk.Response{Id: req.Id, Code: 500, Type: "text", Data: []byte(err.Error())}
						m.Data, _ = proto.Marshal(resp)
						sender.Send(&m)
					} else {
						m := kk.Message{Type: "resp", To: to}
						resp.Id = req.Id
						m.Data, _ = proto.Marshal(resp)
						fmt.Println(resp)
						sender.Send(&m)
					}
				}(&req, s, message.From, sender)
				return nil
			}
		}

		resp := kk.Response{}
		m := kk.Message{Type: "resp", To: message.From}
		resp.Id = req.Id
		resp.Code = 404
		resp.Type = "text"
		resp.Data = []byte("Not Found")
		m.Data, _ = proto.Marshal(&resp)
		sender.Send(&m)
	}
	return nil
}
