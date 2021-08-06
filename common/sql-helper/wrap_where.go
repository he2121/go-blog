package sql_helper

import (
	"reflect"

	"github.com/go-xorm/builder"
)

type Option struct {
	Offset  int
	Limit   int
	OrderBy *string
}

func WrapWhere(where interface{}) (conds []builder.Cond, err error) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	rt := reflect.TypeOf(where)
	rv := reflect.ValueOf(where)

	for i := 0; i < rt.NumField(); i++ {
		structField := rt.Field(i)
		fieldName := structField.Name
		tag := structField.Tag
		colName := tag.Get("db")
		operator := tag.Get("operator")
		value := rv.FieldByName(fieldName)
		if value.IsZero() {
			continue
		}
		value = reflect.Indirect(value)
		switch {
		case operator == "" || operator == "=":
			conds = append(conds, builder.Eq{colName: value.Interface()})
		case operator == "in":
			conds = append(conds, builder.In(colName, value.Interface()))
		case operator == "like":
			conds = append(conds, builder.Like{colName, value.Interface().(string)})
		case operator == ">=":
			conds = append(conds, builder.Gte{colName: value.Interface()})
		case operator == ">":
			conds = append(conds, builder.Gt{colName: value.Interface()})
		case operator == "<=":
			conds = append(conds, builder.Lte{colName: value.Interface()})
		case operator == "<":
			conds = append(conds, builder.Lte{colName: value.Interface()})
		case operator == "null":
			conds = append(conds, builder.IsNull{colName})
		}
	}
	return conds, err
}
