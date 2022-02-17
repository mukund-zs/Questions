package main

import "testing"

//TestQues1 is function for testing of Leap year Algorithm
func TestQues1(t *testing.T) {
	testCases := []struct {
		desc   string
		input  int
		output string
	}{
		{"Non Leap year", 100, "It's not a leap year"},
		{"Leap Year", 4, "It's a leap year"},
		{"Leap Year", 400, "It's a leap year"},
		{"Non Leap year", 1900, "It's not a leap year"},
		{"Leap year", 2000, "It's a leap year"},
		{"Leap year", 1600, "It's a leap year"},
	}
	for i, tc := range testCases {
		res := checkLeapYear(tc.input)
		if res != tc.output {
			t.Errorf("testcase fails %v", i)
		}
	}
}

//BenchmarkQues1 is function for benchmark testing of Leap Year Algorithm
func BenchmarkQues1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkLeapYear(2000)
	}
}
