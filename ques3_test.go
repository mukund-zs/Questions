package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//TestQues3 is function for testing the pascalTriangle algorithm
func TestQues3(t *testing.T) {
	t1 := [][]int{{1}}
	t2 := [][]int{{1}, {1, 1}}
	t3 := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	t4 := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}, {1, 7, 21, 35, 35, 21, 7, 1}}
	t5 := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}}
	tests := []struct {
		desc   string
		input  int
		output [][]int
	}{
		{"1", 1, t1},
		{"2", 2, t2},
		{"5", 5, t3},
		{"8", 8, t4},
		{"7", 7, t5},
	}
	for _, tc := range tests {
		res := pascalTriangle(tc.input)
		assert.Equal(t, tc.output, res)
	}
}

//BenchmarkQues3 is function for benchmark testing of pascalTriangle algorithm
func BenchmarkQues3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pascalTriangle(2)
	}
}
