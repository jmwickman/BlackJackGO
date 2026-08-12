package main

import "reflect"

func IsNilOrZero(a any) bool {
	return reflect.ValueOf(a).IsNil() || reflect.ValueOf(a).IsZero()
}
