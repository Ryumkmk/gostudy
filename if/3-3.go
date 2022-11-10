package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var e, o int = 0, 0
	data := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		in, _ := strconv.Atoi(data[i])
		// im = append(im, in)
		if in%2 == 0 {
			e++
		} else {
			o++
		}
	}
	fmt.Println(e, o)

}
