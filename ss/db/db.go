package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/dynamic"
	less "github.com/hailongz/golang/serverless/app"
)

func getConn(app *less.App, name string) (*sql.DB, error) {

	v, err := app.GetSharedObject("db."+name, func() (less.SharedObject, error) {

		config := dynamic.GetWithKeys(app.GetConfig(), []string{"db", name})

		drive := dynamic.StringValue(dynamic.Get(config, "name"), "mysql")
		url := dynamic.StringValue(dynamic.Get(config, "url"), "root:123456@tcp(127.0.0.1:3306)/kk?charset=utf8mb4")

		conn, err := sql.Open(drive, url)

		if err != nil {
			return nil, err
		}

		err = conn.Ping()

		if err != nil {
			return nil, err
		}

		conn.SetMaxIdleConns(int(dynamic.IntValue(dynamic.Get(config, "maxIdleConns"), 1)))
		conn.SetMaxOpenConns(int(dynamic.IntValue(dynamic.Get(config, "maxOpenConns"), 6)))
		conn.SetConnMaxLifetime(time.Duration(dynamic.IntValue(dynamic.Get(config, "maxLifetime"), 6)) * time.Second)

		return conn, nil
	})

	if err != nil {
		return nil, err
	}

	return v.(*sql.DB), nil
}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		config := dynamic.Get(app.GetConfig(), "db")

		prefix := dynamic.StringValue(dynamic.Get(config, "prefix"), "")

		var errmsg string

		ctx.PushObject()

		ctx.PushString(prefix)
		ctx.PutPropString(-2, "prefix")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 && ctx.IsString(-n) && ctx.IsString(-n+1) {

				name := ctx.ToString(-n)
				sql := ctx.ToString(-n + 1)

				conn, err := getConn(app, name)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s\n", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				args := []interface{}{}

				for i := 2; i < n; i++ {
					args = append(args, less.Decode(ctx, -n+i))
				}

				rs, err := conn.Exec(sql, args...)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s\n", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				ctx.PushObject()

				id, _ := rs.LastInsertId()
				less.Encode(ctx, id)
				ctx.PutPropString(-2, "id")

				count, _ := rs.RowsAffected()
				less.Encode(ctx, count)
				ctx.PutPropString(-2, "count")

				return 1
			}

			return 0
		})
		ctx.PutPropString(-2, "exec")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 && ctx.IsString(-n) && ctx.IsString(-n+1) {

				name := ctx.ToString(-n)
				sql := ctx.ToString(-n + 1)

				conn, err := getConn(app, name)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				args := []interface{}{}

				for i := 2; i < n; i++ {
					args = append(args, less.Decode(ctx, -n+i))
				}

				rs, err := conn.Query(sql, args...)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				defer rs.Close()

				ctx.PushArray()

				var idx uint = 0

				var vs []interface{} = nil
				var names []string = nil
				var vsptr []interface{} = nil
				var n int = 0

				for rs.Next() {

					if vs == nil {
						names, err = rs.Columns()
						if err != nil {
							errmsg = err.Error()
							ctx.Pop()
							log.Printf("[%s] [DB] [ERROR] %s", trace, err.Error())
							ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
							return duktape.ErrRetError
						}
						n = len(names)
						vs = make([]interface{}, n)
						vsptr = make([]interface{}, n)
					}

					for i := 0; i < n; i++ {
						vs[i] = nil
						vsptr[i] = &vs[i]
					}

					err = rs.Scan(vsptr...)

					if err != nil {
						errmsg = err.Error()
						ctx.Pop()
						log.Printf("[%s] [DB] [ERROR] %s", trace, err.Error())
						ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
						return duktape.ErrRetError
					}

					row := map[string]interface{}{}

					for i := 0; i < n; i++ {

						v := vs[i]

						if v == nil {
							continue
						}

						{
							b, ok := v.([]byte)
							if ok {
								row[names[i]] = string(b)
								continue
							}
						}

						row[names[i]] = v

					}

					less.Encode(ctx, row)

					ctx.PutPropIndex(-2, idx)

					idx = idx + 1
				}

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "query")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 && ctx.IsString(-n) && ctx.IsString(-n+1) {

				name := ctx.ToString(-n)
				sql := ctx.ToString(-n + 1)

				conn, err := getConn(app, name)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				args := []interface{}{}

				for i := 2; i < n; i++ {
					args = append(args, less.Decode(ctx, -n+i))
				}

				rs, err := conn.Query(sql, args...)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				defer rs.Close()

				ctx.PushArray()

				var idx uint = 0

				var vs []interface{} = nil
				var names []string = nil
				var vsptr []interface{} = nil
				var n int = 0

				names, err = rs.Columns()

				if err != nil {
					errmsg = err.Error()
					ctx.Pop()
					log.Printf("[%s] [DB] [ERROR] %s", trace, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				n = len(names)
				vs = make([]interface{}, n)
				vsptr = make([]interface{}, n)

				{
					ctx.PushArray()
					for i := 0; i < n; i++ {
						ctx.PushString(names[i])
						ctx.PutPropIndex(-2, uint(i))
					}
					ctx.PutPropIndex(-2, idx)
					idx = idx + 1
				}

				for rs.Next() {

					for i := 0; i < n; i++ {
						vs[i] = nil
						vsptr[i] = &vs[i]
					}

					err = rs.Scan(vsptr...)

					if err != nil {
						errmsg = err.Error()
						ctx.Pop()
						log.Printf("[%s] [DB] [ERROR] %s", trace, err.Error())
						ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
						return duktape.ErrRetError
					}

					ctx.PushArray()

					for i := 0; i < n; i++ {

						v := vs[i]

						if v != nil {
							b, ok := v.([]byte)
							if ok {
								ctx.PushString(string(b))
							} else {
								less.Encode(ctx, v)
							}
						} else {
							ctx.PushNull()
						}

						ctx.PutPropIndex(-2, uint(i))

					}

					ctx.PutPropIndex(-2, idx)

					idx = idx + 1
				}

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "table")

		ctx.PushGoFunction(func() int {

			ctx.PushString(errmsg)

			return 1
		})

		ctx.PutPropString(-2, "getErrmsg")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 4 && ctx.IsString(-n) && ctx.IsString(-n+1) && ctx.IsArray(-n+2) && ctx.IsString(-n+3) && ctx.IsString(-n+4) {

				dst_name := ctx.ToString(-n)
				dst_table := ctx.ToString(-n + 1)
				dst_uniqueKeys := []string{}

				{
					ctx.Enum(-n+2, duktape.DUK_ENUM_ARRAY_INDICES_ONLY)
					for ctx.Next(-1, true) {
						dst_uniqueKeys = append(dst_uniqueKeys, ctx.ToString(-1))
						ctx.Pop2()
					}
					ctx.Pop()
				}

				name := ctx.ToString(-n + 3)
				sql := ctx.ToString(-n + 4)

				conn, err := getConn(app, name)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				dst_conn, err := getConn(app, dst_name)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				args := []interface{}{}

				for i := 5; i < n; i++ {
					args = append(args, less.Decode(ctx, -n+i))
				}

				rs, err := conn.Query(sql, args...)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				defer rs.Close()

				err = db.Copy(rs, dst_conn, dst_table)

				if err != nil {
					errmsg = err.Error()
					log.Printf("[%s] [DB] %s [ERROR] %s", trace, sql, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				return 0
			}

			return 0
		})

		ctx.PutPropString(-2, "copy")

		ctx.PutGlobalString("db")
	})

}
