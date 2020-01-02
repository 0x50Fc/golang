package svc

import (
	"errors"
	"log"
	"sync"
)

type Server interface {
	In() error
	Out()
	Done()
}

type server struct {
	count int64
	done  chan bool
	lock  sync.Mutex
}

func New() Server {
	return &server{}
}

func (S *server) In() error {
	S.lock.Lock()
	defer S.lock.Unlock()
	if S.done != nil {
		return errors.New("server is done")
	}
	S.count = S.count + 1
	// log.Printf("[SVC] [IN] %d\n", S.count)
	return nil
}

func (S *server) Out() {
	S.lock.Lock()
	defer S.lock.Unlock()
	S.count = S.count - 1
	// log.Printf("[SVC] [OUT] %d\n", S.count)
	if S.count == 0 && S.done != nil {
		S.done <- true
	}
}

func (S *server) Done() {
	var done chan bool = nil
	S.lock.Lock()
	if S.done == nil && S.count > 0 {
		done = make(chan bool)
		S.done = done
		log.Printf("[SVC] [DONE] [WAIT] %d\n", S.count)
	}
	S.lock.Unlock()
	if done != nil {
		<-done
		close(done)
	}
}
