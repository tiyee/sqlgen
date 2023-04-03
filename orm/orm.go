package orm

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var rdb *sql.DB
var wdb *sql.DB

func init() {
	fmt.Println("initialize mysql...")
	// rdb = <db>
	// wdb = <db>
}
func CopyPoint[T ITable](m T) T {
	vt := reflect.TypeOf(m).Elem()
	newoby := reflect.New(vt)
	newoby.Elem().Set(reflect.ValueOf(m).Elem())
	return newoby.Interface().(T)
}

func Binds[T ITable](m T) []interface{} {
	xv := reflect.ValueOf(m)
	n := xv.Elem().NumField()
	values := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		ptr := xv.Elem().Field(i).Addr().Interface()
		values = append(values, ptr)
	}
	return values
}

type ITable interface {
	TableName() string
	Pk() string
}
type ORM[T ITable] struct {
	Record T
}

func New[T ITable](data T) *ORM[T] {
	return &ORM[T]{
		Record: data,
	}
}
func (c *ORM[T]) TableName() string {
	return c.Record.TableName()
}
func (c *ORM[T]) Pk() string {
	return c.Record.Pk()
}
func (c *ORM[T]) Fields() string {
	xt := reflect.TypeOf(c.Record)
	n := xt.Elem().NumField()
	fields := make([]string, 0, n)
	for i := 0; i < n; i++ {
		if field, ok := xt.Elem().Field(i).Tag.Lookup("json"); ok {
			if field != "" {
				fields = append(fields, "`"+field+"`")
			}
		}
	}
	return strings.Join(fields, ",")
}

func (c *ORM[T]) Row(where string, values []interface{}) (T, error) {
	db := rdb
	if err := db.Ping(); err != nil {
		return c.Record, err
	}
	conditions := []string{"select", c.Fields(), "from", c.TableName()}
	if len(where) > 3 {
		conditions = append(conditions, "where", where)
	}
	if err := db.QueryRow(strings.Join(conditions, " "), values...).Scan(Binds(c.Record)...); err == nil {
		return c.Record, nil
	} else {
		return c.Record, err
	}
}
func (c *ORM[T]) Count(where string, values []interface{}) (int64, error) {
	db := rdb
	conditions := []string{"select count(*) as n  from", c.TableName()}
	if len(where) > 3 {
		conditions = append(conditions, "where", where)
	}
	var n int64
	if err := db.QueryRow(strings.Join(conditions, " "), values...).Scan(&n); err == nil {
		return n, nil
	} else {
		return 0, err
	}
}

func (c *ORM[T]) Rows(where string, values []interface{}) ([]T, error) {
	db := rdb
	conditions := []string{
		"select",
		c.Fields(),
		"from", c.TableName()}
	if len(where) > 3 {
		conditions = append(conditions, "where", where)
	}

	var myErr error
	if rows, err := db.Query(strings.Join(conditions, " "), values...); err == nil {
		results := make([]T, 0)
		for rows.Next() {
			row := CopyPoint(c.Record)
			if err := rows.Scan(Binds(row)...); err == nil {
				results = append(results, row)
			} else {
				fmt.Println(err.Error())
			}
		}
		rows.Close()
		return results, nil
	} else {
		if err == sql.ErrNoRows {
			return []T{}, nil
		}
		myErr = err
	}
	return nil, myErr
}
func (c *ORM[T]) Save() (int64, error) {
	xv := reflect.ValueOf(c.Record)
	n := xv.Elem().NumField()
	values := make([]interface{}, 0, n)
	placements := make([]string, 0, n)
	for i := 0; i < n; i++ {
		ptr := xv.Elem().Field(i).Interface()
		values = append(values, ptr)
		placements = append(placements, "?")
	}
	db := wdb
	sqlFmt := fmt.Sprintf(" insert into %s (%s) values (%s)", c.TableName(), c.Fields(), strings.Join(placements, ","))

	if result, err := db.Exec(sqlFmt, values...); err == nil {
		return result.LastInsertId()
	} else {
		return 0, err
	}

}
func (c *ORM[T]) Saves(rows []T) (int64, error) {
	xv := reflect.ValueOf(c.Record)
	n := xv.Elem().NumField()

	placements := make([]string, 0, n)
	for i := 0; i < n; i++ {
		placements = append(placements, "?")
	}
	values := make([][]interface{}, 0, len(rows))
	for idx := range rows {
		v := reflect.ValueOf(rows[idx])
		value := make([]interface{}, 0, n)
		for i := 0; i < n; i++ {
			ptr := v.Elem().Field(i).Interface()
			value = append(value, ptr)

		}
		values = append(values, value)

	}
	db := wdb
	sqlFmt := fmt.Sprintf(" insert into %s (%s) values (%s)", c.TableName(), c.Fields(), strings.Join(placements, ","))
	if stmt, err := db.Prepare(sqlFmt); err == nil {
		defer stmt.Close()
		for i := range values {
			if _, err := stmt.Exec(values[i]...); err != nil {
				return 0, err
			}

		}
	} else {
		return 0, err
	}
	return 1, nil

}

