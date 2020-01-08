package oss

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/features/oss/oss"
	"github.com/minio/minio-go"
)

func init() {
	oss.AddSource("minio", func(config interface{}) (oss.ISource, error) {
		return NewMinioSource(dynamic.StringValue(dynamic.Get(config, "endpoint"), ""),
			dynamic.StringValue(dynamic.Get(config, "accessKey"), ""),
			dynamic.StringValue(dynamic.Get(config, "secretKey"), ""),
			dynamic.StringValue(dynamic.Get(config, "bucket"), ""),
			dynamic.BooleanValue(dynamic.Get(config, "useSSL"), false))
	})
}

type MinioSource struct {
	cli     *minio.Client
	bucket  string
	baseURL string
}

func NewMinioSource(endpoint string, accessKey string, secretKey string, bucket string, useSSL bool) (*MinioSource, error) {
	cli, err := minio.New(endpoint, accessKey, secretKey, useSSL)
	if err != nil {
		return nil, err
	}
	baseURL := "http://" + endpoint
	if useSSL {
		baseURL = "https://" + endpoint
	}
	return &MinioSource{cli: cli, bucket: bucket, baseURL: baseURL}, nil
}

func (S *MinioSource) Get(key string) ([]byte, error) {
	object, err := S.cli.GetObject(S.bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()
	return ioutil.ReadAll(object)
}

func (S *MinioSource) GetURL(key string) string {
	return fmt.Sprintf("%s/%s/%s", S.baseURL, S.bucket, key)
}

func (S *MinioSource) GetSignURL(key string, expires time.Duration) (string, error) {
	u, err := S.cli.PresignedGetObject(S.bucket, key, expires, nil)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (S *MinioSource) Put(key string, data []byte, header map[string]string) error {

	opt := minio.PutObjectOptions{ContentType: "application/octet-stream", UserMetadata: map[string]string{}}

	if header != nil {
		for key, value := range header {
			if key == "Content-Type" {
				opt.ContentType = value
			} else if key == "Content-Encoding" {
				opt.ContentEncoding = value
			} else if key == "Content-Disposition" {
				opt.ContentDisposition = value
			} else {
				opt.UserMetadata[key] = value
			}
		}
	}

	_, err := S.cli.PutObject(S.bucket, key, bytes.NewReader(data), int64(len(data)), opt)

	if err != nil {
		return err
	}

	return nil
}

func (S *MinioSource) PutSignURL(key string, expires time.Duration) (string, error) {
	u, err := S.cli.PresignedPutObject(S.bucket, key, expires)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (S *MinioSource) PostSignURL(key string, expires time.Duration) (string, map[string]string, error) {

	p := minio.NewPostPolicy()
	p.SetExpires(time.Now().UTC().Add(expires))
	p.SetBucket(S.bucket)
	p.SetKey(key)

	u, data, err := S.cli.PresignedPostPolicy(p)

	if err != nil {
		return "", nil, err
	}

	return u.String(), data, nil
}

func (S *MinioSource) Del(key string) error {
	err := S.cli.RemoveObject(S.bucket, key)
	if err != nil {
		return err
	}
	return nil
}

func (S *MinioSource) Has(key string) error {
	_, err := S.cli.GetObjectACL(S.bucket, key)
	if err != nil {
		return err
	}
	return nil
}
