/*
(20 points)

	Write a program that performs the following two actions.
	  - Iterate over the following slice and prints the slice
	    values.
	    ```
	    []string{"Actions", "speak", "louder", "than", "words"}
	    ```
	    Do not print any associated index.
	  - Iterate over the following map and prints each key-value
	    pair (k, v) in the format "k => v".
	    ```
	    map[int]string{1: "Actions", 3: "speak", 5: "louder", 7: "than", 10: "words"}
	    ```
*/
package main

import "fmt"

func main() {
	slice := []string{"Actions", "speak", "louder", "than", "words"}

	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Println()

	m := map[int]string{1: "Actions", 3: "speak", 5: "louder", 7: "than", 10: "words"}
	for k, v := range m {
		fmt.Printf("%d => %s\n", k, v)
	}
}
