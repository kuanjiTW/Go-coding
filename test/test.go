package main

func main() {
	xs := []int{7, 3, 1, 9, 6, 2, 8}
	//find the lagest number
	largest := xs[0]
	for _, v := range xs {
		if v > largest {
			largest = v
		}
	}
	//find the smallest number
	smallest := xs[0]
	for _, v := range xs {
		if v < smallest {
			smallest = v
		}
	}
	println("largest:", largest, "smallest:", smallest)

	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			if xs[j] < xs[i] {
				xs[i], xs[j] = xs[j], xs[i]
			}
		}
	}
	println("smallest:", xs[0], "largest:", xs[len(xs)-1])
}
