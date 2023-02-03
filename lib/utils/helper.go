package utils

import (
	"errors"
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

func ToSlug(s string, sep string) string {
	if sep != "_" && sep != "-" {
		sep = "_"
	}
	slug := strings.ToLower(s)
	slug = strings.ReplaceAll(slug, " ", sep)
	return slug
}

func FindFieldByTag(obj interface{}, tag, key string) (string, error) {
	reflectType := reflect.TypeOf(obj)
	switch reflectType.Kind() {
	case reflect.Ptr:
		reflectType = reflectType.Elem()
		fallthrough
	case reflect.Struct:
		for i := 0; i < reflectType.NumField(); i++ {
			field := reflectType.Field(i)
			if ft := field.Tag.Get(tag); ft == key || strings.HasPrefix(ft, key+",") {
				return field.Name, nil
			}
		}
		return "", errors.New("field not found")
	case reflect.Map:
		return key, nil
	default:
		return "", errors.New("unsupported type")
	}
}

func IsExistFieldByTag(obj interface{}, tag, key string) bool {
	if _, err := FindFieldByTag(obj, tag, key); err != nil {
		return false
	}
	return true
}
