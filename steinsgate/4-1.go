package main

import "fmt"

func main() {

	var n, m, q int

	fmt.Scan(&n, &m, &q)

	n_i := make([]int, n)
	s_i := make([]int, m)
	e_i := make([]int, q)
	t_i := make([]int, q)

	for i := 0; i < m; i++ {
		fmt.Scan(&s_i[i])
		n_i[s_i[i]] = 1
	}
	for i := 0; i < q; i++ {
		fmt.Scan(&e_i[i], &t_i[i])
	}
	for i := 0; i < q; i++ {
		if n_i[t_i[i]] == 1 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}
