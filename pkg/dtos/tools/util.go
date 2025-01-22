package tools

import (
	"strings"
)

func EsStringVacio(valor string) bool {
	return len(strings.TrimSpace(valor)) == 0

}
