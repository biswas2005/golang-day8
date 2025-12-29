package pointers

import "fmt"

func Basics() {

	value := 100
	x := 10
	y := 20

	var p *int

	fmt.Println("The value initially was", value)
	fmt.Println("The address of value", &value)
	p = &value
	*p = 200
	fmt.Println("The value was later changed to", *p)
	fmt.Println("Address is", p)

	//swap() swaps the values
	//address of variable remains same
	fmt.Printf("\nValue of a %d and b %d before swap", x, y)
	fmt.Printf("\nAddress before swap a is %x and b is %x", &x, &y)
	swap(&x, &y)
	fmt.Printf("\nValue of a %d and b %d after swap", x, y)
	fmt.Printf("\nAddress after swap a is %x and b is %x", &x, &y)

	//array() called
	array()

}

//swaps the values
func swap(a, b *int) {

	*a, *b = *b, *a

}

//prints address of each element of array
func array() {

	arr := [3]int{2, 3, 1}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("\nAddress of %d is %x", arr[i], &arr[i])
	}
}
