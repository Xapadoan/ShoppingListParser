package application

import (
	"strings"

	dom "github.com/Xapadoan/shplsprsr/logger/domain"
)

func CompactFormatter(level dom.LogLevel, msg ...string) string {
	return "[" + strings.ToUpper(level.String()) + "]: " + strings.Join(msg, " ")
}
