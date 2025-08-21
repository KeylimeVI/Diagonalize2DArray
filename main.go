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
	Array   [][]int
	i       *int
	n       int
	Current *int
	Next    Direction
	x       *int
	y       *int
}

func (array *DiagonalizedArray) Diagonalize() {

	*array.Current += *array.i
	*array.i++

	if *array.i == array.n {
		return
	}

	switch array.Next {
	case Down:
		*array.y++
		array.Next = UpRight
	case Right:
		*array.x++
		array.Next = DownLeft
	case UpRight:
		*array.y--
		*array.x++
		if *array.y == 0 {
			array.Next = Down
		} else {
			array.Next = UpRight
		}
	case DownLeft:
		*array.y++
		*array.x--
		if *array.x == 0 {
			array.Next = Right
		} else {
			array.Next = DownLeft
		}
	}
	array.Current = &array.Array[*array.y][*array.x]
	array.Diagonalize()
	return
}

func Diagonal(n int) error {
	if !isSquare(n) {
		return fmt.Errorf("Invalid input")
	}
	sqrtn := int(math.Floor(math.Sqrt(float64(n))))
	arr := make([][]int, sqrtn)
	for i := range arr {
		arr[i] = make([]int, sqrtn)
	}
	i := new(int)
	*i = 1
	x := new(int)
	y := new(int)
	d := &DiagonalizedArray{Array: arr, i: i, Next: Down, Current: &arr[0][0], x: x, y: y, n: n}
	d.Diagonalize()

	for _, row := range d.Array {
		fmt.Println(row)
	}
	return nil
}

func isSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	if sqrt == float64(int(sqrt)) {
		return true
	}
	return false
}
func main() {
	Diagonal(9)
}
