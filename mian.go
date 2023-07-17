package main

import "fmt"

func arrange(a []int, l, r int) {
	for ; l <= r; l++ {
		for m := l; m <= r; m++ {
			if a[m] < a[l] {
				c := a[m]
				a[m] = a[l]
				a[l] = c
			}
		}
	}
}
func main() {
	var n, l, r int
	fmt.Println("请输入n,l,r")
	fmt.Scan(&n, &l, &r)
	a := make([]int, n)
	fmt.Println("请输入n个整数")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	arrange(a, l, r)
	fmt.Println(a)
}
