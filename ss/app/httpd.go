package app

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/svc"
)

var ClientIpKeys = []string{"X-Real-IP", "x-real-ip", "X-Forwarded-For", "x-forwarded-for"}
var TraceIDKeys = []string{"Trace-ID", "trace-id"}
var HostKeys = []string{"X-Host", "x-host", "Host", "host"}
var SchemaKeys = []string{"X-Schema", "x-schema"}

func GetHeaderValue(header http.Header, keys []string, defaultValue string) string {

	for _, key := range keys {

		v := header.Get(key)

		if v != "" {
			return v
		}
	}

	return defaultValue
}

func HandleFunc(app *App, s svc.Server) func(http.ResponseWriter, *http.Request) {

	return func(resp http.ResponseWriter, req *http.Request) {

		if s != nil {
			err := s.In()
			if err != nil {
				resp.WriteHeader(http.StatusServiceUnavailable)
				resp.Write([]byte(err.Error()))
				return
			}
			defer s.Out()
		}

		input := Input{
			Method:   req.Method,
			Path:     req.URL.Path,
			Host:     req.Host,
			Protocol: req.Proto,
			Query:    req.URL.RawQuery,
			Header:   map[string]string{},
			Cookie:   map[string]string{},
		}

		{
			var ip = GetHeaderValue(req.Header, ClientIpKeys, "")

			if ip == "" {
				ip = req.RemoteAddr
				i := strings.LastIndex(ip, ":")
				if i > 0 {
					ip = ip[:i]
				}
			}

			input.ClientIp = ip
		}

		{
			var v = GetHeaderValue(req.Header, TraceIDKeys, "")

			if v != "" {
				input.Trace = v
			}
		}

		{
			var v = GetHeaderValue(req.Header, HostKeys, "")

			if v != "" {
				input.Host = v
			}
		}

		{
			var v = GetHeaderValue(req.Header, SchemaKeys, "")

			if v != "" {
				input.Protocol = v
			}
		}

		var maxMemory = dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"httpd", "maxMemory"}), 1024*1024*1024)

		ctype := req.Header.Get("Content-Type")

		if strings.Contains(ctype, "multipart/form-data") {
			input.Data = map[string]interface{}{}
			req.ParseMultipartForm(maxMemory)
			if req.MultipartForm != nil {
				for key, values := range req.MultipartForm.Value {
					dynamic.Set(input.Data, key, values[0])
				}
				for key, values := range req.MultipartForm.File {
					dynamic.Set(input.Data, key, values[0])
				}
			}
		} else if strings.Contains(ctype, "json") {

			b, err := ioutil.ReadAll(req.Body)

			defer req.Body.Close()

			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				resp.Write([]byte(err.Error()))
				return
			}

			log.Println(string(b))

			err = json.Unmarshal(b, &input.Data)

			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				resp.Write([]byte(err.Error()))
				return
			}

		} else if strings.Contains(ctype, "text") {

			b, err := ioutil.ReadAll(req.Body)

			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				resp.Write([]byte(err.Error()))
				return
			}

			defer req.Body.Close()

			input.Data = map[string]interface{}{}

			req.ParseForm()

			for key, values := range req.Form {
				dynamic.Set(input.Data, key, values[0])
			}

			dynamic.Set(input.Data, "$body", string(b))

		} else {

			input.Data = map[string]interface{}{}

			req.ParseForm()

			for key, values := range req.Form {
				dynamic.Set(input.Data, key, values[0])
			}

		}

		for key, h := range req.Header {
			input.Header[key] = h[0]
		}

		for _, cookie := range req.Cookies() {
			input.Cookie[cookie.Name] = cookie.Value
		}

		{
			var v = req.Header.Get("cookie")
			if v != "" {
				for _, item := range strings.Split(v, ";") {
					vs := strings.Split(item, "=")
					if len(vs) > 1 {
						input.Cookie[vs[0]] = vs[1]
					}
				}
			}
		}

		output, err := app.Exec("/main.js", &input)

		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(err.Error()))
			return
		}

		if output == nil {
			resp.WriteHeader(http.StatusNotImplemented)
			resp.Write([]byte(err.Error()))
			return
		}

		if output.Status == 200 && output.Data == nil {
			resp.WriteHeader(http.StatusNotImplemented)
			resp.Write([]byte(err.Error()))
			return
		}

		{
			h := resp.Header()
			for key, value := range output.Header {
				h[key] = []string{value}
			}
		}

		if output.Status != 200 {
			resp.WriteHeader(output.Status)
			return
		}

		if output.Data == nil {
			resp.Write([]byte{})
			return
		}

		{
			b, ok := output.Data.(string)
			if ok {
				resp.Write([]byte(b))
				return
			}
		}

		{
			b, ok := output.Data.([]byte)
			if ok {
				resp.Write(b)
				return
			}
		}

		{
			b, _ := json.Marshal(output.Data)
			resp.Write(b)
			return
		}

	}
}
