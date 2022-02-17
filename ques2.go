package main

//checkTriangle is function to check type of triangle
func checkTriangle(x, y, z int) string {
	if x+y < z || y+z < x || x+z < y {
		return "Triangle is not possible"
	}
	if x == y && y == z {
		return "Equilateral and isosceles"
	} else if (x == y && y != z) || (x == z && y != x) || (y == z && x != y) {
		return "Isosceles"
	} else {
		return "scalene"
	}
}
