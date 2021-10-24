package main

import (
	"fmt"
	"strings"
)

func main() {
	
		//string , integer

		var numOne string = "Monty"
		var numTwo = "Singh"
		numThree := "Mel"

		var numFour int = 30

		var numFive float64 = 34444.433

		var ages = [3]int{20, 30, 45}
		names := [4]string{"monty", "lucky", "harman", "neha"}

		fmt.Println(numOne, numTwo, numThree, numFour, numFive)
		fmt.Printf("numfive  %0.3f", numFive)
		fmt.Print("my name is khan")
		fmt.Printf("my name is khan %v  and length %v", ages, len(ages))
		fmt.Printf("my name is khan %v  and length %v \n", names, len(names))

		list := []int{10, 20, 30, 50, 7}
		greet := "Hello World"
		name1 := [4]string{"monty", "lucky", "harman", "neha"}
		list = append(list, 78)
		rangeOne := list[1:]

		fmt.Println(strings.Index(greet, "ll"))

		fmt.Println(strings.Split(greet, " "))
		sort.Ints(list)
		fmt.Println(list)
		index := sort.SearchInts(list, 20)
		fmt.Println(list, index)

		fmt.Println(list, name1)
		fmt.Println(rangeOne)

		//loops
		x := 0
		for x < 5 {

			fmt.Println("Value of x is :", x)
			x++

		}

		for x := 0; x < 5; x++ {

			fmt.Println(x)
		}

		for index, value := range name1 {
			fmt.Printf("Index %v and Value %v", index, value)

		}

		// conditions and booleans

		num := 10
		fmt.Println((num == 10)) // boolean

		if num > 5 {
			// if condition
			fmt.Println("Test 1 ")
		} else if num < 5 {
			// if condition
			fmt.Println("Test 2 ")
		} else {
			fmt.Println("nothing")
		}
		sayGreeting("Monty")
		sayGoodbye("Ghaman")
		cycleNames([]string{"Mario", "Monty"}, sayGoodbye)
		ans := cycleArea(13.23)
		fmt.Println(ans)*/
	fn1, sn1 := getInitials("Balvinder Singh Ghaman")
	fmt.Println(fn1, sn1)

}


//Using functions
func sayGreeting(s string) {
	fmt.Printf("Good Morning %v\n", s)
}
func sayGoodbye(s string) {
	fmt.Printf("Good Bye %v\n", s)
}
func cycleNames(n []string, f func(string)) {

	for _, v := range n {
		f(v)
	}

}

func cycleArea(r float64) float64 {

	return math.Pi * r * r
}*/

//Return Multiply Values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string

	for _, v := range names {

		initials = append(initials, v[:1])

	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
