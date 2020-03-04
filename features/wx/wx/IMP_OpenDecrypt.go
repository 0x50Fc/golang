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
	"strings"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func parseXML(s string) (map[string]interface{}, error) {

	dec := xml.NewDecoder(bytes.NewBufferString(s))

	ret := map[string]interface{}{}

	var names = []string{}
	var value = bytes.NewBuffer(nil)

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
			value.Reset()
		case xml.EndElement:
			if len(names) > 1 {
				ret[names[1]] = strings.Trim(value.String(), " \n\r")
			}
			names = names[0 : len(names)-1]
		case xml.CharData:
			value.Write(token.(xml.CharData))
		}

	}

	return ret, nil
}

func (S *Service) OpenDecrypt(app micro.IContext, task *OpenDecryptTask) (interface{}, error) {

	ret, err := parseXML(task.Content)

	if err != nil {
		return nil, err
	}

	log.Println(ret)

	vs := []string{task.Token, task.Timestamp, task.Nonce, dynamic.StringValue(dynamic.Get(ret, "Encrypt"), "")}

	sort.Strings(vs)

	b := bytes.NewBuffer(nil)

	for _, v := range vs {
		b.WriteString(v)
	}

	m := sha1.New()
	m.Write(b.Bytes())

	sign := hex.EncodeToString(m.Sum(nil))

	log.Println(vs, sign, task.Signature)

	if sign != task.Signature {
		return nil, micro.NewError(ERROR_SIGN, "签名错误")
	}

	if task.EncodingKey != "" {

		var decrypt map[string]interface{} = nil

		encrypt := dynamic.StringValue(dynamic.Get(ret, "Encrypt"), "")

		data, err := base64.StdEncoding.DecodeString(encrypt)

		if err != nil {
			log.Println("[EncryptedData]", encrypt)
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

		s := make([]byte, len(data))

		cdc.CryptBlocks(s, data)

		s = PKCS5UnPadding(s)

		if len(s) > 20 {
			size := binary.BigEndian.Uint32(s[16:20])
			if len(s) >= int(size+20) {
				decrypt, err = parseXML(string(s[20 : size+20]))
				if err != nil {
					log.Println("[parseXML]", s[20:size+20])
					return nil, err
				}
			}
		}

		ret["Decrypt"] = decrypt
	}

	return ret, nil
}
