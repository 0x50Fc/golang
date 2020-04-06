package tn

import (
	"github.com/hailongz/golang/tunnel/kk"
)

type Break struct {
}

func (B *Break) Error() string {
	return "break"
}

type Sender interface {
	Addr() string
	Send(message *kk.Message) error
	Close() error
}

type Handler interface {
	Handle(message *kk.Message, sender Sender) error
}

type Node interface {
	Name() string
	Match(name string) bool
	Send(message *kk.Message) error
	Priority() int
}

type Service interface {
	Match(message *kk.Message) bool
	Handle(message *kk.Message, sender Sender) error
}

type Channel interface {
	Addr() string
	Read() (*kk.Message, error)
	Write(message *kk.Message) error
	Close() error
}
