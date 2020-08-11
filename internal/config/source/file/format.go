package file

import (
	"strings"

	"github.com/infinytum/baerenhoehle/internal/config/encoder"
)

func format(p string, e encoder.Encoder) string {
	parts := strings.Split(p, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return e.String()
}
