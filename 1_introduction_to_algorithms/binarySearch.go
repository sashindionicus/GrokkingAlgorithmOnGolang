package main

import "fmt"

func binarySearch(list []int, item int) (int, bool) {
	low := 0
	high := len(list) - 1

	for low <= high{
		mid := (low + high)
		guess := list[mid]
		if guess == item{
			return mid, true
		}
		if guess > item{
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return - 1, false
}

func main()  {
	myList := []int{1, 3, 5, 7, 9}
	fmt.Println(binarySearch(myList, 1))
}