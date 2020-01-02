package geoip

import (
	"fmt"
	"net"
	"reflect"
	"sync"

	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/dynamic"
	less "github.com/hailongz/golang/serverless/app"
	"github.com/oschwald/geoip2-golang"
)

var geoip *geoip2.Reader = nil
var lock sync.Mutex = sync.Mutex{}

func pushGeoObject(ctx *duktape.Context, v reflect.Value) {

	switch v.Kind() {
	case reflect.String:
		ctx.PushString(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		ctx.PushInt(int(v.Int()))
	case reflect.Int64:
		{
			i64 := v.Int()
			f64 := float64(i64)
			if i64 == int64(f64) {
				ctx.PushNumber(f64)
			} else {
				ctx.PushString(fmt.Sprintf("%d", i64))
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		ctx.PushUint(uint(v.Uint()))
	case reflect.Uint64:
		{
			i64 := v.Uint()
			f64 := float64(i64)
			if i64 == uint64(f64) {
				ctx.PushNumber(f64)
			} else {
				ctx.PushString(fmt.Sprintf("%u", i64))
			}
		}
	case reflect.Float32, reflect.Float64:
		ctx.PushNumber(v.Float())
	case reflect.Bool:
		ctx.PushBoolean(v.Bool())
	case reflect.Ptr:
		if v.IsNil() {
			ctx.PushUndefined()
		} else {
			pushGeoObject(ctx, v.Elem())
		}
	case reflect.Map:
		if v.IsNil() {
			ctx.PushUndefined()
		} else {
			ctx.PushObject()
			for _, key := range v.MapKeys() {
				pushGeoObject(ctx, key)
				pushGeoObject(ctx, v.MapIndex(key))
				ctx.PutProp(-3)
			}
		}
	case reflect.Slice:
		if v.IsNil() {
			ctx.PushUndefined()
		} else {
			ctx.PushArray()
			for i := 0; i < v.Len(); i++ {
				pushGeoObject(ctx, v.Index(i))
				ctx.PutPropIndex(-2, uint(i))
			}
		}
	case reflect.Struct:
		{
			ctx.PushObject()
			count := v.NumField()
			tp := v.Type()

			for i := 0; i < count; i++ {

				tf := tp.Field(i)
				fd := v.Field(i)

				name := tf.Tag.Get("maxminddb")

				if name == "-" {
					continue
				}

				if name == "" {
					name = tf.Name
				}

				ctx.PushString(name)
				pushGeoObject(ctx, fd)
				ctx.PutProp(-3)
			}
		}
	default:
		ctx.PushUndefined()
	}
}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		config := dynamic.Get(app.GetConfig(), "geoip")

		getGeoip := func() (*geoip2.Reader, error) {
			if geoip != nil {
				return geoip, nil
			}
			var err error = nil
			geoip, err = geoip2.Open(dynamic.StringValue(dynamic.Get(config, "db"), "./GeoLite2-City.mmdb"))

			if err != nil {
				return nil, err
			}

			return geoip, nil
		}

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				ip := net.ParseIP(ctx.ToString(-n))

				lock.Lock()
				defer lock.Unlock()

				geo, err := getGeoip()

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[Geoip] %s", err.Error())
					return duktape.ErrRetError
				}

				city, err := geo.City(ip)

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[Geoip] %s", err.Error())
					return duktape.ErrRetError
				}

				pushGeoObject(ctx, reflect.ValueOf(city))

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "city")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				ip := net.ParseIP(ctx.ToString(-n))

				lock.Lock()
				defer lock.Unlock()

				geo, err := getGeoip()

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[Geoip] %s", err.Error())
					return duktape.ErrRetError
				}

				country, err := geo.Country(ip)

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[Geoip] %s", err.Error())
					return duktape.ErrRetError
				}

				pushGeoObject(ctx, reflect.ValueOf(country))

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "country")

		ctx.PutGlobalString("geoip")
	})
}
