package advancefunc

import (
	"math"
)

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func Square(a int) int {
	return a * a
}
func Cube(a int) int {
	return a * a * a
}
func Root(a, b int) int {
	return int(math.Pow(float64(a), float64(1.0/b)))

}

func CubeRoot(a int) int {
	return int(math.Pow(float64(a), 1.0/3))
}

func SquareRoot(a int) int {
	return int(math.Sqrt(float64(a)))
}
