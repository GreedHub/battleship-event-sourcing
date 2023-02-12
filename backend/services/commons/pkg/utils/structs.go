package utils

import "reflect"

func StructName(s interface{}) string{
	return reflect.TypeOf(s).Name()
}

func EventName(e interface{}) string{
	return StructName(e)
}