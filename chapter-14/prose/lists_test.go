package prose

import "testing"

// Table-driven tests
// Table-driven tests are tests that process “tables” of inputs and
//expected outputs. They pass each set of input to the code being
//tested, and check that the code’s output matches the expected values

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{
			list: []string{},
			want: "",
		},
		{
			list: []string{"apple"},
			want: "apple",
		},
		{
			list: []string{"apple", "orange"},
			want: "apple and orange",
		},
		{
			list: []string{"apple", "orange", "pear"},
			want: "apple, orange, and pear",
		},
		{
			list: []string{"apple", "orange", "pear", "banana", "water melon", "melon", "strawberry"},
			want: "apple, orange, pear, banana, water melon, melon, and strawberry",
		},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = got \"%s\" but want \"%s\"", test.list, got, test.want)
		}
	}

}
