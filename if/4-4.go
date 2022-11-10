package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n, out int
	fmt.Scanln(&n)
	var a = make([]string, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a = strings.Split(sc.Text(), " ")
	for _, st := range a {
		in, _ := strconv.Atoi(st)
		fmt.Println(in)
		if in%2 != 0 {
			break
		}
		out += in
	}
	fmt.Println(out)
}
