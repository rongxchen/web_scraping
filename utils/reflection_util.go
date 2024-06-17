package utils

import (
	"reflect"
	"strings"
)

func GetFieldValue(obj interface{}, field string) interface{} {
	value := reflect.ValueOf(obj)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value.FieldByName(field).Interface()
}

func GetGormColumnTag(obj interface{}, fieldName string) string {
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, _ := t.FieldByName(fieldName)
	tag := field.Tag.Get("gorm")

	if tag == "" {
		return fieldName
	}

	col := strings.Split(tag, ":")
	if len(col) == 2 {
		return col[1]
	}
	return fieldName
}
