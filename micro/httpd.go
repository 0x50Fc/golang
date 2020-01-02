package micro

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/svc"
)

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

		prefix := "/" + app.GetName() + "/"

		if !strings.HasPrefix(req.URL.Path, prefix) {
			resp.WriteHeader(http.StatusNotFound)
			resp.Write([]byte("Not Found"))
			return
		}

		name := req.URL.Path[len(prefix):]

		entry := app.GetEntry(name)
		if entry == nil {
			resp.WriteHeader(http.StatusNotFound)
			resp.Write([]byte("Not Found"))
			return
		}

		var inputData interface{} = nil
		var maxMemory = dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"httpd", "maxMemory"}), 1024*1024*1024)

		ctype := req.Header.Get("Content-Type")

		if strings.Contains(ctype, "multipart/form-data") {
			inputData = map[string]interface{}{}
			req.ParseMultipartForm(maxMemory)
			if req.MultipartForm != nil {
				for key, values := range req.MultipartForm.Value {
					dynamic.Set(inputData, key, values[0])
				}
				for key, values := range req.MultipartForm.File {
					dynamic.Set(inputData, key, values[0])
				}
			}
		} else if strings.Contains(ctype, "json") {

			b, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()

			if err == nil {
				json.Unmarshal(b, &inputData)
			}

		} else {

			inputData = map[string]interface{}{}

			req.ParseForm()

			for key, values := range req.Form {
				dynamic.Set(inputData, key, values[0])
			}

		}

		var trace string = req.Header.Get("Trace-ID")

		if trace == "" {
			trace = fmt.Sprintf("%d", app.NewID())
			resp.Header().Set("Trace-ID", trace)
		}

		log.Printf(fmt.Sprintf("[%s]", trace), "[IN]", req.URL.Path, inputData)

		task := entry.NewTask(func(name string) interface{} {
			return dynamic.Get(inputData, name)
		})

		rs, err := entry.Handle(app, task, trace)

		if err != nil {
			b, _ := json.Marshal(err)
			resp.Header().Set("Content-Type", "application/json; charset=utf-8")
			resp.Write(b)
			return
		}

		outputData := map[string]interface{}{"errno": ERRNO_OK}

		if rs != nil {
			outputData["data"] = rs
		}

		b, _ := json.Marshal(outputData)
		resp.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp.Write(b)
		return

	}
}
