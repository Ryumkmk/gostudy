package main

import (
	"fmt"
)

var s_i = make([]int, 0)
var t_i = make([]int, 0)
var out = make([]int, 0)

func move(s int, t int, m int) {

	if s != t {

		if t > m {
			return
		}
		out = append(out, t)
		move(s_i[t-1], t_i[t-1], m)
	} else {
		return
	}

}

func main() {

	var n, m int

	fmt.Scan(&n, &m)
	for i := 0; i < m; i++ {

		var s, t int
		fmt.Scan(&s, &t)
		s_i = append(s_i, s)
		t_i = append(t_i, t)
	}
	out = append(out, 1)
	if s_i[0] != t_i[0] {
		move(s_i[0], t_i[0], m)
	}

	for i := len(out); i > 0; i-- {
		fmt.Printf("%d", out[i-1])
		if i != 1 {
			fmt.Print(" ")
		}
	}

}
