package counter

import "fmt"

type Counter interface {
	fmt.Stringer
	Inc() error
}
