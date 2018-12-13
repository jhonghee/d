package d

import (
	"fmt"

	"github.com/jhonghee/e"
)

// Version returns the tagged version of the module
func Version() string {
	return fmt.Sprint("D v1.2", "->", e.Version())
}
