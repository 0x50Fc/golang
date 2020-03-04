package wx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"io"
	"log"
	"sort"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppRecv(app micro.IContext, task *AppRecvTask) (interface{}, error) {

	vs := []string{task.Token, task.Timestamp, task.Nonce}

	sort.Strings(vs)

	b := bytes.NewBuffer(nil)

	for _, v := range vs {
		b.WriteString(v)
	}

	m := sha1.New()
	m.Write(b.Bytes())

	sign := hex.EncodeToString(m.Sum(nil))

	if sign != task.Signature {
		return nil, micro.NewError(ERROR_SIGN, "签名错误")
	}

	var data interface{} = nil

	content := task.Content

	if task.EncodingKey != "" {

		data, err := base64.StdEncoding.DecodeString(content)

		if err != nil {
			log.Println("[EncryptedData]", content)
			return nil, err
		}

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

		cdc := cipher.NewCBCDecrypter(block, iv[0:block.BlockSize()])

		ret := make([]byte, len(data))

		cdc.CryptBlocks(ret, data)

		ret = PKCS5UnPadding(ret)

		if len(ret) > 20 {
			size := binary.BigEndian.Uint32(ret[16:20])
			if len(ret) >= int(size+20) {
				content = string(ret[20 : size+20])
			}
		}

	}

	log.Println(content)

	if task.Type == "xml" {

		data = map[string]interface{}{}

		dec := xml.NewDecoder(bytes.NewBufferString(content))

		var names = []string{}
		var value = ""

		for {

			token, err := dec.Token()

			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}

			switch token.(type) {
			case xml.StartElement:
				names = append(names, token.(xml.StartElement).Name.Local)
				value = ""
			case xml.EndElement:
				if len(names) > 1 {
					dynamic.Set(data, names[1], value)
				}
				names = names[0 : len(names)-1]
			case xml.CharData:
				value = string(token.(xml.CharData))
			}

		}
	} else {

		err := json.Unmarshal([]byte(content), &data)

		if err != nil {
			return nil, err
		}

	}

	return data, nil
}
