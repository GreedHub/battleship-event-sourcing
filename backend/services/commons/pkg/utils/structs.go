package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func StructName(s interface{}) string{
    if t := reflect.TypeOf(s); t.Kind() == reflect.Ptr {
        return "*" + t.Elem().Name()
    } else {
        return t.Name()
    }
}

func EventName(e interface{}) string{
	return StructName(e)
}

func ValidateRequestParams(method string, requiredParams []string, body map[string]interface{}) error{
	for _, s := range requiredParams{
		if body[s] == nil {
			return errors.New(fmt.Sprintf("Error invoking %s, parameter '%s' missing", method, s))
		}
	}
	
	return nil
}