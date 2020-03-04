package wx

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) MpConfig(app micro.IContext, task *MpConfigTask) (*MPConfigData, error) {

	ticket, err := MP_GetTicket(app, task.Appid, TicketType_Jsapi)

	if err != nil {
		return nil, err
	}

	noncestr := dynamic.StringValue(task.Noncestr, MP_NewNonceStr())

	timestamp := dynamic.IntValue(task.Timestamp, time.Now().Unix())

	data := map[string]interface{}{}
	data["noncestr"] = noncestr
	data["timestamp"] = timestamp
	data["jsapi_ticket"] = ticket
	data["url"] = task.Url

	keys := []string{}

	for key, _ := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	b := bytes.NewBuffer(nil)

	for i, key := range keys {
		if i != 0 {
			b.WriteString("&")
		}
		b.WriteString(key)
		b.WriteString("=")
		b.WriteString(dynamic.StringValue(data[key], ""))
	}

	m := sha1.New()
	m.Write(b.Bytes())

	return &MPConfigData{Appid: task.Appid, Signature: hex.EncodeToString(m.Sum(nil)), NonceStr: noncestr, Timestamp: timestamp}, nil

}
