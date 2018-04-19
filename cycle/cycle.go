package cycle

import (
    "github.com/jonbodner/leftpad"
)
var DEFAULT_CHAR=' '

func FormatDouble(s string, i int) string {
    return leftpad.Format(s+s,i)
}