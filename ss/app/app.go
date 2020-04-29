package app

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/iid"
	"github.com/hailongz/golang/stat"
)

type SharedObject interface {
}

type App struct {
	store         IStore
	config        interface{}
	lock          sync.RWMutex
	sharedObjects map[string]SharedObject
	id            *iid.IID
	stat          stat.Client
	qname         string
}

func NewApp(store IStore, config interface{}) (*App, error) {

	var st stat.Client = nil
	var err error = nil

	qname := fmt.Sprintf("%s-%d-%d-%s", dynamic.StringValue(dynamic.Get(config, "name"), "app"),
		dynamic.IntValue(dynamic.Get(config, "aid"), 0),
		dynamic.IntValue(dynamic.Get(config, "nid"), 0),
		dynamic.StringValue(dynamic.Get(config, "version"), "1.0"))

	{
		cfg := dynamic.Get(config, "stat")
		if cfg != nil {
			stype := dynamic.StringValue(dynamic.Get(cfg, "type"), "")
			if stype != "" {
				st, err = stat.OpenClient(stype, cfg)

				if err != nil {
					return nil, err
				}

				if st != nil {
					stat.Sys(st,
						time.Duration(dynamic.IntValue(dynamic.Get(cfg, "keepalive"), 6))*time.Second,
						qname)
				}
			}

		} else {
			stat.SysLog(60*time.Second, qname)
		}
	}

	{
		cfg := dynamic.Get(config, "log")
		if cfg != nil {
			stype := dynamic.StringValue(dynamic.Get(cfg, "type"), "")
			if stype != "" {

				cli, err := stat.OpenClient(stype, cfg)

				if err != nil {
					return nil, err
				}

				if cli != nil {
					stat.SetLog(cli,
						qname)
				}
			}

		}
	}

	var id *iid.IID = nil

	{
		cfg := dynamic.Get(config, "iid")
		id = iid.NewIID(dynamic.IntValue(dynamic.Get(cfg, "aid"), 0), dynamic.IntValue(dynamic.Get(cfg, "nid"), 0))
	}

	return &App{store: store, config: config, id: id, stat: st, sharedObjects: map[string]SharedObject{}, qname: qname}, nil
}

