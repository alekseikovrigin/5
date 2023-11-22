package structocaster

import (
	"reflect"
	"strings"
)

func castV(src interface{}, dest reflect.Value) {
	tagName := "out"
	if dest.Kind() == reflect.Interface && !dest.IsNil() {
		elm := dest.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			dest = elm
		}
	}
	if dest.Kind() == reflect.Ptr && !dest.IsNil() {
		dest = dest.Elem()
	}

	if dest.Kind() != reflect.Struct || !dest.IsValid() {
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

		SrcV := reflect.ValueOf(src)
		if SrcV.Kind() != reflect.Ptr || SrcV.IsNil() {
			continue
		}
		SrcV = SrcV.Elem()
		if !SrcV.IsValid() {
			continue
		}

		var FieldName string
		var Value reflect.Value

		if len(typeField.Tag.Get(tagName)) > 0 {
			FieldName = typeField.Tag.Get(tagName)

			if strings.Contains(FieldName, ".") {
				Value = SrcV
				words := strings.Split(FieldName, ".")
				for _, word := range words {
					if !Value.IsValid() {
						break
					}
					Value = Value.FieldByName(word)
				}
			} else {
				Value = SrcV.FieldByName(FieldName)
			}
		} else {
			FieldName = typeField.Name
			Value = SrcV.FieldByName(FieldName)
		}

		if !Value.IsValid() || !Value.Type().AssignableTo(valueField.Type()) {
			continue
		}

		valueField.Set(Value)
	}
}

func Cast(SrcV interface{}, v interface{}) {
	castV(SrcV, reflect.ValueOf(v))
}
