package utils

import (
	"reflect"
	"strings"
)

func Slash(text string) (string, string) {
	res := strings.SplitN(text, "/", 2)
	if len(res) == 1 {
		return res[0], ""
	}
	return res[0], res[1]
}

func HasElem(arr interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(arr)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			// panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}
