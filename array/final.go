package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func swap_a_b(data []string, lin_p *[]int) {
// 	var a, b int
// 	a, _ = strconv.Atoi(data[1])
// 	b, _ = strconv.Atoi(data[2])
// 	x := lin_p*[a-1]
// 	lin_p[a-1] = lin_p[b-1]
// 	lin_p[b-1] = x

// }
func reverse(data []string, lin_p *[]int) {

}
func resize_c(data []string, lin_p *[]int) {
	var c int
	c, _ = strconv.Atoi(data[1])
}

var sc = bufio.NewScanner(os.Stdin)

func main() {

	var n, q int
	fmt.Scanln(&n, &q)
	var lin = make([]int, n)

	for i := 0; i < n; i++ {
		lin[i] = i + 1
	}
	lin_p := &lin

	for i := 0; i < q; i++ {
		sc.Scan()
		data := strings.Split(sc.Text(), " ")
		switch len(data) {
		case 1:
			reverse(data, lin_p)
		case 2:
			resize_c(data, lin_p)
		case 3:
			swap_a_b(data, lin_p)
		}
	}
}
