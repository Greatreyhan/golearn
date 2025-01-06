package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	t := make([]string, 3)
	fmt.Println("emp:", t, "len:", len(t), "cap:", cap(t))

	t[0] = "a"
	t[1] = "b"
	t[2] = "c"
	fmt.Println("set:", t)
	fmt.Println("get", t[2])

	fmt.Println("len", len(t))

	t = append(t, "d")
	t = append(t, "e")
	fmt.Println("append:", t)

	c := make([]string, len(t))
	copy(c, t)
	fmt.Println("cpy:", c)

	l := c[2:5]
	fmt.Println("sl1", l)

	l = c[:5]
	fmt.Println("sl2", l)

	l = c[2:]
	fmt.Println("sl3", l)

	d := []string{"f", "g", "h"}
	fmt.Println("dcl:", d)

	d2 := []string{"f", "g", "h"}
	if slices.Equal(d, d2) {
		fmt.Println("d == d2")
	}

	tds := make([][]int, 3)
	for i := range 3 {
		innLen := i + 1
		tds[i] = make([]int, innLen)
		for j := range innLen {
			tds[i][j] = i + j
		}
	}
	fmt.Println(tds)
}
