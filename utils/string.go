package utils

import "fmt"

func SliceInterfaceToString(i []interface{}) []string {
	s := make([]string, len(i))
	for i, v := range i {
		s[i] = fmt.Sprint(v)
	}
	return s
}
