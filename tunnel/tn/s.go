package tn

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/hailongz/golang/tunnel/kk"
)

type Done struct {
}

func (D *Done) Error() string {
	return "done"
}

var defaultService = map[string]Service{}

func AddDefaultService(name string, s Service) {
	defaultService[name] = s
}

func GetDefaultService(name string) Service {
	return defaultService[name]
}

type SNode struct {
	s []Service
}

func (S *SNode) AddService(s Service) {
	if S.s == nil {
		S.s = []Service{}
	}
	S.s = append(S.s, s)
}

func (S *SNode) Handle(message *kk.Message, sender Sender) error {
	if S.s != nil {
		var err error = nil
		for _, s := range S.s {
			if s.Match(message) {
				err = s.Handle(message, sender)
				if err != nil {
					_, ok := err.(*Break)
					if ok {
						break
					}
					return err
				}
			}
		}
	}
	return nil
}

type NodeSet struct {
	s    map[int64]Node
	lock sync.Mutex
}

func (N *NodeSet) AddNode(id int64, n Node) {
	N.lock.Lock()
	defer N.lock.Unlock()
	if N.s == nil {
		N.s = map[int64]Node{}
	}
	N.s[id] = n
}

func (N *NodeSet) RemoveNode(id int64) {
	N.lock.Lock()
	defer N.lock.Unlock()
	if N.s != nil {
		delete(N.s, id)
	}
}

func (N *NodeSet) Send(message *kk.Message) error {
	var err error = nil
	nodes := []Node{}

	N.lock.Lock()

	if N.s != nil {
		for _, s := range N.s {
			if s.Match(message.To) {
				nodes = append(nodes, s)
			}
		}
	}

	defer N.lock.Unlock()

	if strings.HasSuffix(message.To, "*") {

		for _, s := range nodes {
			s.Send(message)
		}

	} else {

		if len(nodes) == 0 {
			return &Done{}
		}

		sort.Slice(nodes, func(i, j int) bool {
			a := nodes[i]
			b := nodes[j]
			v := a.Priority() - b.Priority()
			if v == 0 {
				v = rand.Intn(3) - 1
			}
			return v > 0
		})

		for _, s := range nodes {
			err = s.Send(message)
			if err == nil {
				break
			}
		}
	}

	return nil
}

type Transport struct {
	Id        int64
	c         Channel
	s         Handler
	w         chan *kk.Message
	loopbreak bool
}

func NewTransport(id int64, c Channel, s Handler, e chan int64) *Transport {

	w := make(chan *kk.Message, 2048)

	v := Transport{Id: id, c: c, s: s, loopbreak: false, w: w}

	go func() {

		for {

			m, err := c.Read()

			if err != nil {
				fmt.Println(err)
				v.loopbreak = true
				c.Close()
				break
			}

			if m == nil {
				break
			}

			// fmt.Println(m)

			s.Handle(m, &v)

			if v.loopbreak {
				break
			}
		}

		w <- nil
		e <- id

	}()

	go func() {

		for {

			m := <-w

			if m == nil {
				break
			}

			err := c.Write(m)

			if err != nil {
				v.loopbreak = true
				c.Close()
				break
			}

			if v.loopbreak {
				break
			}
		}

		close(w)

	}()

	return &v
}

func (T *Transport) Send(message *kk.Message) error {

	if T.loopbreak {
		return errors.New("transport closed")
	}

	T.w <- message

	return nil
}

func (T *Transport) Addr() string {
	return T.c.Addr()
}

func (T *Transport) Close() error {
	T.loopbreak = true
	return T.c.Close()
}

type CNode struct {
	name     string
	sender   Sender
	s        Handler
	priority int
}

func NewCNode(name string, sender Sender, s Handler) *CNode {
	return &CNode{name: name, sender: sender, s: s}
}

func (N *CNode) Name() string {
	return N.name
}

func (N *CNode) Match(name string) bool {

	if N.name == name {
		return true
	}

	if strings.HasPrefix(name, N.name) {
		return true
	}

	return false
}

func (N *CNode) Send(message *kk.Message) error {
	if message.From == "" {
		message.From = N.name
	}
	return N.sender.Send(message)
}

func (N *CNode) Handle(message *kk.Message) error {
	return N.s.Handle(message, N)
}

func (N *CNode) Priority() int {
	return N.priority
}

func (N *CNode) Addr() string {
	return N.sender.Addr()
}

func (N *CNode) Close() error {
	return N.sender.Close()
}

type MNode struct {
	SNode
	name        string
	id          int64
	childrenSet NodeSet
	parentSet   NodeSet
}

func (N *MNode) NewId() int64 {
	id := N.id + 1
	N.id = id
	return id
}

func (N *MNode) mnodeInit(name string) {
	N.name = name
	N.childrenSet = NodeSet{}
	N.parentSet = NodeSet{}
}

func (N *MNode) AddChidren(id int64, n Node) int64 {
	N.childrenSet.AddNode(id, n)
	return id
}

func (N *MNode) RemoveChidren(id int64) {
	N.childrenSet.RemoveNode(id)
}

func (N *MNode) AddParent(id int64, n Node) {
	N.parentSet.AddNode(id, n)
}

func (N *MNode) RemoveParent(id int64) {
	N.parentSet.RemoveNode(id)
}

func (N *MNode) Name() string {
	return N.name
}

func (N *MNode) Match(name string) bool {

	if N.name == name {
		return true
	}

	if strings.HasPrefix(name, N.name+".") {
		return true
	}

	return false
}

func (N *MNode) Send(message *kk.Message, sender Sender) error {

	if message.From == "" {
		message.From = N.name
	}

	var err error = nil

	if N.Match(message.To) {
		err = N.childrenSet.Send(message)
	} else {
		err = N.parentSet.Send(message)
	}

	if err != nil {

		_, done := err.(*Done)

		if done {

			if message.Type == "req" {

				req := kk.Request{}
				err = proto.Unmarshal(message.Data, &req)

				if err != nil {
					return err
				}

				resp := kk.Response{Id: req.Id, Code: 404, Type: "text", Data: []byte("Not Found")}
				m := kk.Message{Type: "resp", To: message.From}
				m.Data, _ = proto.Marshal(&resp)

				return sender.Send(&m)

			}

			return nil
		}
		return err
	}

	return nil
}
