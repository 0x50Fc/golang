package app

import (
	"fmt"
	"reflect"

	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/dynamic"
)

func Marshal(ctx *duktape.Context, v interface{}) {
	Encode(ctx, v)
}

func Encode(ctx *duktape.Context, v interface{}) {

	if v == nil {
		ctx.PushUndefined()
		return
	}

	{
		b, ok := v.([]byte)
		if ok {
			ctx.PushBytes(b)
			return
		}
	}

	vv := reflect.ValueOf(v)
	switch vv.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Int, reflect.Uint:
		ctx.PushInt(int(vv.Int()))
		break
	case reflect.Int64:
		{
			i64 := vv.Int()
			i32 := int32(i64)
			if int64(i32) == i64 {
				ctx.PushInt(int(i32))
			} else {
				ctx.PushString(fmt.Sprintf("%d", i64))
			}
		}
		break
	case reflect.Uint64:
		{
			i64 := vv.Uint()
			i32 := uint32(i64)
			if uint64(i32) == i64 {
				ctx.PushUint(uint(i32))
			} else {
				ctx.PushString(fmt.Sprintf("%d", i64))
			}
		}
		break
	case reflect.Bool:
		ctx.PushBoolean(vv.Bool())
		break
	case reflect.Float32, reflect.Float64:
		ctx.PushNumber(vv.Float())
		break
	case reflect.String:
		ctx.PushString(vv.String())
		break
	case reflect.Slice:
		ctx.PushArray()
		var i uint = 0
		dynamic.Each(v, func(key interface{}, value interface{}) bool {
			if value == nil {
				return true
			}
			Encode(ctx, value)
			ctx.PutPropIndex(-2, i)
			i = i + 1
			return true
		})
		break
	default:
		ctx.PushObject()
		dynamic.Each(v, func(key interface{}, value interface{}) bool {
			if value == nil {
				return true
			}
			ctx.PushString(dynamic.StringValue(key, ""))
			Encode(ctx, value)
			ctx.PutProp(-3)
			return true
		})
		break
	}
}

func Unmarshal(ctx *duktape.Context, idx int, object interface{}) {
	dynamic.SetValue(object, Decode(ctx, idx))
}

func Decode(ctx *duktape.Context, idx int) interface{} {

	switch ctx.GetType(idx) {
	case duktape.TypeString:
		return ctx.ToString(idx)
	case duktape.TypeNumber:
		return ctx.ToNumber(idx)
	case duktape.TypeBoolean:
		return ctx.ToBoolean(idx)
	case duktape.TypeBuffer:
		return ctx.ToBytes(idx)
	case duktape.TypeObject:
		if ctx.IsBufferData(idx) {
			return ctx.ToBytes(idx)
		} else if ctx.IsArray(idx) {

			object := []interface{}{}

			ctx.Enum(idx, duktape.DUK_ENUM_ARRAY_INDICES_ONLY)

			for ctx.Next(-1, true) {
				value := Decode(ctx, -1)
				if value != nil {
					object = append(object, value)
				}
				ctx.Pop2()
			}

			ctx.Pop()

			return object

		} else {

			object := map[interface{}]interface{}{}

			ctx.Enum(idx, duktape.DUK_ENUM_INCLUDE_SYMBOLS)

			for ctx.Next(-1, true) {
				key := Decode(ctx, -2)
				value := Decode(ctx, -1)
				if value != nil {
					object[key] = value
				}
				ctx.Pop2()
			}

			ctx.Pop()

			return object

		}
	}
	return nil
}
