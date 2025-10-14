package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// yes, I know this is a useless test file
// just trying to get CI/CD working
type T struct {
	t int
}

func TestSplit(t *testing.T) {

	x := []*T{{1}, {2}, {32}}
	y := []*T{{1}, {2}, {32}}

	compar := reflect.DeepEqual(x, y)

	fmt.Printf("fount that %v", compar)
}

func Test_api(t *testing.T) {
	header := http.Header{}

	header.Set("Authorization", "ApiKey 1234567890")

	if s, err := GetAPIKey(header); s != "1234567890" {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		t.Fatal("Key returend mismatch.")
	}

	header.Set("Authorization", "Failing 12345567")

	if _, err := GetAPIKey(header); err == nil {
		t.Fatal("failed at failing API key detecting ")
	}
	fmt.Print("Testing APi Secusses  succeeded ✅")

}
