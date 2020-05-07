package db

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/hailongz/golang/dynamic"
)

type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type IObject interface {
	GetTitle() string
	GetName() string
	GetId() int64
	SetId(id int64)
}

type Object struct {
	Id int64 `json:"id" title:"ID"`
}

type Field struct {
	F            reflect.StructField `json:"-"`
	V            reflect.Value       `json:"-"`
	Name         string              `json:"name"`
	IsObject     bool                `json:"-"`
	IsJSONObject bool                `json:"-"`
}

func (O *Object) GetTitle() string {
	return "数据对象"
}

func (O *Object) GetName() string {
	return "object"
}

func (O *Object) GetId() int64 {
	return O.Id
}

func (O *Object) SetId(id int64) {
	O.Id = id
}

func TableName(prefix string, object IObject) string {
	return prefix + object.GetName()
}

func Transaction(db *sql.DB, fn func(conn Database) error) error {

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	err = fn(tx)

	if err == nil {
		err = tx.Commit()
	}

	if err != nil {
		tx.Rollback()
	}

	return err
}

func getColumnTypeSQL(columnType *sql.ColumnType) string {
	n, _ := columnType.Length()
	name := columnType.DatabaseTypeName()
	if n == 0 && name == "VARCHAR" {
		n = 2048
	}
	if n > 0 {
		return fmt.Sprintf("%s(%d)", name, n)
	}
	return name
}

func Copy(rs *sql.Rows, db Database, table string, uniqueKeys ...string) error {

	columns, err := rs.Columns()

	if err != nil {
		return err
	}

	columnTypes, err := rs.ColumnTypes()

	if err != nil {
		return err
	}

	n := len(columns)

	{
		r, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE 0", table))

		if err != nil {

			s := bytes.NewBuffer(nil)
			s.WriteString("CREATE TABLE ")
			s.WriteString(table)
			s.WriteString(" (")

			for i := 0; i < n; i++ {
				name := columns[i]
				columnType := columnTypes[i]
				if i != 0 {
					s.WriteString(",")
				}
				s.WriteString("`")
				s.WriteString(name)
				s.WriteString("` ")
				s.WriteString(getColumnTypeSQL(columnType))
			}

			s.WriteString(")")

			log.Println("[SQL]", s.String())

			_, err = db.Exec(s.String())

			if err != nil {
				return err
			}

		} else {

			cols, err := r.Columns()

			if err != nil {
				r.Close()
				return err
			}

			colTypes, err := r.ColumnTypes()

			if err != nil {
				r.Close()
				return err
			}

			r.Close()

			columnSet := map[string]*sql.ColumnType{}

			for i := 0; i < n; i++ {
				columnSet[columns[i]] = columnTypes[i]
			}

			for i := 0; i < len(cols); i++ {
				name := cols[i]
				colType := colTypes[i]
				_, ok := columnSet[name]

				if !ok {

					s := fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s;", table, name, getColumnTypeSQL(colType))

					log.Println("[SQL]", s)

					_, err = db.Exec(s)

					if err != nil {
						return nil
					}
				}
			}

			return nil
		}
	}

	un := 0

	if uniqueKeys != nil {
		un = len(uniqueKeys)
	}

	var vs []interface{} = make([]interface{}, n+un)
	var vsptr []interface{} = make([]interface{}, n+un)
	var ui []int = make([]int, un)

	s := bytes.NewBuffer(nil)
	s.WriteString("INSERT INTO `")
	s.WriteString(table)
	s.WriteString("`(")

	for i := 0; i < n; i++ {
		name := columns[i]
		if i != 0 {
			s.WriteString(",")
		}
		s.WriteString("`")
		s.WriteString(name)
		s.WriteString("` ")
	}

	s.WriteString(") VALUSE(")

	for i := 0; i < n; i++ {
		if i != 0 {
			s.WriteString(",")
		}
		s.WriteString("?")
	}

	s.WriteString(")")

	var u *bytes.Buffer = nil

	if un > 0 {

		for k := 0; k < un; k++ {
			ui[k] = -1
		}

		u = bytes.NewBuffer(nil)

		u.WriteString("UPDATE `")
		u.WriteString(table)
		u.WriteString("` SET ")

		for i := 0; i < n; i++ {
			name := columns[i]
			if i != 0 {
				u.WriteString(",")
			}
			u.WriteString("`")
			u.WriteString(name)
			u.WriteString("`=?")

			for k := 0; k < un; k++ {
				if name == uniqueKeys[k] {
					ui[k] = i
				}
			}
		}

		u.WriteString(" WHERE ")

		for i := 0; i < un; i++ {
			uniqueKey := uniqueKeys[i]
			if i != 0 {
				u.WriteString(" AND ")
			}
			u.WriteString("`")
			u.WriteString(uniqueKey)
			u.WriteString("`=?")
		}

		for k := 0; k < un; k++ {
			if ui[k] == -1 {
				return errors.New("Not Found uniqueKey " + uniqueKeys[k])
			}
		}

	}

	for rs.Next() {

		for i := 0; i < n+un; i++ {
			vs[i] = nil
			vsptr[i] = &vs[i]
		}

		err = rs.Scan(vsptr[:n]...)

		if err != nil {
			return err
		}

		if un > 0 {

			for i := 0; i < un; i++ {
				vs[n+i] = vs[ui[i]]
			}

			log.Println(u.String(), vs)

			r, err := db.Exec(u.String(), vs...)

			if err != nil {
				return err
			}

			c, err := r.RowsAffected()

			if err != nil {
				return err
			}

			if c > 0 {
				continue
			}
		}

		log.Println(s.String(), vs[:n])

		_, err = db.Exec(s.String(), vs[:n]...)

		if err != nil {
			return err
		}
	}

	return nil
}

