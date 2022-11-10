package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func st_to_int(s string, n int) (sl_int []int) {

	sl := strings.Split(s, " ")
	sl_int = make([]int, n)
	for i := 0; i < n; i++ {
		sl_int[i], _ = strconv.Atoi(sl[i])

	}

	return sl_int

}

func main() {

	var m, n int
	sc.Scan()
	data := strings.Split(sc.Text(), " ")
	m, _ = strconv.Atoi(data[0])
	n, _ = strconv.Atoi(data[1])

	a_mn := make([][]int, 0)

	for i := 0; i < m; i++ {
		sc.Scan()

		t_i := st_to_int(sc.Text(), n)
		a_mn = append(a_mn, t_i)

	}
	sc.Scan()
	x_i := st_to_int(sc.Text(), n)
	for i := 0; i < m; i++ {

		if reflect.DeepEqual(a_mn[i], x_i) == true {
			fmt.Println(i + 1)
		}
	}
}
