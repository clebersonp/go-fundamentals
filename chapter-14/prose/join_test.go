// join_test.go
// _test.go is an important part; the go test tools looks for files named with that suffix
package prose

import (
	"fmt"
	"testing"
)

// Function name should begin with "Test"
// Name after "Test" can be whatever you want
// command line:
// outside of package: go test "package name" -> go test github.com/clebersonp/prose
// inside of package: go test
// go test -v -> for verbose
// limit which tests will be run: go test -v -run TestOneElement
// limit which tests will be run which matches the name: go test -v -run Elements

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) { // function will be passed a pointer to a testing.T value
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestWithManyElements(t *testing.T) {
	list := []string{"apple", "orange", "pear", "banana", "water melon", "melon", "strawberry"}
	want := "apple, orange, pear, banana, water melon, melon, and strawberry"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = got \"%s\" but want \"%s\"", list, got, want)
}
