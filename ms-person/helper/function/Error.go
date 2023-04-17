package function

import (
	"br.com.charlesrodrigo/ms-person/helper/logger"
)

func IfErrorPanic(text string, err error) {
	if err != nil {
		logger.Panic(err)
	}
}

func IfErrorFatal(text string, err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
