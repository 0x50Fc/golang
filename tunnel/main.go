package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/hailongz/golang/tunnel/kk"
	"github.com/hailongz/golang/tunnel/tn"
)

var flag_name = flag.String("name", "plus.proj.", "--name <name>")
var flag_addr = flag.String("addr", ":9080", "--addr <addr>")
var flag_help = flag.Bool("help", false, "--help")
var flag_remote = flag.String("remote", "", "--remote <name>@<addr>,<name>@<addr>")

func main() {

	flag.Parse()

	if *flag_help {
		flag.PrintDefaults()
		return
	}

	resp := tn.NewRespService()
	req := tn.NewRService()

	resp.AddHandler("/", func(req *kk.Request) (*kk.Response, error) {
		resp := kk.Response{}
		resp.Code = 200
		resp.Type = req.Type
		resp.Data = req.Data
		fmt.Println(req)
		return &resp, nil
	})

	m, err := tn.NewTCPMNode(*flag_name, *flag_addr)

	if err != nil {
		log.Panicln(err)
	}

	{
		vs := strings.Split(*flag_remote, ",")
		for _, v := range vs {
			ns := strings.Split(v, "@")
			if len(ns) > 1 {
				m.AddParent(m.NewId(), tn.NewTCPCNode(ns[0], ns[1]))
				fmt.Printf("[REMOTE] %s@%s\n", ns[0], ns[1])
			}
		}
	}
	m.AddService(req)
	m.AddService(resp)

	fmt.Printf("[LOCAL] %s@%s\n", *flag_name, *flag_addr)

	m.Run()

}
