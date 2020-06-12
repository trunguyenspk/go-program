package main

import (
	"fmt"
	//	"math"
	"time"
	//	"strings"
	//	"Godemo/service"
	//	"github.com/google/go-cmp/cmp"
)

func main() {

	defer fmt.Println("==================>defer world")

	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0:
			fmt.Println(" Today.")
		case today + 1:
			fmt.Println(" Tomorrow.")
		case today + 2:
			fmt.Println(" In two days.")
		default:
			fmt.Println(" Too far away.")
	}

	/*sum := 1
	for sum < 1000 {
		sum += sum
		if sum == 8 {
			fmt.Printf("Break at %d", sum)
			break
		}
	}

	u, _ := split(12)
	fmt.Println(u)*/

	//fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	//fmt.Println(untils.CustomConcat("!oG ,olleH"))

	/*f := "apple" // var f string = "apple"
	fmt.Println(f)

	b := []int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	delete(m, "k2")

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	mul, _ := muldiv(105, 7)

	// only using the mul variable
	// compiler will give an error
	fmt.Println("105 x 7 = ", mul)

	ff, one, two := vals(1, 2, 3)

	fmt.Println(ff, one, two)*/

}

func muldiv(n1 int, n2 int) (int, int) {
	return n1 * n2, n1 / n2 // returning the values
}

func vals(p1 int, p2 int, p3 int) (int, int, int) {
	return p1 * 2, p2 * 3, p3 * 4
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
