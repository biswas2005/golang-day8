package set1

import "fmt"

func Students() {

	var student int
	fmt.Println("Enter student number: ")
	fmt.Scan(&student)

	switch student {
	case 1:
		fmt.Println("Average marks:", average([]int{40, 50, 32, 34, 55})/5)

	case 2:
		fmt.Println("Average marks:", average([]int{65, 55, 52, 44, 59})/5)
	case 3:
		fmt.Println("Average marks:", average([]int{67, 59, 39, 30, 55})/5)
	case 4:
		fmt.Println("Average marks:", average([]int{49, 80, 33, 74, 65})/5)
	default:
		fmt.Printf("Wrong student ID")
	}

}

func average(marks []int) float64 {

	if len(marks) == 0 {
		fmt.Println("Invalid operation.")

	}
	sum := 0

	for _, value := range marks {
		sum += value

	}
	return float64(sum)

}
