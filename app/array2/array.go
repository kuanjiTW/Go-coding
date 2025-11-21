package main

import "fmt"

type Counter struct {
	count int
}

func incr(n any) any {
	switch m := n.(type) {
	case int:
		return any(m + 1)
	case float64:
		return any(m + 1.0)
	case string:
		return any(m + "")
	case *Counter:
		m.count += 1
		return any(m)
	default:
		fmt.Println("unsupported type")
		return n
	}

	// 	if m, ok := n.(int); ok {
	// 		return any(m + 1)
	// 	} else if m, ok := n.(float64); ok {
	// 		return any(m + 1.0)
	// 	} else if m, ok := n.(string); ok {
	// 		return any(m + "")
	// 	} else if m, ok := n.(*Counter); ok {
	// 		m.count += 1
	// 		return any(m)
	// 	} else {
	// 		fmt.Println("unsupported type")
	// 		return n
	// 	}
}

func main() {
	fmt.Println(incr(5))
	fmt.Println(incr(3.14))
	fmt.Println(incr("hello"))
	fmt.Println(incr(&Counter{7}))
	// var x any = 5
	// y := x.(int)
	// fmt.Println(x, y)
	// if z, ok := x.(string); ok {
	// 	fmt.Println(z)
	// }
}
