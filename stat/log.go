package stat

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

type ClientLog struct {
	ch   chan string
	data *bytes.Buffer
	rd   *bufio.Reader
}

func (L *ClientLog) Write(p []byte) (n int, err error) {

	L.data.Write(p)

	for {

		s, err := L.rd.ReadString('\n')

		if err != nil {
			break
		}

		L.ch <- s
	}

	os.Stdout.Write(p)

	return len(p), nil
}

func SetLog(st Client, name string) {

	log.SetPrefix(fmt.Sprintf("%s[%s] ", name, log.Prefix()))

	data := bytes.NewBuffer(nil)

	ch := make(chan string, 20480)

	log.SetOutput(&ClientLog{ch: ch, data: data, rd: bufio.NewReader(data)})

	go func() {

		for {

			s, ok := <-ch

			if !ok {
				break
			}

			err := st.Write("log", map[string]string{"name": name}, map[string]interface{}{"text": s}, time.Now())

			if err != nil {
				fmt.Println("[LOG] [ERROR]", err)
			}

		}
	}()

}
