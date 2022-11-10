package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var s string
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s = sc.Text()

	s = strings.Replace(s, ", maybe.", "!!", 1)
	fmt.Println(s)
}
