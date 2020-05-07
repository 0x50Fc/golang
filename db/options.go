package db

import "github.com/hailongz/golang/dynamic"

func Merge(options ...interface{}) interface{} {
	v := map[string]interface{}{}
	if options != nil {
		for _, opt := range options {
			dynamic.Each(opt, func(key interface{}, value interface{}) bool {
				v[dynamic.StringValue(key, "")] = value
				return true
			})
		}
	}
	return v
}