func Query(db Database, object IObject, prefix string, sql string, args ...interface{}) (*sql.Rows, error) {
	var tbname = prefix + object.GetName()
	return db.Query(fmt.Sprintf("SELECT * FROM `%s` %s", tbname, sql), args...)
}

func QueryWithKeys(db Database, object IObject, prefix string, keys map[string]bool, sql string, args ...interface{}) (*sql.Rows, error) {

	s := bytes.NewBuffer(nil)

	if keys == nil {
		s.WriteString("SELECT *")
	} else {

		s.WriteString("SELECT id")

		Each(object, func(field Field) bool {

			if keys[field.Name] {
				s.WriteString(fmt.Sprintf(",`%s`", field.Name))
			}

			return true
		})

	}

	s.WriteString(fmt.Sprintf(" FROM `%s%s` %s", prefix, object.GetName(), sql))

	return db.Query(s.String(), args...)
}

func Delete(db Database, object IObject, prefix string) (sql.Result, error) {
	var tbname = prefix + object.GetName()
	return db.Exec(fmt.Sprintf("DELETE FROM `%s` WHERE id=?", tbname), object.GetId())
}

func DeleteWithSQL(db Database, object IObject, prefix string, sql string, args ...interface{}) (sql.Result, error) {
	var tbname = prefix + object.GetName()
	return db.Exec(fmt.Sprintf("DELETE FROM `%s` %s", tbname, sql), args...)
}

