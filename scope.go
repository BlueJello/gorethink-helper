package main

import "reflect"

type Scope struct {
	Value         interface{}
	db            *DB
	indirectValue *reflect.Value
	fields        map[string]interface{}
}

func (db *DB) NewScope(value interface{}) *Scope {
	db.Value = value
	return &Scope{db: db, Value: value}
}

func (scope *Scope) IndirectValue() reflect.Value {
	if scope.indirectValue == nil {
		value := reflect.Indirect(reflect.ValueOf(scope.Value))
		scope.indirectValue = &value
	}
	return *scope.indirectValue
}

func (scope *Scope) TableName() string {
	if scope.Value == nil {
		//scope.Err(errors.New("can't get table name"))
		return ""
	}

	data := scope.IndirectValue()
	if data.Kind() == reflect.Slice {
		elem := data.Type().Elem()
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		data = reflect.New(elem).Elem()
	}

	if fm := data.MethodByName("TableName"); fm.IsValid() {
		if v := fm.Call([]reflect.Value{}); len(v) > 0 {
			if result, ok := v[0].Interface().(string); ok {
				return result
			}
		}
	}

	// str := ToSnake(data.Type().Name())

	// if scope.db == nil || !scope.db.parent.singularTable {
	// 	for index, reg := range pluralMapKeys {
	// 		if reg.MatchString(str) {
	// 			return reg.ReplaceAllString(str, pluralMapValues[index])
	// 		}
	// 	}
	// }

	return data.Type().Name()
}
func (s *Scope) createTable() *Scope {

	return s
}
