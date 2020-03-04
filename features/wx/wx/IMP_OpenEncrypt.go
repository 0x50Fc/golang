package wx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func encodeXML(data interface{}) []byte {
	s := bytes.NewBuffer(nil)

	s.WriteString("<xml>")

	dynamic.Each(data, func(key interface{}, value interface{}) bool {
		skey := dynamic.StringValue(key, "")
		s.WriteString("<")
		s.WriteString(skey)
		s.WriteString("><![CDATA[")
		s.WriteString(dynamic.StringValue(value, ""))
		s.WriteString("]]>")
		s.WriteString("</")
		s.WriteString(skey)
		s.WriteString(">")
		return true
	})

	s.WriteString("</xml>")

	return s.Bytes()
}

func encodeXMLWithJSON(v string) ([]byte, error) {

	var data interface{} = nil

	err := json.Unmarshal([]byte(v), &data)

	if err != nil {
		return nil, err
	}

	return encodeXML(data), nil
}

func (S *Service) OpenEncrypt(app micro.IContext, task *OpenEncryptTask) (*OpenEncryptData, error) {

	appid := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"open", "appid"}), "")
	nonce := dynamic.StringValue(task.Nonce, fmt.Sprintf("%d", rand.Int()))
	timestamp := dynamic.StringValue(task.Timestamp, fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)))
	content := ""

	enc, err := encodeXMLWithJSON(task.Content)

	if err != nil {
		return nil, err
	}

	pk := bytes.NewBuffer(nil)

	{
		rs := fmt.Sprintf("%d", rand.Int())

		for len(rs) < 16 {
			rs = fmt.Sprintf("%s%d", rs, rand.Int())
		}

		pk.Write([]byte(rs)[0:16])
	}

	{
		n := make([]byte, 4)
		binary.BigEndian.PutUint32(n, uint32(len(enc)))
		pk.Write(n)
	}

	pk.Write(enc)
	pk.WriteString(appid)

	{
		key, err := base64.StdEncoding.DecodeString(task.EncodingKey + "=")

		if err != nil && len(key) >= 16 {
			log.Println("[EncodingKey]", err, task.EncodingKey+"=")
			return nil, err
		}

		iv := key[0:16]

		block, err := aes.NewCipher(key)

		if err != nil {
			log.Println("[aes.NewCipher]", err)
			return nil, err
		}

		cdc := cipher.NewCBCEncrypter(block, iv[0:block.BlockSize()])

		bs := PKCS5Padding(pk.Bytes(), block.BlockSize())

		s := make([]byte, len(bs))

		cdc.CryptBlocks(s, bs)

		content = base64.StdEncoding.EncodeToString(s)

	}

	vs := []string{task.Token, timestamp, nonce, content}

	sort.Strings(vs)

	b := bytes.NewBuffer(nil)

	for _, v := range vs {
		b.WriteString(v)
	}

	m := sha1.New()
	m.Write(b.Bytes())

	sign := hex.EncodeToString(m.Sum(nil))

	log.Println(vs, sign)

	s := fmt.Sprintf(`<xml>
		<Encrypt><![CDATA[%s]]></Encrypt>
		<MsgSignature><![CDATA[%s]]></MsgSignature>
		<TimeStamp>%s</TimeStamp>
		<Nonce><![CDATA[%s]]></Nonce>
	</xml>`,
		content, sign, timestamp, nonce)

	return &OpenEncryptData{Nonce: nonce, Timestamp: timestamp, Signature: sign, Content: s}, nil
}
