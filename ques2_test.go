package main

import "testing"

//TestQues2 is function to test checkTriangle algorithm
func TestQues2(t *testing.T) {
	testCases2 := []struct {
		desc   string
		inputx int
		inputy int
		inputz int
		output string
	}{
		{"Greater than 2", 3, 3, 3, "Equilateral and isosceles"},
		{"Only 2", 5, 3, 3, "Isosceles"},
		{"only 2", 3, 6, 6, "Isosceles"},
		{"All equal", 5, 5, 5, "Equilateral and isosceles"},
		{"No equal", 5, 6, 7, "scalene"},
		{"Triangle not possible", 1, 5, 1, "Triangle is not possible"},
	}
	for i, tc := range testCases2 {
		res := checkTriangle(tc.inputx, tc.inputy, tc.inputz)
		if res != tc.output {
			t.Errorf("testcase fails %v", i)
		}
	}
}

//BenchmarkQues2 is function for benchmark testing of checkTriangle algorithm
func BenchmarkQues2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkTriangle(1, 2, 3)
	}
}
