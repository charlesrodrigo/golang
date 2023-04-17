package function

import (
	"strings"

	"br.com.charlesrodrigo/ms-person/helper/logger"
	"github.com/gin-gonic/gin"
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

func CreateResponseError(codeError int, messages ...string) (int, interface{}) {
	if len(messages) == 1 {
		return codeError, gin.H{"code": codeError, "errors": strings.Split(messages[0], "\n")}
	}
	return codeError, gin.H{"code": codeError, "errors": messages[0]}
}
