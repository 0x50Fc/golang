package micro

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/iid"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/mq"
	_ "github.com/hailongz/golang/mq/ali"
	_ "github.com/hailongz/golang/mq/nsq"
	"github.com/hailongz/golang/stat"
	"gopkg.in/redis.v5"
)

type SharedObject interface {
}

type Entry struct {
	taskType reflect.Type
	invoke   reflect.Value
}

type dbConnection struct {
	prefix string
	conn   *sql.DB
}

type redisConnection struct {
	prefix string
	conn   *redis.Client
}

type IContext interface {
	GetTrace() string
	GetConfig() interface{}
	GetApp() *App
	GetDB(name string) (*sql.DB, string, error)
	GetRedis(name string) (*redis.Client, string, error)
	GetCache(name string) (cache.ICache, error)
	GetSharedObject(key string, fn func() (SharedObject, error)) (SharedObject, error)
	SendMessage(name string, data interface{}) error
	NewID() int64
	Printf(fomrat string, v ...interface{})
	Println(v ...interface{})
}

type Context struct {
	app        *App
	trace      string
	prefix     string
	dbCount    int64
	cacheCount int64
}

func NewContext(app *App, trace string, prefix string) *Context {
	return &Context{app: app, trace: trace, prefix: prefix}
}

func (C *Context) GetTrace() string {
	return C.trace
}

func (C *Context) GetConfig() interface{} {
	return C.app.GetConfig()
}

func (C *Context) GetApp() *App {
	return C.app
}

func (C *Context) GetDB(name string) (*sql.DB, string, error) {
	C.dbCount = C.dbCount + 1
	return C.app.GetDB(name)
}

func (C *Context) GetRedis(name string) (*redis.Client, string, error) {
	C.cacheCount = C.cacheCount + 1
	return C.app.GetRedis(name)
}

func (C *Context) GetCache(name string) (cache.ICache, error) {
	C.cacheCount = C.cacheCount + 1
	cli, prefix, err := C.app.GetRedis(name)
	if err != nil {
		return nil, err
	}
	return cache.NewRedisCache(cli, prefix+"cache_"), nil
}

func (C *Context) GetSharedObject(key string, fn func() (SharedObject, error)) (SharedObject, error) {
	return C.app.GetSharedObject(key, fn)
}

func (C *Context) SendMessage(name string, data interface{}) error {
	return C.app.SendMessage(name, data)
}

func (C *Context) NewID() int64 {
	return C.app.NewID()
}

func (C *Context) Printf(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("[%s] [%s] %s %s", C.trace, C.app.qname, C.prefix, format), v...)
}

func (C *Context) Println(v ...interface{}) {
	vs := []interface{}{fmt.Sprintf("[%s] [%s] %s", C.trace, C.app.qname, C.prefix)}
	vs = append(vs, v...)
	log.Println(vs...)
}

type App struct {
	errno         int
	name          string
	config        interface{}
	entrys        map[string]*Entry
	lock          sync.RWMutex
	conns         map[string]*dbConnection
	redisConns    map[string]*redisConnection
	sharedObjects map[string]SharedObject
	q             mq.Producer
	stat          stat.Client
	id            *iid.IID
	qname         string
}

