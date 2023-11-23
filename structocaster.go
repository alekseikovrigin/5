package structocaster

import (
	"reflect"
	"strings"
)

func castV(src interface{}, dest reflect.Value) {
	tagName := "out"

	dest = dereferenceValue(dest)
	if dest.Kind() != reflect.Struct || !dest.IsValid() {
		return
	}

	SrcV := dereferenceValue(reflect.ValueOf(src))
	if !SrcV.IsValid() {
		return
	}

	for i := 0; i < dest.NumField(); i++ {
		valueField := dest.Field(i)
		typeField := dest.Type().Field(i)

		if !valueField.IsValid() || !valueField.CanSet() {
			continue
		}

		if valueField.Kind() == reflect.Struct {
			castV(src, valueField)
			continue
		}

		Value := getValueFromSrc(SrcV, typeField, tagName)
		if !Value.IsValid() || !Value.Type().AssignableTo(valueField.Type()) {
			continue
		}

		valueField.Set(Value)
	}
}

func dereferenceValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		if val.IsNil() {
			return reflect.Value{}
		}
		val = val.Elem()
	}
	return val
}

func getValueFromSrc(src reflect.Value, field reflect.StructField, tagName string) reflect.Value {
	var fieldName string
	var value reflect.Value

	if len(field.Tag.Get(tagName)) > 0 {
		fieldName = field.Tag.Get(tagName)
		if strings.Contains(fieldName, ".") {
			value = src
			words := strings.Split(fieldName, ".")
			for _, word := range words {
				if !value.IsValid() {
					break
				}
				value = value.FieldByName(word)
			}
		} else {
			value = src.FieldByName(fieldName)
		}
	} else {
		fieldName = field.Name
		value = src.FieldByName(fieldName)
	}

	return value
}

func Cast(SrcV interface{}, v interface{}) {
	castV(SrcV, reflect.ValueOf(v))
}
