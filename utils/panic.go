package utils

import "errors"

func PanicIfError(err error, message string) {
	if err != nil {
		e := errors.New(message)
		panic(e)
	}
}
