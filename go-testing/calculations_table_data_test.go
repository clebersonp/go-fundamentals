package go_testing

import (
	"testing"
)

type testData struct {
	numbers []int
	want    int
}

// run a specific file test: go test -v calculations_table_data_test.go calculations.go

func TestShouldSumNumbers(t *testing.T) {

	tests := []testData{
		{
			numbers: []int{},
			want:    0,
		},
		{
			numbers: []int{10},
			want:    10,
		},
		{
			numbers: []int{5, 6},
			want:    11,
		},
		{
			numbers: []int{8, 5, 6},
			want:    19,
		},
	}
	for _, test := range tests {
		result := Add(test.numbers...)
		if result != test.want {
			t.Errorf("expected '%d' but current '%d'", result, test.want)
		}
	}
}
