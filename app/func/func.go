package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	var f func(int, int) int = add
// 	r := f(5, 3)
// 	fmt.Println(r)
// }

func smap[T1 any, T2 any](f func(T1) T2, ns []T1) []T2 {
	rs := []T2{}
	for _, n := range ns {
		rs = append(rs, f(n))
	}
	return rs
}

func mul2(n int) int {
	return n * 2
}

func incr1(n int) int {
	return n + 1
}

func sfold[T1 any, T2 any](f func(T2, T1) T2, i T2, ns []T1) T2 {
	for _, n := range ns {
		i = f(i, n)
	}
	return i
}

func addr(r int, n int) int {
	return r + n
}

func mulr(r int, n int) int {
	return r * n
}

func addf(x float64) float64 {
	return x + 1.0
}

func main() {
	ns := []int{1, 2, 3, 4, 5}

	//smap(addf,ns)

	ms1 := smap(mul2, ns)
	ms2 := smap(incr1, ns)
	fmt.Println(ms1)
	fmt.Println(ms2)
	sum := sfold(addr, 0, ns)
	fmt.Println(sum)
	prod := sfold(mulr, 1, ns)
	fmt.Println(prod)
}
