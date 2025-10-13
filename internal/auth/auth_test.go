package auth

import (
	"fmt"
	"reflect"
	"testing"
)

// yes, I know this is a useless test file
// just trying to get CI/CD working
type T struct {
	t int
}

func TestSplit(t *testing.T) {

	x := []*T{{1}, {2}, {3}}
	y := []*T{{1}, {2}, {3}}

	compar := reflect.DeepEqual(x, y)

	fmt.Printf("fount that %v", compar)
}
