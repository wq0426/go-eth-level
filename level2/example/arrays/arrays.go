package arrays

import "fmt"

func Arrays() {
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("fix array:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("fix array2:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("fix array3:", b)

	twoD := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("more array: ", twoD)
}
