package app

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type IStore interface {
	Dir() string
	GetContent(path string) ([]byte, error)
	Has(path string) (bool, time.Time)
}

type storeContent struct {
	content []byte
	endTime time.Time
	modTime time.Time
}

type MemStore struct {
	dir      string
	contents map[string]*storeContent
	expires  time.Duration
	lock     sync.RWMutex
}

func NewMemStore(dir string, expires time.Duration) *MemStore {
	v := MemStore{}
	v.dir = dir
	v.contents = map[string]*storeContent{}
	v.expires = expires
	return &v
}

func (S *MemStore) Dir() string {
	return S.dir
}

func (S *MemStore) Has(path string) (bool, time.Time) {

	atime := time.Now()

	p := filepath.Clean(filepath.Join(S.dir, path))

	S.lock.RLock()

	v, ok := S.contents[p]

	S.lock.RUnlock()

	if ok && atime.Before(v.endTime) {
		return true, v.modTime
	}

	ts, err := os.Stat(p)

	if err != nil {
		return false, atime
	}

	return true, ts.ModTime()

}

func (S *MemStore) GetContent(path string) ([]byte, error) {

	atime := time.Now()

	p := filepath.Clean(filepath.Join(S.dir, path))

	S.lock.RLock()

	v, ok := S.contents[p]

	S.lock.RUnlock()

	if ok && atime.Before(v.endTime) {
		return v.content, nil
	}

	p = S.dir + p

	st, err := os.Stat(p)

	if err != nil {
		return nil, err
	}

	if ok {
		if v.modTime.Equal(st.ModTime()) {
			v.endTime = atime.Add(S.expires)
			return v.content, nil
		}
	}

	fd, err := os.Open(p)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	b, err := ioutil.ReadAll(fd)

	if err != nil {
		return nil, err
	}

	vv := &storeContent{}
	vv.content = b
	vv.modTime = st.ModTime()
	vv.endTime = atime.Add(S.expires)

	S.lock.Lock()
	S.contents[p] = vv
	S.lock.Unlock()

	return b, nil
}
