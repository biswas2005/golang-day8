package set1

import "fmt"

func MCQ() {

	questions := [3]string{"1.What is the capital of India?\nA)Hyderabad\nB)Kolkata\nC)Delhi\nD)Mumbai",
		"2.Who is refered as 'Modern day master'(In cricket)?\nA)Virat Kohli\nB)Rohit Sharma\nC)Dhoni\nD)Babar",
		"3.Who won the most ODI Men's world cups?\nA)England\nB)India\nC)Newzealand\nD)Australia"}

	answers := [3]string{"C", "A", "D"}
	score := 0

	for i := 0; i < len(questions); i++ {
		fmt.Println(questions[i])
		var userAns string
		fmt.Println("Enter your choice: ")
		fmt.Scan(&userAns)
		if userAns == answers[i] {
			fmt.Println("Correct")
			score++
		} else {
			fmt.Println("Incorrect")
		}

	}
	fmt.Printf("Your final Score is %d/%d", score, len(questions))

}
