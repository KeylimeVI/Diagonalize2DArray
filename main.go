package main

import (
	"fmt"
	"math"
)

type Direction int

const (
	Down Direction = iota
	Right
	UpRight
	DownLeft
)

type DiagonalizedArray struct {
	Array [][]int
	i     int
	n     int
	Next  Direction
	x, y  int
}

func (array *DiagonalizedArray) Diagonalize() {
	last := int(math.Sqrt(float64(array.n))) - 1

	for array.i = 1; array.i <= array.n; array.i++ {
		array.Array[array.y][array.x] = array.i

		if array.i == array.n {
			return
		}

		switch array.Next {
		case Down:
			array.y++
			array.Next = UpRight

		case Right:
			array.x++
			array.Next = DownLeft

		case UpRight:
			if array.x == last {
				array.y++
				array.Next = DownLeft
			} else if array.y == 0 {
				array.x++
				array.Next = DownLeft
			} else {
				array.x++
				array.y--
			}

		case DownLeft:
			if array.y == last {
				array.x++
				array.Next = UpRight
			} else if array.x == 0 {
				array.y++
				array.Next = UpRight
			} else {
				array.x--
				array.y++
			}
		}
	}
}

func Diagonal(n int) error {
	if !isSquare(n) {
		return fmt.Errorf("Invalid input")
	}
	sqrtn := int(math.Sqrt(float64(n)))
	arr := make([][]int, sqrtn)
	for i := range arr {
		arr[i] = make([]int, sqrtn)
	}

	d := &DiagonalizedArray{Array: arr, n: n, Next: Down}
	d.Diagonalize()

	for _, row := range d.Array {
		fmt.Println(row)
	}
	return nil
}

func isSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return sqrt == float64(int(sqrt))
}

func main() {
	Diagonal(16)
}
