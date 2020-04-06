package tn

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/hailongz/golang/tunnel/kk"
)

type TCPChannel struct {
	conn net.Conn
	data []byte
	len  int
}

func NewTCPChannel(conn net.Conn) *TCPChannel {
	return &TCPChannel{conn: conn}
}

func (C *TCPChannel) Grow(size int) {
	if C.data == nil {
		C.data = make([]byte, C.len+size)
	} else if len(C.data) < size+C.len {
		b := make([]byte, size+C.len)
		if C.len > 0 {
			copy(b, C.data[0:C.len])
		}
		C.data = b
	}
}

func (C *TCPChannel) Read() (*kk.Message, error) {

	var n int = 0
	var i64 int64 = 0
	var err error = nil
	if C.len < 4 {
		C.Grow(2048)
		n, err = C.conn.Read(C.data)
		if err != nil {
			return nil, err
		}
		C.len = C.len + n
	}

	if C.len < 4 {
		return nil, nil
	}

	i64, n = binary.Varint(C.data[0:4])

	if n <= 0 || i64 <= 0 {
		return nil, errors.New("stream error")
	}

	size := int(i64) + 4

	C.Grow(size - 4)

	if C.len < size {
		n, err = C.conn.Read(C.data[4:])
		if err != nil {
			return nil, err
		}
		C.len = C.len + n
	}

	if C.len < size {
		return nil, nil
	}

	m := kk.Message{}

	err = proto.Unmarshal(C.data[4:size], &m)

	if err != nil {
		return nil, err
	}

	C.len = C.len - size

	if C.len > 0 {
		copy(C.data, C.data[size:size+C.len])
	}

	return &m, nil
}

func (C *TCPChannel) Write(message *kk.Message) error {

	data, err := proto.Marshal(message)

	if err != nil {
		return err
	}

	size := make([]byte, 4)

	binary.PutVarint(size, int64(len(data)))

	for len(size) > 0 {

		n, err := C.conn.Write(size)

		if err != nil {
			return err
		}

		if n >= len(size) {
			break
		}

		size = size[n:]
	}

	for len(data) > 0 {

		n, err := C.conn.Write(data)

		if err != nil {
			return err
		}

		if n >= len(data) {
			break
		}

		data = data[n:]
	}

	return nil
}

func (C *TCPChannel) Close() error {
	return C.conn.Close()
}

func (C *TCPChannel) Addr() string {
	return C.conn.RemoteAddr().String()
}

type TCPCNode struct {
	SNode
	name      string
	addr      string
	e         chan int64
	transport *Transport
	lock      sync.Mutex
	priority  int
	ping      time.Duration
	atime     time.Duration
}

func NewTCPCNode(name string, addr string) *TCPCNode {

	e := make(chan int64, 32)
	v := TCPCNode{}

	v.name = name
	v.addr = addr
	v.e = e
	v.priority = 0

	go func() {

		var b int64
		var ping = time.NewTimer(time.Second * 6)
		var loopbreak = false
		var m = kk.Message{Type: "ping", Data: make([]byte, binary.MaxVarintLen64)}
		defer close(e)
		defer ping.Stop()

		for !loopbreak {

			select {
			case b = <-e:
				if b == 0 {
					loopbreak = true
					break
				} else if b == 1 {
					v.lock.Lock()
					v.transport = nil
					v.lock.Unlock()
				}
				break
			case <-ping.C:
				if !strings.HasSuffix(v.name, "*") {
					m.From = v.name
					binary.PutVarint(m.Data, time.Now().UnixNano())
					v.Send(&m)
				}
				break
			}

		}

	}()

	return &v
}

func (N *TCPCNode) Name() string {
	return N.name
}

func (N *TCPCNode) Ping() time.Duration {
	return N.ping
}

func (N *TCPCNode) Atime() time.Duration {
	return N.atime
}

func (N *TCPCNode) Match(name string) bool {

	if N.name == name {
		return true
	}

	if strings.HasPrefix(name, N.name) {
		return true
	}

	return false
}

func (N *TCPCNode) Send(message *kk.Message) error {

	N.lock.Lock()
	defer N.lock.Unlock()

	if N.transport == nil {
		conn, err := net.Dial("tcp", N.addr)
		if err != nil {
			return err
		}
		channel := NewTCPChannel(conn)
		N.transport = NewTransport(1, channel, N, N.e)
		N.ping = 0
	}

	if message.From == "" {
		message.From = N.name
	}

	return N.transport.Send(message)
}

func (N *TCPCNode) Handle(message *kk.Message, sender Sender) error {
	if message.Type == "login.done" {
		N.name = message.To
		fmt.Printf("[LOGIN] %s\n", message.To)
		return nil
	}
	if message.Type == "pong" {
		t, n := binary.Varint(message.Data)
		if n > 0 {
			N.atime = time.Duration(time.Now().UnixNano())
			N.ping = N.atime - time.Duration(t)
		}
		return nil
	}
	return N.SNode.Handle(message, N)
}

func (N *TCPCNode) Priority() int {
	return N.priority
}

func (N *TCPCNode) Addr() string {
	return N.addr
}

func (N *TCPCNode) Close() error {

	var err error = nil
	N.lock.Lock()
	defer N.lock.Unlock()

	if N.transport != nil {
		err = N.transport.Close()
		N.transport = nil
	}

	N.e <- 0

	return err
}

type TCPMNode struct {
	MNode
	fd net.Listener
}

func NewTCPMNode(name string, addr string) (*TCPMNode, error) {

	fd, err := net.Listen("tcp", addr)

	if err != nil {
		return nil, err
	}

	v := TCPMNode{}
	v.mnodeInit(name)
	v.fd = fd

	return &v, nil
}

func (N *TCPMNode) Handle(message *kk.Message, sender Sender) error {
	if message.Type == "login" {
		transport := sender.(*Transport)
		if transport != nil {
			if strings.HasSuffix(message.From, "*") {
				message.From = fmt.Sprintf("%s%d.", message.From[0:len(message.From)-1], transport.Id)
			}
			node := NewCNode(message.From, sender, N)
			N.AddChidren(transport.Id, node)
			{
				m := kk.Message{Type: "login.done", To: message.From}
				sender.Send(&m)
			}
		}
	} else if message.Type == "ping" {
		m := kk.Message{}
		m.Type = "pong"
		m.From = N.name
		m.To = message.From
		m.Data = message.Data
		sender.Send(&m)
	}
	return N.SNode.Handle(message, sender)
}

func (N *TCPMNode) Run() error {

	defer N.fd.Close()

	t := make(chan *Transport, 64)
	e := make(chan int64, 64)

	go func() {
		var id int64 = 0
		var transport *Transport = nil
		var loopbreak = false
		var transportSet = map[int64]*Transport{}
		for !loopbreak {

			select {
			case id = <-e:
				if id == 0 {
					loopbreak = true
				} else {
					delete(transportSet, id)
					N.RemoveChidren(id)
				}
				break
			case transport = <-t:
				transportSet[transport.Id] = transport
				break
			}
		}
		close(e)
		close(t)
	}()

	for {

		conn, err := N.fd.Accept()

		if err != nil {
			return err
		}

		id := N.NewId()
		transport := NewTransport(id, NewTCPChannel(conn), N, e)

		fmt.Printf("%s id:%d\n", conn.RemoteAddr().String(), id)

		t <- transport
	}

}
