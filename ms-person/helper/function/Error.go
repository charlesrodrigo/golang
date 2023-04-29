package function

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateResponseError(codeError int, messages ...string) (int, interface{}) {
	if len(messages) == 1 {
		return codeError, gin.H{"code": codeError, "errors": strings.Split(messages[0], "\n")}
	}
	return codeError, gin.H{"code": codeError, "errors": messages[0]}
}
