package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func st_to_int(s string, m int) (sl_int []int) {

	sl := strings.Split(s, " ")
	sl_int = make([]int, m)
	for i := 0; i < m; i++ {
		sl_int[i], _ = strconv.Atoi(sl[i])

	}

	return sl_int

}

func main() {

	var n, m, l int
	fmt.Scan(&n, &m, &l)

	sc := bufio.NewScanner(os.Stdin)

	d_i := make([][]int, 0)
	p_i := make([][]int, 0)

	for i := 0; i < n; i++ {
		sc.Scan()

		d_i = append(d_i, st_to_int(sc.Text(), m))
	}
	for i := 0; i < l; i++ {
		sc.Scan()

		p_i = append(p_i, st_to_int(sc.Text(), m))

	}

	for i := 1; i < l; i++ {

		diff := make([]int, 0)
		for j := 0; j < m; j++ {
			diff = append(diff, p_i[i][j]-p_i[i-1][j])

		}
		for q := 0; q < n; q++ {

			if reflect.DeepEqual(diff, d_i[q]) {
				fmt.Println(q + 1)

			}

		}
	}
}