func NewApp(name string, errno int, config interface{}) (*App, error) {

	var q mq.Producer = nil
	var err error = nil

	qname := fmt.Sprintf("%s-%d-%d-%s", name,
		dynamic.IntValue(dynamic.Get(config, "aid"), 0),
		dynamic.IntValue(dynamic.Get(config, "nid"), 0),
		dynamic.StringValue(dynamic.Get(config, "version"), "1.0"))

	{
		cfg := dynamic.Get(config, "mq")
		if cfg != nil {
			stype := dynamic.StringValue(dynamic.Get(cfg, "type"), "")
			if stype != "" {
				q, err = mq.OpenProducer(stype, cfg)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	var st stat.Client = nil

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
					stat.SetLog(st,
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

	v := App{name: name,
		errno:         errno,
		config:        config,
		entrys:        map[string]*Entry{},
		conns:         map[string]*dbConnection{},
		redisConns:    map[string]*redisConnection{},
		sharedObjects: map[string]SharedObject{},
		q:             q,
		stat:          st,
		id:            id,
		qname:         qname}

	for _, s := range defaultServices {
		v.AddService(s)
	}

	return &v, nil
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

	config, err := GetConfigWithFile(configFile)

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
			j := strings.LastIndex(s, "-")
			if j >= 0 {
				s = s[j+1:]
			}
			dynamic.Set(config, "nid", dynamic.IntValue(s, 0))
		}
	}

	app, err := NewApp(dynamic.StringValue(dynamic.Get(config, "name"), "app"), int(dynamic.IntValue(dynamic.Get(config, "errno"), 0)), config)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func GetConfigWithFile(configFile string) (interface{}, error) {

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

	config, err := GetConfigWithFile(configFile)

	if err != nil {
		return nil, err
	}

	return NewApp(dynamic.StringValue(dynamic.Get(config, "name"), "app"), int(dynamic.IntValue(dynamic.Get(config, "errno"), 0)), config)
}

func (A *App) Errno() int {
	return A.errno
}

func (A *App) GetName() string {
	return A.name
}

func (A *App) GetConfig() interface{} {
	return A.config
}

func (A *App) ParseConfig(name string, config interface{}) {
	dynamic.SetValue(config, dynamic.Get(A.config, name))
}

func (A *App) NewID() int64 {
	return A.id.NewID()
}

var contextType = reflect.TypeOf((*IContext)(nil)).Elem()
var taskType = reflect.TypeOf((*Task)(nil)).Elem()
var errorType = reflect.TypeOf((*error)(nil)).Elem()

func (A *App) AddService(service Service) {
	v := reflect.ValueOf(service)
	num := v.NumMethod()

	for i := 0; i < num; i++ {
		m := v.Method(i)

		inCount := m.Type().NumIn()

		if inCount != 2 {
			continue
		}

		outCount := m.Type().NumOut()

		if outCount != 2 {
			continue
		}

		if !m.Type().In(0).AssignableTo(contextType) {
			continue
		}

		ttype := m.Type().In(1)

		if !ttype.AssignableTo(taskType) || ttype.Kind() != reflect.Ptr || ttype.Elem().Kind() != reflect.Struct {
			continue
		}

		if !m.Type().Out(1).AssignableTo(errorType) {
			continue
		}

		ttype = ttype.Elem()

		task := reflect.New(ttype).Interface().(Task)

		name := task.GetName()

		A.entrys[name] = &Entry{taskType: ttype, invoke: m}

		log.Println("[TASK]", name, taskType, m)
	}
}

func (A *App) GetEntry(name string) *Entry {
	return A.entrys[name]
}

func (A *App) Each(cb func(name string, entry *Entry) bool) {
	for key, entry := range A.entrys {
		if !cb(key, entry) {
			break
		}
	}
}

func (A *App) GetDB(name string) (*sql.DB, string, error) {

	A.lock.RLock()

	rs, ok := A.conns[name]

	A.lock.RUnlock()

	if ok {

		err := rs.conn.Ping()

		if err == nil {
			return rs.conn, rs.prefix, nil
		}

	}

	config := dynamic.GetWithKeys(A.config, []string{"db", name})

	drive := dynamic.StringValue(dynamic.Get(config, "name"), "mysql")
	url := dynamic.StringValue(dynamic.Get(config, "url"), "root:123456@tcp(127.0.0.1:3306)/kk?charset=utf8mb4")

	conn, err := sql.Open(drive, url)

	if err != nil {
		return nil, "", err
	}

	err = conn.Ping()

	if err != nil {
		return nil, "", err
	}

	conn.SetMaxIdleConns(int(dynamic.IntValue(dynamic.Get(config, "maxIdleConns"), 1)))
	conn.SetMaxOpenConns(int(dynamic.IntValue(dynamic.Get(config, "maxOpenConns"), 6)))
	conn.SetConnMaxLifetime(time.Duration(dynamic.IntValue(dynamic.Get(config, "maxLifetime"), 6)) * time.Second)

	rs = &dbConnection{prefix: dynamic.StringValue(dynamic.GetWithKeys(A.config, []string{"db", "prefix"}), ""), conn: conn}

	A.lock.Lock()
	A.conns[name] = rs
	A.lock.Unlock()

	return rs.conn, rs.prefix, nil
}

func (A *App) GetRedis(name string) (*redis.Client, string, error) {

	A.lock.RLock()

	rs, ok := A.redisConns[name]

	A.lock.RUnlock()

	if ok {

		_, err := rs.conn.Ping().Result()

		if err == nil {
			return rs.conn, rs.prefix, nil
		}
	}

	prefix := dynamic.StringValue(dynamic.GetWithKeys(A.config, []string{"redis", "prefix"}), "")
	config := dynamic.GetWithKeys(A.config, []string{"redis", name})

	if config == nil {
		return nil, prefix, errors.New("未找到 Redis 配置")
	}

	addr := dynamic.StringValue(dynamic.Get(config, "addr"), "127.0.0.1:6379")
	password := dynamic.StringValue(dynamic.Get(config, "password"), "")
	db := dynamic.IntValue(dynamic.Get(config, "db"), 0)

	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       int(db),  // use default DB
	})

	_, err := conn.Ping().Result()

	if err != nil {
		return nil, "", err
	}

	rs = &redisConnection{prefix: prefix, conn: conn}

	A.lock.Lock()
	A.redisConns[name] = rs
	A.lock.Unlock()

	return rs.conn, rs.prefix, nil
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

	if rs == nil {
		return nil, nil
	}

	A.lock.Lock()
	A.sharedObjects[key] = rs
	A.lock.Unlock()

	return rs, nil
}

func (A *App) SendMessage(name string, data interface{}) error {
	if A.q != nil {
		return A.q.Send(A.name+"/"+name, data)
	}
	return nil
}

func (E *Entry) NewTask(getValue func(name string) interface{}) Task {
	task := reflect.New(E.taskType)
	if getValue != nil {
		dynamic.EachReflect(task, func(name string, value reflect.Value) bool {
			v := getValue(name)
			dynamic.SetReflectValue(value, v)
			return true
		})
	}
	return task.Interface().(Task)
}

func (E *Entry) Handle(app *App, task Task, trace string) (interface{}, *Error) {

	log.Printf("[%s] [%s/%s] ...\n", trace, app.name, task.GetName())

	context := NewContext(app, trace, fmt.Sprintf("[%s/%s]", app.name, task.GetName()))

	if app.stat != nil {
		tags := map[string]string{}
		fields := map[string]interface{}{}
		tags["name"] = app.qname
		tags["path"] = app.name + "/" + task.GetName()
		tags["trace"] = trace
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
				err := app.stat.Write("task", tags, fields, now)
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
			log.Printf("[STAT] [%s] [%s] [%s] [in:%d] [out:%d] [use:%d] [db:%d] [cache:%d]\n", trace, app.qname, app.name+"/"+task.GetName(), tv_in, tv_out, tv_out-tv_in, context.dbCount, context.cacheCount)
		}()
	}

	rs := E.invoke.Call([]reflect.Value{reflect.ValueOf(context), reflect.ValueOf(task)})

	if rs != nil && len(rs) > 1 {

		var result interface{} = nil

		if rs[0].CanInterface() {
			result = rs[0].Interface()
		}

		var err error = nil

		if rs[1].CanInterface() && !rs[1].IsNil() {
			err = rs[1].Interface().(error)
		}

		if err == nil {
			log.Printf("[%s] [%s] [%s/%s] [OK] %v\n", trace, app.qname, app.name, task.GetName(), result)
			return result, nil
		}

		log.Printf("[%s] [%s] [%s/%s] [ERROR] %v\n", trace, app.qname, app.name, task.GetName(), err)

		e, ok := err.(*Error)

		if ok {
			e.Errno = app.errno | e.Errno
			return result, e
		} else {
			return result, NewError(app.errno, err.Error())
		}

	} else {
		log.Printf("[%s] [%s] [%s/%s] [DONE]\n", trace, app.qname, app.name, task.GetName())
	}

	return nil, nil
}

var defaultServices []Service = []Service{}

func AddDefaultService(service Service) {
	defaultServices = append(defaultServices, service)
}
