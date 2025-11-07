package main

import (
	"fmt"
	"os"
)

func main() {
	ns := []int{1, 2, 3}
	fmt.Println(ns)
	ns = append(ns, 4, 5, 6)
	fmt.Println(ns)

	ms := [...]int{1, 2, 3}
	//ms = append(ms,4) not work
	fmt.Println(ms)

	for i, v := range os.Args {
		fmt.Printf("%d: %#v\n", i, v)

	}

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := a[4:6]
	a[5] = 100
	fmt.Println("切片b:", b)

	ns = []int{1, 2, 3, 4, 5}
	cs := ns[1:4]
	ns[0] = 666
	cs[1] = 777

	fmt.Println(ns)
	fmt.Println(cs)

	ns = []int{0, 1, 2, 3, 4, 5}
	cs = ns[1:4]
	fmt.Println(ns, len(ns), cap(ns))
	fmt.Println(cs, len(cs), cap(cs))
	ns = append(ns, 6)
	ns[1] = 777
	fmt.Println(cs[0])

}
