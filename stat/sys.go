package stat

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sort"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/shirou/gopsutil/process"
)

func GetSysObject() *Object {

	tags := map[string]string{}
	fields := map[string]interface{}{}

	tags["platform"] = "golang"
	tags["os"] = runtime.GOOS
	tags["arch"] = runtime.GOARCH

	fields["CPU_COUNT"] = runtime.GOMAXPROCS(0)

	{
		proc, _ := process.NewProcess(int32(os.Getpid()))

		if proc != nil {

			{
				used, _ := proc.CPUPercent()
				threads, _ := proc.NumThreads()

				fields["CPU_USED"] = int(used * 100)
				fields["THREAD_COUNT"] = threads
			}
			{
				mem, _ := proc.MemoryInfo()
				used, _ := proc.MemoryPercent()
				fields["MEM_USED"] = int(used * 100)
				fields["MEM_RSS"] = int(mem.RSS)
				fields["MEM_VMS"] = int(mem.VMS)
			}
			{
				conns, _ := proc.Connections()
				if conns != nil {
					fields["NET_COUNT"] = len(conns)
				}
			}
			{
				st, _ := proc.IOCounters()
				if st != nil {
					fields["IO_RD_COUNT"] = int(st.ReadCount)
					fields["IO_WD_COUNT"] = int(st.WriteCount)
					fields["IO_RD_BYTES"] = int(st.ReadBytes)
					fields["IO_WD_COUNT"] = int(st.WriteBytes)
				}
			}
		}
	}

	fields["ROUTINE_COUNT"] = runtime.NumGoroutine()

	{
		m := runtime.MemStats{}
		runtime.ReadMemStats(&m)
		fields["GC_LAST"] = int(m.LastGC / uint64(time.Millisecond))
		fields["GC_COUNT"] = int(m.NumGC)
	}

	return &Object{Name: "sys", Tags: tags, Fields: fields, Tv: time.Now().UnixNano()}
}

func Sys(st Client, keepalive time.Duration, name string) {

	go func() {

		for {

			time.Sleep(keepalive)

			v := GetSysObject()

			v.Tags["name"] = name

			err := st.Write(v.Name, v.Tags, v.Fields, time.Unix(v.Tv/1e9, v.Tv%1e9))

			if err != nil {
				log.Println("[STAT] [SYS] [ERROR]", err)
			}

		}

	}()

}

func SysLog(keepalive time.Duration, name string) {

	go func() {

		for {

			time.Sleep(keepalive)

			v := GetSysObject()

			args := []interface{}{}

			args = append(args, fmt.Sprintf("[%s]", name), "[SYS]")

			m := map[string]interface{}{}
			keys := []string{}

			{
				for key, value := range v.Tags {
					_, ok := m[key]
					if !ok {
						keys = append(keys, key)
						m[key] = value
					}
				}
			}

			{
				for key, value := range v.Fields {
					_, ok := m[key]
					if !ok {
						keys = append(keys, key)
						m[key] = value
					}
				}
			}

			sort.Strings(keys)

			{
				for _, key := range keys {
					args = append(args, fmt.Sprintf("[%s:%s]", key, dynamic.StringValue(dynamic.Get(m, key), "")))
				}
			}

			log.Println(args...)

		}

	}()

}

func HandleFunc() func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/__stat" {
			b, _ := json.Marshal(GetSysObject())
			resp.Header().Set("Content-Type", "application/json; charset=utf-8")
			resp.Write(b)
			return
		}
		resp.WriteHeader(200)
		resp.Write([]byte{})
	}
}