func (c *ORM[T]) Limit(where string, values []interface{}, offset, size int64, order string) ([]T, error) {
	db := rdb
	conditions := []string{
		"select",
		c.Fields(),
		"from", c.TableName()}
	if len(where) > 3 {
		conditions = append(conditions, "where", where)
	}
	if len(order) > 5 {
		conditions = append(conditions, "order by ", order)
	}
	if size < 1 {
		size = 20
	}
	if offset > 0 {
		conditions = append(conditions, "limit", fmt.Sprintf(" %d , %d ", offset, size))
	} else {
		conditions = append(conditions, "limit", fmt.Sprintf("  %d ", size))
	}

	var myErr error
	if rows, err := db.Query(strings.Join(conditions, " "), values...); err == nil {

		results := make([]T, 0, size)
		for rows.Next() {
			row := CopyPoint(c.Record)
			if err := rows.Scan(Binds(row)...); err == nil {
				results = append(results, row)
			}
		}
		rows.Close()
		return results, nil
	} else {
		myErr = err
	}
	return nil, myErr
}
func (c *ORM[T]) UpdateByPk(data map[string]interface{}, pk int64) (int64, error) {
	condition := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data)+1)
	for k, v := range data {
		condition = append(condition, "`"+k+"`=?")
		values = append(values, v)
	}
	values = append(values, pk)
	db := wdb
	sqls := []string{"update", c.TableName(), "set", strings.Join(condition, ","), "where " + c.Pk() + "=?"}
	if ret, err := db.Exec(strings.Join(sqls, " "), values...); err == nil {
		return ret.RowsAffected()
	} else {
		return 0, nil
	}
}
func (c *ORM[T]) Update(pk int64) (int64, error) {
	xv := reflect.ValueOf(c.Record)
	xt := reflect.TypeOf(c.Record)
	n := xv.Elem().NumField()
	values := make([]interface{}, 0)
	conditions := make([]string, 0)
	for i := 0; i < n; i++ {
		if field, ok := xt.Elem().Field(i).Tag.Lookup("json"); ok {
			if field == "" {
				continue
			}
			if field == c.Record.Pk() {
				continue
			}
			conditions = append(conditions, "`"+field+"`=?")
		}

		ptr := xv.Elem().Field(i).Interface()
		values = append(values, ptr)

	}
	if len(conditions) == 0 {
		return 0, errors.New("empty conditions")
	}
	values = append(values, pk)
	db := wdb
	sqls := []string{"update", c.Record.TableName(), "set", strings.Join(conditions, ","), "where " + c.Record.Pk() + "=?"}
	if _, err := db.Exec(strings.Join(sqls, " "), values...); err == nil {
		return 1, nil
	} else {
		return 0, err
	}
}
