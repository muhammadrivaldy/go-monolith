package util

import (
	"fmt"
	"reflect"
	"strings"
)

type FilterQuery struct {
	Conditions  []Condition
	Limit, Page int
}

type Condition struct {
	Field     string
	Operation string
	Value     interface{}
}

func (filter FilterQuery) BuildQuery() (query string, arguments []interface{}) {

	conditions := []string{}

	for _, i := range filter.Conditions {

		if i.Field != "" {
			conditions = append(conditions, i.Field)
		}

		var typeOfValue reflect.Kind
		var valueOfValue reflect.Value
		if i.Value != nil {
			typeOfValue = reflect.TypeOf(i.Value).Kind()
			valueOfValue = reflect.ValueOf(i.Value)
		}

		if i.Operation == "and" || i.Operation == "or" || i.Operation == "(" || i.Operation == ")" {
			conditions = append(conditions, i.Operation)
		} else if i.Operation == "=" || i.Operation == ">" || i.Operation == "<" || i.Operation == ">=" || i.Operation == "<=" || i.Operation == "!=" || i.Operation == "like" {
			conditions = append(conditions, fmt.Sprintf("%s ?", i.Operation))
		} else if i.Operation == "in" && typeOfValue == reflect.Array {
			params := []string{}
			for i := 0; i < valueOfValue.Len(); i++ {
				params = append(params, "?")
			}

			conditions = append(conditions, fmt.Sprintf("(%s)", strings.Join(params, ",")))
		}

		if typeOfValue == reflect.String {
			arguments = append(arguments, i.Value)
		} else if typeOfValue == reflect.Int || typeOfValue == reflect.Int8 || typeOfValue == reflect.Int16 || typeOfValue == reflect.Int32 || typeOfValue == reflect.Int64 {
			arguments = append(arguments, i.Value)
		} else if typeOfValue == reflect.Uint || typeOfValue == reflect.Uint8 || typeOfValue == reflect.Uint16 || typeOfValue == reflect.Uint32 || typeOfValue == reflect.Uint64 {
			arguments = append(arguments, i.Value)
		} else if typeOfValue == reflect.Float32 || typeOfValue == reflect.Float64 {
			arguments = append(arguments, i.Value)
		} else if typeOfValue == reflect.Array {
			for i := 0; i < valueOfValue.Len(); i++ {
				arguments = append(arguments, valueOfValue.Index(i))
			}
		}
	}

	return strings.Join(conditions, " "), arguments
}

func (filter FilterQuery) GetOffset() int {
	return (filter.Page - 1) * filter.Limit
}
