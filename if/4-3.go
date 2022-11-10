package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var n, out int
	fmt.Scan(&n)
	var a = make([]string, n)
	var b = make([]string, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a = strings.Split(sc.Text(), " ")
	sc.Scan()
	b = strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			out++
		}
	}
	fmt.Println(out)
}
