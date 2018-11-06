package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Book struct {
	title  string
	author string
	id     int
}

func printSlice(x []int) {
	fmt.Println("len, cap, slice", len(x), cap(x), x)
}

func factoria(n uint64) uint64 {
	if n > 1 {
		return n * factoria(n-1)
	}
	return 1
}

func finbonacci(n uint64) uint64 {
	if n >= 2 {
		return finbonacci(n-1) + finbonacci(n-2)
	}
	return n
}

type Phone interface {
	call(p string) string
}

type HUaWei struct {
}

func (HUaWei) call(p string) string {
	return p + "hw"
}

type Xiaomi struct {
}

func (Xiaomi) call(p string) string {
	return p + "xm"
}

type DivError struct {
	dividee int
	divider int
}

func (de *DivError) Error() string {
	str := `
   can not procced,
   dividee : %d
   divider : 0
	`
	return fmt.Sprintf(str, de.dividee)
}

func Divide(x, y int) (result int, err string) {
	if y == 0 {
		divError := DivError{x, y}
		err = divError.Error()
		return
	}
	return x / y, ""
}

//naked return, returns the return named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 7
	y = sum - x
	return
}

func main() {
	book := Book{"java", "Gaosinlin", 1}
	book2 := Book{title: "c++", author: "none", id: 2}
	fmt.Println(book)
	fmt.Println(book2)

	var b *Book
	b = &book2
	fmt.Println(b == &book2)
	fmt.Println(*b == book2)
	fmt.Println(b.author, b.title, b.id)

	numbers := make([]int, 3, 5)
	printSlice(numbers)

	//slice apend and copy
	var nums []int
	printSlice(nums)

	nums = append(nums, 1)
	printSlice(nums)
	nums = append(nums, 2, 3, 4)
	printSlice(nums)
	num1s := make([]int, len(nums), cap(nums)*2)
	printSlice(num1s)

	copy(num1s, nums)

	printSlice(num1s)

	//range iterate
	for _, nn := range num1s {
		fmt.Print(nn, "\t")
	}
	smap := map[string]string{"a": "aa", "b": "bb"}
	fmt.Println()
	for k, v := range smap {
		fmt.Println("key:", k, "value:", v)
	}

	//map
	maps := make(map[string]string)
	maps["Italy"] = "Rome"
	maps["UK"] = "London"
	maps["America"] = "WC"
	for e := range maps {
		fmt.Println(e)
	}

	s := maps["UK"]
	if s != "" {
		fmt.Println(s)
	} else {
		fmt.Println("not exist")
	}

	delete(maps, "UK")
	fmt.Println(maps)

	u := factoria(15)
	fmt.Println(u)

	for i := 0; i < 10; i++ {
		fmt.Print(finbonacci(uint64(i)), "\t")
	}
	fmt.Println()

	var phone Phone
	phone = new(HUaWei)
	fmt.Println(phone.call("phone"))
	phone = new(Xiaomi)
	fmt.Println(phone.call("xphone"))

	result, err := Divide(10, 0)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	rand.Seed(10)
	rn := rand.Int()
	fmt.Println(rn)
	fmt.Println(split(17))

	sum := 1
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println("sum:", sum)

	f := pow(3, 2, 10)
	i := pow(3, 3, 20)
	fmt.Println(f, i)
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))

	fmt.Println("When is Saturday ?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tommorrow")
	case today + 2:
		fmt.Println("in tow days")
	default:
		fmt.Println("too far away")
	}

	now := time.Now()
	switch  {
	case now.Hour() < 12:
		fmt.Println("Good morning")
	case now.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	//defers the execution of a function until the surrounding function returns
	defer fmt.Println("world")
	fmt.Println("hello")

	//deferd function calls are pushed onto stack
	fmt.Println("counting")
	for i := 0; i < 10 ; i++  {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	//pointer
	h, j := 42, 2701
	var p *int
	p = &h
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
	p = &j
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(*p)

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g>=%g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
