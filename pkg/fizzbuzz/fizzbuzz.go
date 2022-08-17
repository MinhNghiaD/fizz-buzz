package fizzbuzz

import "strconv"

// FizzBuzz returns a list of strings with numbers from 1 to limit,
// where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2
func FizzBuzz(int1, int2, limit int, str1, str2 string) []string {
	output := make([]string, limit)
	for i := 1; i <= limit; i++ {
		if i%int1 == 0 {
			output[i-1] += str1
		}
		if i%int2 == 0 {
			output[i-1] += str2
		}

		if output[i-1] == "" {
			output[i-1] = strconv.Itoa(i)
		}
	}

	return output
}
