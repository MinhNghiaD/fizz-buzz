package fizzbuzz_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/MinhNghiaD/fizz-buzz/pkg/fizzbuzz"
)

// TestFizzBuzz tests functionality of fizz buzz
func TestFizzBuzz(t *testing.T) {
	testcases := []struct {
		int1           int
		int2           int
		limit          int
		str1           string
		str2           string
		expectedOutput []string
	}{
		{
			3,
			5,
			15,
			"fizz",
			"buzz",
			[]string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
		{
			2,
			4,
			15,
			"fizz",
			"buzz",
			[]string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
		{
			1,
			5,
			20,
			"fizz",
			"buzz",
			[]string{"fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz", "fizz", "fizz", "fizz", "fizz", "fizzbuzz"},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			t.Parallel()
			output := fizzbuzz.FizzBuzz(tc.int1, tc.int2, tc.limit, tc.str1, tc.str2)

			if !reflect.DeepEqual(output, tc.expectedOutput) {
				t.Errorf("Test case %d returns %v, while %v was expected", i, output, tc.expectedOutput)
			}
		})
	}
}