func Get(db Database, object IObject, prefix string, sql string, args ...interface{}) (IObject, error) {

	var rs, err = Query(db, object, prefix, sql, args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	if rs.Next() {
		scaner := NewScaner(object)
		err = scaner.Scan(rs)
		if err != nil {
			return nil, err
		}
		return object, nil
	}

	return nil, nil
}

func Count(db Database, object IObject, prefix string, sql string, args ...interface{}) (int64, error) {

	var tbname = prefix + object.GetName()

	var rows, err = db.Query(fmt.Sprintf("SELECT COUNT(*) as c FROM `%s` %s", tbname, sql), args...)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	if rows.Next() {
		var v int64 = 0
		err = rows.Scan(&v)
		if err != nil {
			return 0, err
		}
		return v, nil
	}

	return 0, nil
}

func Update(db Database, object IObject, prefix string) (sql.Result, error) {
	return UpdateWithKeys(db, object, prefix, nil)
}

func UpdateWithKeys(db Database, object IObject, prefix string, keys map[string]bool) (sql.Result, error) {

	var tbname = prefix + object.GetName()
	var s bytes.Buffer
	var fs = []interface{}{}
	var n = 0

	s.WriteString(fmt.Sprintf("UPDATE `%s` SET ", tbname))

	Each(object, func(field Field) bool {

		if field.Name == "id" {
			return true
		}

		if keys == nil || keys[field.Name] {
			if n != 0 {
				s.WriteString(",")
			}
			if field.IsJSONObject && field.IsObject {
				s.WriteString(fmt.Sprintf(" `%s`=JSON_MERGE_PATCH(`%s`,?)", field.Name, field.Name))
				b, _ := json.Marshal(field.V.Interface())
				fs = append(fs, string(b))
			} else {
				s.WriteString(fmt.Sprintf(" `%s`=?", field.Name))
				if field.IsObject {
					b, _ := json.Marshal(field.V.Interface())
					fs = append(fs, string(b))
				} else {
					fs = append(fs, field.V.Interface())
				}

			}
			n += 1
		}

		return true
	})

	s.WriteString(" WHERE id=?")

	fs = append(fs, object.GetId())

	// log.Printf("[SQL] %s %s\n", s.String(), fs)

	return db.Exec(s.String(), fs...)
}

func Insert(db Database, object IObject, prefix string) (sql.Result, error) {
	var tbname = prefix + object.GetName()
	var s bytes.Buffer
	var w bytes.Buffer
	var fs = []interface{}{}
	var n = 0

	s.WriteString(fmt.Sprintf("INSERT INTO `%s`(", tbname))
	w.WriteString(" VALUES (")

	Each(object, func(field Field) bool {

		if field.Name == "id" && object.GetId() == 0 {
			return true
		}

		if n != 0 {
			s.WriteString(",")
			w.WriteString(",")
		}
		s.WriteString("`" + field.Name + "`")
		w.WriteString("?")
		if field.IsObject {
			b, _ := json.Marshal(field.V.Interface())
			fs = append(fs, string(b))
		} else {
			fs = append(fs, field.V.Interface())
		}

		n += 1

		return true
	})

	s.WriteString(")")

	w.WriteString(")")

	s.Write(w.Bytes())

	// log.Printf("%s %s\n", s.String(), fs)

	var rs, err = db.Exec(s.String(), fs...)

	if err == nil && object.GetId() == 0 {
		id, err := rs.LastInsertId()
		if err == nil {
			object.SetId(id)
		}
	}

	return rs, err
}

type booleanValue struct {
	v        reflect.Value
	intValue int
}

type jsonValue struct {
	v         reflect.Value
	byteValue interface{}
}

type Scaner struct {
	object        interface{}
	fields        []interface{}
	jsonObjects   []*jsonValue
	nilValue      interface{}
	booleanValues []*booleanValue
}

func NewScaner(object interface{}) *Scaner {
	return &Scaner{object, nil, nil, nil, nil}
}

func (o *Scaner) Scan(rows *sql.Rows) error {

	if o.fields == nil {

		var columns, err = rows.Columns()

		if err != nil {
			return err
		}

		var fdc = len(columns)
		var mi = map[string]int{}

		for i := 0; i < fdc; i += 1 {
			mi[columns[i]] = i
		}

		o.booleanValues = []*booleanValue{}
		o.jsonObjects = []*jsonValue{}
		o.fields = make([]interface{}, fdc)

		for i := 0; i < fdc; i += 1 {
			o.fields[i] = &o.nilValue
		}

		Each(o.object, func(field Field) bool {

			idx, ok := mi[field.Name]

			if ok {
				if field.F.Type.Kind() == reflect.Bool {
					b := booleanValue{}
					b.v = field.V
					b.intValue = 0
					o.fields[idx] = &b.intValue
					o.booleanValues = append(o.booleanValues, &b)
				} else if field.IsObject {
					b := jsonValue{}
					b.v = field.V
					b.byteValue = nil
					o.fields[idx] = &b.byteValue
					o.jsonObjects = append(o.jsonObjects, &b)
				} else {
					o.fields[idx] = field.V.Addr().Interface()
				}
			}

			return true
		})

	}

	err := rows.Scan(o.fields...)

	if err != nil {
		return err
	}

	for _, fd := range o.jsonObjects {

		if fd.byteValue == nil {
			dynamic.SetReflectValue(fd.v, nil)
			continue
		}

		{
			b, ok := (fd.byteValue).([]byte)
			if ok {
				_ = json.Unmarshal(b, fd.v.Addr().Interface())
				continue
			}
		}

		{
			s, ok := (fd.byteValue).(string)
			if ok {
				_ = json.Unmarshal([]byte(s), fd.v.Addr().Interface())
				continue
			}
		}

		dynamic.SetReflectValue(fd.v, nil)
	}

	for _, fd := range o.booleanValues {
		dynamic.SetReflectValue(fd.v, fd.intValue != 0)
	}

	return nil
}