func NewAppWithEnv() (*App, error) {

	configFile := "./app.json"

	{
		s := os.Getenv("KK_CONFIG_FILE")
		if s != "" {
			configFile = s
			log.Println("[KK_CONFIG_FILE]", s)
		}
	}

	config, err := OpenConfigFile(configFile)

	if err != nil {
		return nil, err
	}

	{
		s := os.Getenv("KK_VERSION")
		if s != "" {
			dynamic.Set(config, "version", s)
			log.Println("[KK_VERSION]", s)
		}
	}

	{
		s := os.Getenv("KK_AID")
		if s != "" {
			dynamic.Set(config, "aid", dynamic.IntValue(s, 0))
			log.Println("[KK_AID]", s)
		}
	}

	{
		s := os.Getenv("KK_NID")
		if s != "" {
			dynamic.Set(config, "nid", dynamic.IntValue(s, 0))
			log.Println("[KK_NID]", s)
		}
	}

	{
		s := os.Getenv("KK_NODE")
		if s != "" {
			log.Println("[KK_NODE]", s)
			name := s
			j := strings.LastIndex(s, "-")
			if j >= 0 {
				name = s[:j]
				s = s[j+1:]
			}
			dynamic.Set(config, "nid", dynamic.IntValue(s, 0))
			dynamic.Set(config, "name", name)
		}
	}

	app, err := NewApp(NewMemStore("./", time.Second*time.Duration(dynamic.IntValue(dynamic.GetWithKeys(config, []string{"store", "expires"}), 6))), config)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func OpenConfigFile(configFile string) (interface{}, error) {

	var config interface{} = nil

	fd, err := os.Open(configFile)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	b, err := ioutil.ReadAll(fd)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func NewAppWithConfigFile(configFile string) (*App, error) {

	config, err := OpenConfigFile(configFile)

	if err != nil {
		return nil, err
	}

	return NewApp(NewMemStore("./", time.Second*time.Duration(dynamic.IntValue(dynamic.GetWithKeys(config, []string{"store", "expires"}), 6))), config)
}

func (A *App) GetConfig() interface{} {
	return A.config
}

func (A *App) GetStore() IStore {
	return A.store
}

func (A *App) QName() string {
	return A.qname
}

func (A *App) GetSharedObject(key string, fn func() (SharedObject, error)) (SharedObject, error) {

	A.lock.RLock()

	rs, ok := A.sharedObjects[key]

	A.lock.RUnlock()

	if ok {
		return rs, nil
	}

	rs, err := fn()

	if err != nil {
		return nil, err
	}

	A.lock.Lock()
	A.sharedObjects[key] = rs
	A.lock.Unlock()

	return rs, nil
}

type Openlib func(app *App, ctx *duktape.Context, trace string)

var openlibs []Openlib = []Openlib{}

func AddOpenlib(openlib Openlib) {
	openlibs = append(openlibs, openlib)
}

type Input struct {
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	Protocol    string            `json:"protocol"`
	Host        string            `json:"host"`
	Query       string            `json:"query"`
	QueryObject map[string]string `json:"queryObject"`
	SessionId   string            `json:"sessionId"`
	ClientIp    string            `json:"clientIp"`
	Data        interface{}       `json:"data"`
	Cookie      map[string]string `json:"cookie"`
	Header      map[string]string `json:"header"`
	Trace       string            `json:"trace"`
}

type Output struct {
	Header map[string]string `json:"header"`
	Cookie map[string]string `json:"cookie"`
	Data   interface{}       `json:"data"`
	Status int               `json:"status"`
}

func (A *App) NewID() int64 {
	return A.id.NewID()
}

func (A *App) Exec(path string, input *Input) (*Output, error) {

	if input.Trace == "" {
		input.Trace = fmt.Sprintf("%d", A.NewID())
	}

	if input.QueryObject == nil {

		input.QueryObject = map[string]string{}

		if input.Query != "" {
			for _, item := range strings.Split(input.Query, "&") {
				vs := strings.Split(item, "=")
				if len(vs) > 1 {
					input.QueryObject[vs[0]], _ = url.QueryUnescape(vs[1])
				}
			}
		}

	}

	if A.stat != nil {
		tags := map[string]string{}
		fields := map[string]interface{}{}
		tags["name"] = A.qname
		tags["path"] = input.Path
		tags["trace"] = input.Trace
		tv_in := int32(time.Now().UnixNano() / int64(time.Millisecond))
		fields["name"] = tags["name"]
		fields["path"] = tags["path"]
		fields["trace"] = tags["trace"]
		fields["in"] = tv_in
		defer func() {
			now := time.Now()
			tv_out := int32(now.UnixNano() / int64(time.Millisecond))
			fields["out"] = tv_out
			fields["use"] = tv_out - tv_in
			go func() {
				err := A.stat.Write("task", tags, fields, now)
				if err != nil {
					log.Println("[STAT] [TASK] [ERROR]", err)
				}
			}()
		}()
	} else {
		tv_in := time.Now().UnixNano() / int64(time.Millisecond)
		defer func() {
			now := time.Now()
			tv_out := now.UnixNano() / int64(time.Millisecond)
			log.Printf("[STAT] [%s] [%s] [%s] [in:%d] [out:%d] [use:%d] [clientIp:%s] [session:%s]\n", input.Trace, A.qname, input.Path, tv_in, tv_out, tv_out-tv_in, input.ClientIp, input.SessionId)
		}()
	}

	store := A.GetStore()

	b, err := store.GetContent(path)

	if err != nil {
		return nil, err
	}

	ctx := duktape.New()

	defer ctx.Recycle()

	for _, openlib := range openlibs {
		openlib(A, ctx, input.Trace)
	}

	output := Output{
		Header: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Cookie: map[string]string{},
		Data: map[string]interface{}{
			"errno":  ERRNO_NOT_IMPLEMENTED,
			"errmsg": "未实现",
		},
		Status: 200,
	}

	sessionKey := dynamic.StringValue(dynamic.Get(A.GetConfig(), "sessionKey"), "kk")

	input.SessionId = input.Cookie[sessionKey]

	if input.SessionId == "" {
		m := md5.New()
		m.Write([]byte(fmt.Sprintf("FJI)#(*YRH(GUN!OWJF%dKFLDJKFO%d%s", time.Now().UnixNano(), rand.Int(), input.Trace)))
		input.SessionId = hex.EncodeToString(m.Sum(nil))
		output.Cookie[sessionKey] = input.SessionId
	}

	ctx.PushObject()
	ctx.PushObject()
	ctx.Dup(-1)
	ctx.PutGlobalString("exports")
	ctx.PutPropString(-2, "exports")
	ctx.PutGlobalString("module")

	Encode(ctx, input)
	ctx.PutGlobalString("input")

	Encode(ctx, &output)
	ctx.PutGlobalString("output")

	Encode(ctx, dynamic.Get(A.GetConfig(), "app"))
	ctx.PutGlobalString("app")

	ctx.PushString(path)
	ctx.CompileStringFilename(0, string(b))

	if ctx.Pcall(0) != duktape.ExecSuccess {
		err = ctx.ToError(-1)
		ctx.Pop()
		return nil, err
	} else {
		ctx.Pop()
	}

	ctx.GetGlobalString("output")

	Unmarshal(ctx, -1, &output)

	ctx.Pop()

	if len(output.Cookie) > 0 {

		cookie := bytes.NewBuffer(nil)

		for key, value := range output.Cookie {
			cookie.WriteString(key)
			cookie.WriteString("=")
			cookie.WriteString(value)
			cookie.WriteString("; path=/; httponly;")
		}

		output.Header["Set-Cookie"] = cookie.String()
	}

	output.Header["Trace-ID"] = input.Trace

	return &output, nil
}
