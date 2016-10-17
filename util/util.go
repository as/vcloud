package util

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"reflect"

)

// NextContext returns the offset of the next context
// located in the slice of strings, if it exists.
func NextContext(c []string) (offset int, exists bool) {
	if len(c) == 0 {
		return 0, false
	}

	var allow bool

	for i, v := range c {
		if allow || v[0] == '-' {
			if !allow {
				allow = true
			} else {
				allow = false
			}
		} else {
			offset = i
			exists = true
			break
		}
	}

	if exists {
		return offset, true
	}

	return 0, false
}

// TmpParams holds flag.Values provided by PrepareParams.
var TmpParams []string

// PrepareParams fills the TmpParams slice
// with --GNU-style argument specifier strings
func PrepareParams(fn *flag.Flag) {
	TmpParams = append(TmpParams, "--"+fn.Name)
}

// PrepareArgs converts each dashed argument
// specifier string to --GNU-style.
func PrepareArgs(args []string) {
	for i, v := range args {
		if len(v) < 2 {
			continue
		}

		if v[0] == '-' && v[1] != '-' {
			args[i] = "-" + args[i]
		}
	}
}



// FieldPrintNames prints struct field names by index, seperated by
// the string 'sep'
func FieldPrintNames(out io.Writer, rec interface{}, sep string, args ...int) error {
	val := reflect.TypeOf(rec)
	if val.Kind() != reflect.Struct {
		return errors.New("FieldPrintNames expects a struct as input")
	}

	for _, v := range args {
		fmt.Fprintf(out, "%v%s", val.Field(v).Name, sep)
	}
	fmt.Fprintf(out, "\r\n")

	return nil
}

// FieldNames prints struct fields by named string
func FieldNames(out io.Writer, rec interface{}, sep string, args ...string) error {
	val := reflect.TypeOf(rec)
	if val.Kind() != reflect.Struct {
		return errors.New("FieldNames expects a struct as input")
	}
	for _, v := range args {
		field, ok := val.FieldByName(v)
		if !ok{
			continue
		}
		fmt.Fprintf(out, "%v%s", field.Name, sep)
	}
	fmt.Fprintf(out, "\r\n")
	return nil
}

// FieldPrintValues prints struct field values by index, seperated by
// the string 'sep'.
func FieldValues(out io.Writer, rec interface{}, sep string, args ...string) error {
	val := reflect.ValueOf(rec)
	if val.Kind() != reflect.Struct {
		return errors.New("FieldPrintValues expects a struct as input")
	}
	for _, v := range args {
		field := val.FieldByName(v)
		fmt.Fprintf(out, "%v\t", field.Interface())
	}
	fmt.Fprintf(out, "\r\n")

	return nil
}

// FieldPrintValues prints struct field values by index, seperated by
// the string 'sep'.
func FieldPrintValues(out io.Writer, rec interface{}, sep string, args ...int) error {
	val := reflect.ValueOf(rec)
	if val.Kind() != reflect.Struct {
		return errors.New("FieldPrintValues expects a struct as input")
	}

	for _, v := range args {
		fmt.Fprintf(out, "%v\t", val.Field(v).Interface())
	}
	fmt.Fprintf(out, "\r\n")

	return nil
}

// SliceStructFields takes a struct and converts
// it to a slice of field names. Each slice index corresponds
// to the index of the field arranged in the struct.
func SliceStructFields(rec interface{}) []string {
	val := reflect.TypeOf(rec)
	if val.Kind() != reflect.Struct {
		panic("FieldToNo expects a struct as input")
	}

	length := val.NumField()
	fields := make([]string, length)
	for i := 0; i < length; i++ {
		fields[i] = val.Field(i).Name
	}

	return fields
}

func APItoC9(s string) string {
	b := []byte(s)
	if len(b) == 0 {
		return ""
	}

	if b[0] > 0x60 && b[0] < 0x7B {
		b[0] -= 0x20
	}

	return string(b)
}

func C9toAPI(s string) string {
	b := []byte(s)
	if len(b) == 0 {
		return ""
	}

	if b[0] > 0x40 && b[0] < 0x5B {
		b[0] += 0x20
	}

	return string(b)
}
