package main

import (
	"Comvita/morestrings"
	"Comvita/untils"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	fmt.Println(untils.CustomConcat("!oG ,olleH"))

	f := "apple" // var f string = "apple"
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

	fmt.Println(ff, one, two)
}

func muldiv(n1 int, n2 int) (int, int) {
	// returning the values
	return n1 * n2, n1 / n2
}

func vals(p1 int, p2 int, p3 int) (int, int, int) {
	return p1*2, p2*3, p3*4
}
