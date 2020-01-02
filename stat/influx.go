package stat

import (
	"time"

	"github.com/hailongz/golang/dynamic"
	influx "github.com/influxdata/influxdb1-client/v2"
)

func init() {
	AddOpenlib("influx", func(config interface{}) (Client, error) {
		return NewInfluxClient(dynamic.StringValue(dynamic.Get(config, "addr"), ""),
			dynamic.StringValue(dynamic.Get(config, "user"), ""),
			dynamic.StringValue(dynamic.Get(config, "password"), ""),
			dynamic.StringValue(dynamic.Get(config, "db"), ""))
	})
}

type InfluxClient struct {
	c        influx.Client
	db       string
	addr     string
	user     string
	password string
}

func NewInfluxClient(addr string, user string, password string, db string) (*InfluxClient, error) {
	return &InfluxClient{c: nil, db: db, addr: addr, user: user, password: password}, nil
}

func (C *InfluxClient) Close() {
	if C.c == nil {
		return
	}
	C.c.Close()
}

func (C *InfluxClient) Write(name string, tags map[string]string, fields map[string]interface{}, tv time.Time) error {

	if C.c == nil {

		v, err := influx.NewHTTPClient(influx.HTTPConfig{
			Addr:     C.addr,
			Username: C.user,
			Password: C.password,
		})

		if err != nil {
			return err
		}

		C.c = v
	}

	p, err := influx.NewPoint(name, tags, fields, tv)

	if err != nil {
		return nil
	}

	bp, _ := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  C.db,
		Precision: "us",
	})

	bp.AddPoint(p)

	return C.c.Write(bp)

}
