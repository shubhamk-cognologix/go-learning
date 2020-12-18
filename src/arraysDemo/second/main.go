package main

import "fmt"

func main() {
	eBooks := [...]string{
		"kafka",
		"Go",
		"Java",
		"React",
	}

	newBooks := eBooks

	for _, v := range newBooks {
		fmt.Printf("%v \n", v)
	}

	studentGrades := [...][3]float64{
		{67, 80, 65},
		{74, 91, 86},
	}
	sum := 0.0
	for i, grades := range studentGrades {
		sum = 0
		for _, grade := range grades {
			sum += grade
		}
		fmt.Printf("\n Sum of student %d: %v", i+1, sum)
	}

}
