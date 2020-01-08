package ali

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	Ali "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/features/oss/oss"
)

func init() {
	oss.AddSource("ali", func(config interface{}) (oss.ISource, error) {
		return NewAliSource(dynamic.StringValue(dynamic.Get(config, "endpoint"), ""),
			dynamic.StringValue(dynamic.Get(config, "accessKey"), ""),
			dynamic.StringValue(dynamic.Get(config, "secretKey"), ""),
			dynamic.StringValue(dynamic.Get(config, "bucket"), ""),
			dynamic.StringValue(dynamic.Get(config, "baseURL"), ""))
	})
}

type AliSource struct {
	client    *Ali.Client
	bucket    *Ali.Bucket
	baseURL   string
	endpoint  string
	accessKey string
	secretKey string
}

func NewAliSource(endpoint string, accessKey string, secretKey string, bucket string, baseURL string) (*AliSource, error) {

	cli, err := Ali.New(endpoint, accessKey, secretKey)
	if err != nil {
		return nil, err
	}

	buk, err := cli.Bucket(bucket)

	if err != nil {
		return nil, err
	}

	return &AliSource{client: cli, bucket: buk, baseURL: baseURL, endpoint: endpoint, accessKey: accessKey, secretKey: secretKey}, nil
}

func (S *AliSource) Get(key string) ([]byte, error) {
	rd, err := S.bucket.GetObject(key)
	if err != nil {
		return nil, err
	}
	defer rd.Close()
	return ioutil.ReadAll(rd)
}

func (S *AliSource) GetURL(key string) string {
	return fmt.Sprintf("%s%s", S.baseURL, key)
}

func (S *AliSource) GetSignURL(key string, expires time.Duration) (string, error) {
	u, err := S.bucket.SignURL(key, Ali.HTTPGet, int64(expires/time.Second))
	if err != nil {
		return "", err
	}
	return u, nil
}

func (S *AliSource) Put(key string, data []byte, header map[string]string) error {

	options := []Ali.Option{}

	if header != nil {
		for key, value := range header {
			if key == "Content-Type" {
				options = append(options, Ali.ContentType(value))
			} else if key == "Content-Encoding" {
				options = append(options, Ali.ContentEncoding(value))
			} else if key == "Content-Disposition" {
				options = append(options, Ali.ContentDisposition(value))
			} else {
				options = append(options, Ali.Meta(key, value))
			}
		}
	}

	err := S.bucket.PutObject(key, bytes.NewReader(data), options...)

	if err != nil {
		return err
	}

	return nil
}

func (S *AliSource) PutSignURL(key string, expires time.Duration) (string, error) {
	u, err := S.bucket.SignURL(key, Ali.HTTPPut, int64(expires/time.Second))
	if err != nil {
		return "", err
	}
	return u, nil
}

func (S *AliSource) PostSignURL(key string, expires time.Duration) (string, map[string]string, error) {

	u, err := url.Parse(S.endpoint)

	if err != nil {
		return "", nil, err
	}

	data := map[string]string{}

	data["OSSAccessKeyId"] = S.accessKey
	data["key"] = key

	policy := fmt.Sprintf(`{"expiration": "%s","conditions":[{"bucket": "%s" },["content-length-range", 0, 1073741824]]}`,
		time.Now().Add(expires).Format("2006-01-02T15:04:05Z"), S.bucket.BucketName)

	log.Println(policy)

	v := base64.StdEncoding.EncodeToString([]byte(policy))

	data["policy"] = v

	m := hmac.New(sha1.New, []byte(S.secretKey))

	m.Write([]byte(v))

	data["Signature"] = base64.StdEncoding.EncodeToString(m.Sum(nil))

	s := fmt.Sprintf("%s://%s.%s", u.Scheme, S.bucket.BucketName, u.Host)

	log.Println(s, data)

	return s, data, nil
}

func (S *AliSource) Del(key string) error {
	err := S.bucket.DeleteObject(key)
	if err != nil {
		return err
	}
	return nil
}

func (S *AliSource) Has(key string) error {
	_, err := S.bucket.GetObjectACL(key)
	if err != nil {
		return err
	}
	return nil
}
