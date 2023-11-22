package structocaster

import (
	"reflect"
	"testing"
)

type TestStructA struct {
	Field1 string
	Field2 int
}

type TestStructB struct {
	Field1 string `out:"Field1"`
	Field2 int    `out:"Field2"`
}

type I interface{}

type A struct {
	Greeting string
	Message  string
	Pi       float64
}

type B struct {
	Struct    A
	Ptr       *A
	Answer    int
	Map       map[string]string
	StructMap map[string]interface{}
	Slice     []string
}

func TestCastVBasic(t *testing.T) {
	src := TestStructA{Field1: "test", Field2: 123}
	dest := TestStructB{}

	srcV := reflect.ValueOf(&src)
	destV := reflect.ValueOf(&dest)

	castV(srcV.Interface(), destV)

	if dest.Field1 != src.Field1 || dest.Field2 != src.Field2 {
		t.Errorf("castV did not copy struct fields correctly")
	}
}

func TestCastVNilPointerToStruct(t *testing.T) {
	var original *B
	var dest B
	destV := reflect.ValueOf(&dest)

	castV(original, destV)

	if !reflect.DeepEqual(dest, B{}) {
		t.Errorf("Expected empty struct, got: %+v", dest)
	}
}

func TestCastVNilPointerToInterface(t *testing.T) {
	var original *I
	var dest I
	destV := reflect.ValueOf(&dest)

	castV(original, destV)

	if dest != nil {
		t.Errorf("Expected nil, got: %+v", dest)
	}
}

func TestCastVEmptyStruct(t *testing.T) {
	original := B{}
	var dest B
	destV := reflect.ValueOf(&dest)

	castV(&original, destV)

	if !reflect.DeepEqual(original, dest) {
		t.Errorf("Expected %+v, got: %+v", original, dest)
	}
}

func TestCastVStructWithNoElements(t *testing.T) {
	type E struct{}
	original := E{}
	var dest E
	destV := reflect.ValueOf(&dest)

	castV(&original, destV)

	if !reflect.DeepEqual(original, dest) {
		t.Errorf("Expected %+v, got: %+v", original, dest)
	}
}

func TestCastVEmptyStructPointer(t *testing.T) {
	original := &B{}
	var dest B
	destV := reflect.ValueOf(&dest)

	castV(original, destV)

	if !reflect.DeepEqual(*original, dest) {
		t.Errorf("Expected %+v, got: %+v", *original, dest)
	}
}
