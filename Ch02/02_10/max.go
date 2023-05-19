// Calculate maximal value in a slice
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	fmt.Println(nums)

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}

	fmt.Println("Max number: ", max)
}
