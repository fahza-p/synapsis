package utils

import (
	"reflect"
	"strings"
)

func IsMapStringInterface(obj interface{}) bool {
	reflectType := reflect.TypeOf(obj)
	m := make(map[string]interface{})
	return reflectType == reflect.TypeOf(m) || reflectType == reflect.TypeOf(&m)
}

func IsStructOrPointerOf(obj interface{}) bool {
	return IsStruct(obj) || IsPointerOfStruct(obj)
}

func IsStruct(obj interface{}) bool {
	reflectType := reflect.TypeOf(obj)
	return reflectType.Kind() == reflect.Struct
}

func IsPointerOfStruct(obj interface{}) bool {
	reflectType := reflect.TypeOf(obj)

	if reflectType.Kind() != reflect.Ptr {
		return false
	}

	if reflectType.Elem().Kind() != reflect.Struct {
		return false
	}

	return true
}

func GetAllJsonTagName(obj interface{}) (fields []string) {
	reflectType := reflect.TypeOf(obj)

	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}

	if reflectType.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < reflectType.NumField(); i++ {
		ft := reflectType.Field(i).Tag.Get("json")
		tagName := ft
		if i := strings.IndexByte(ft, ','); i >= 0 {
			tagName = ft[:i]
		}

		fields = append(fields, tagName)
	}

	return
}
