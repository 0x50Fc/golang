package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hailongz/golang/tunnel/kk"
	"github.com/hailongz/golang/tunnel/tn"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: tn-cli <name>@<addr>")
		return
	}

	s := os.Args[1]

	vs := strings.Split(s, "@")

	if len(vs) < 2 {
		fmt.Println("Usage: tn-cli <name>@<addr>")
		return
	}

	c := tn.NewTCPCNode(vs[0], vs[1])

	req := tn.NewRService()

	c.AddService(req)

	defer c.Close()

	if strings.HasSuffix(vs[0], "*") {
		m := kk.Message{Type: "login"}
		c.Send(&m)
	}

	rd := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("%s@%s > ", vs[0], vs[1])

		line, err := rd.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}

		ns := strings.Split(line, " ")

		switch ns[0] {
		case "req":
			if len(ns) > 3 {
				resp, err := req.Send(c, ns[1], ns[2], "text", []byte(ns[3]), nil, time.Second*6)
				if err != nil {
					fmt.Println(err)
					continue
				} else if resp.Code == 200 {
					if resp.Type == "text" {
						fmt.Println(string(resp.Data))
					} else {
						fmt.Println(hex.EncodeToString(resp.Data))
					}
				} else {
					fmt.Printf("%d\n", resp.Code)
				}
			} else {
				fmt.Println("req <to> <uri> <data>")
			}
			break
		case "help":
			fmt.Println("req <to> <uri> <data>")
			break
		}

	}

}
