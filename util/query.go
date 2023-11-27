package util

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type FilterQuery struct {
	Conditions  []Condition
	Limit, Page int64
	Sort        []Sorting
}

type Sorting struct {
	Field, Type string
}

type Condition struct {
	Field, Operation string
	Value            interface{}
}

func (filter FilterQuery) BuildQuery() (query, sorting string, arguments []interface{}) {

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
		} else if i.Operation == "=" || i.Operation == ">" || i.Operation == "<" || i.Operation == ">=" || i.Operation == "<=" || i.Operation == "!=" || i.Operation == "like" || i.Operation == "<>" {
			conditions = append(conditions, fmt.Sprintf("%s ?", i.Operation))
		} else if i.Operation == "in" && (typeOfValue == reflect.Slice || typeOfValue == reflect.Array) {
			params := []string{}
			for i := 0; i < valueOfValue.Len(); i++ {
				params = append(params, "?")
			}

			conditions = append(conditions, fmt.Sprintf("in (%s)", strings.Join(params, ",")))
		}

		if typeOfValue == reflect.String {
			arguments = append(arguments, i.Value)
		} else if typeOfValue == reflect.Int || typeOfValue == reflect.Int8 || typeOfValue == reflect.Int16 || typeOfValue == reflect.Int32 || typeOfValue == reflect.Int64 {
			arguments = append(arguments, i.Value)
		} else if typeOfValue == reflect.Uint || typeOfValue == reflect.Uint8 || typeOfValue == reflect.Uint16 || typeOfValue == reflect.Uint32 || typeOfValue == reflect.Uint64 {
			arguments = append(arguments, i.Value)
		} else if typeOfValue == reflect.Float32 || typeOfValue == reflect.Float64 {
			arguments = append(arguments, i.Value)
		} else if typeOfValue == reflect.Slice || typeOfValue == reflect.Array {
			for i := 0; i < valueOfValue.Len(); i++ {
				arguments = append(arguments, valueOfValue.Index(i).Interface())
			}
		} else {
			if reflect.TypeOf(i.Value) == reflect.TypeOf(time.Time{}) {
				arguments = append(arguments, i.Value)
			}
		}
	}

	sorts := []string{}
	for _, i := range filter.Sort {
		if i.Field != "" && i.Type != "" {
			sorts = append(sorts, fmt.Sprintf("%s %s", i.Field, i.Type))
		}
	}

	return strings.Join(conditions, " "), strings.Join(sorts, ","), arguments
}

func (filter FilterQuery) GetOffset() int64 {
	return (filter.Page - 1) * filter.Limit
}
