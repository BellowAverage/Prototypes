package main

import "fmt"

func main() {
	var arr [3]int32 = [3]int32{1, 2, 3}
	arr2 := [10]int32{2, 3, 1}
	fmt.Println(arr)

	arr[0] = 3
	fmt.Println(arr)

	fmt.Println(arr[0])

	fmt.Println(arr2)

	nums := make([]int32, 5, 8)
	appended := [4]int32{12, 13, 14, 15}
	nums = append(nums, appended[:]...)
	fmt.Println(nums)

	var mp = map[int32]string{1: "one", 2: "two", 3: "three"}

	var age, ok = mp[4]
	if ok {
		fmt.Println(age)
	}

	fmt.Println(mp)

	for key, value := range mp {
		fmt.Printf("%v: %v\n", key, value)
	}

	fmt.Println("-------------")

	name := []rune("Hello")
	for key, value := range name {
		fmt.Printf("%v: %c\n", key, value)
	}
}
