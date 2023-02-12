package utils

import "reflect"

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