package extratime

import "strings"

func cut(b []byte) string {
	return strings.TrimSuffix(strings.TrimPrefix(string(b), `"`), `"`)
}
