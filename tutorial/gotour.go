package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"math"
	"strings"
)

type Vertext struct {
	X int
	Y int
}

//func Pic(dx, dy int) [][]uint8 {
//	xy := [][]uint8
//	for _, v := range xy {
//		vs := append(v, uint8(dx))
//
//	}
//}
type Vertex struct {
	Lat, Long float64
}

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	fields := strings.Fields(s)
	for _, v := range fields {
		fmt.Println(v)
		words[v]++
	}
	return words
}

func compute(fn func(x, y float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	v := Vertext{X: 1, Y: 2}
	vptr := &v
	fmt.Println(vptr.X)

	//struct array
	vertexes := []Vertext{{1, 1}, {2, 2}, {3, 3}}
	for _, e := range vertexes {
		fmt.Println(e)
	}

	var arr []int
	if arr == nil {
		fmt.Println(len(arr), cap(arr))
	} else {
		fmt.Println("arr is not nil")
	}
	ints := make([]int, 5)
	fmt.Println(ints)

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "x"
	board[2][2] = "o"
	board[1][2] = "x"
	board[1][0] = "o"
	board[0][2] = "x"
	for _, e := range board {
		fmt.Println(strings.Join(e, ""))
	}

	newints := append(ints, 11, 22)
	fmt.Println(newints)

	verticesMap := make(map[string]Vertex)
	verticesMap["bell lab"] = Vertex{11, 22}
	fmt.Println(verticesMap)

	vMap := map[string]Vertex{
		"11": {11, 22},
		"22": {11, 22},
	}
	fmt.Println(vMap)

	score := make(map[string]int)
	score["answer"] = 42
	fmt.Println(score["answer"])
	score["answer"] = 48
	fmt.Println(score["answer"])
	delete(score, "answer")
	answer, ok := score["answer"]
	if ok {
		fmt.Println(answer)
	} else {
		fmt.Println("answer is not exist")
	}

	wc.Test(WordCount)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	u := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(u())
	}
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func Fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		i, j = j, i + j
		return i
	}
}
