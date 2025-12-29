package pointers

import "fmt"

func Practice() {

	nums := [4]int{2, 3, 4, 5}
	nums2 := []int{2, 3, 4, 5}
	//prints the original array
	fmt.Println("Array before modify", nums)
	modify(&nums)

	fmt.Println("After modify", nums)
	modify2(nums2)
	fmt.Println("Without pointers", nums2)

}

//modify() using pointer
func modify(arr *[4]int) {

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
		if arr[i] == 4 {
			arr[i] = arr[i] + 7
		}
	}
}

//modify2() without using pointer
func modify2(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] * 2
	}
}
