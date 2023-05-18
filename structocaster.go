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
	if dest.Kind() == reflect.Ptr {
		dest = dest.Elem()
	}

	for i := 0; i < dest.NumField(); i++ {
		valueField := dest.Field(i)
		typeField := dest.Type().Field(i)

		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
				valueField = elm
			}
		}

		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()
		}

		if valueField.Kind() == reflect.Struct {
			castV(src, valueField)
		}

		SrcV := reflect.ValueOf(src).Elem()
		var FieldName string
		var Value reflect.Value
		if len(typeField.Tag.Get(tagName)) > 0 {
			FieldName = typeField.Tag.Get(tagName)

			if strings.Contains(FieldName, ".") {
				Value = SrcV
				words := strings.Split(FieldName, ".")
				for _, word := range words {
					Value = Value.FieldByName(word)
				}
			} else {
				Value = SrcV.FieldByName(FieldName)
			}
		} else {
			FieldName = typeField.Name
			Value = SrcV.FieldByName(FieldName)
		}

		if !Value.IsValid() || !Value.Type().AssignableTo(typeField.Type) {
			continue
		}

		valueField.Set(Value)
	}
}

func cast(SrcV interface{}, v interface{}) {
	castV(SrcV, reflect.ValueOf(v))
}
