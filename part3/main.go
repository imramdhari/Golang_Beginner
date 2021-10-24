package main

import "fmt"

func updateinfo(n *string) {

	*n = "Singh"
	//return n
}

func updatelist(y map[string]int) {

	y["Balvinder"] = 388
}

func main() {
	name := "Monty"
	//s := updateinfo(name)
	//fmt.Println(s)

	info := map[string]int{
		"Balvinder": 303465,
		"Singh":     339930,
		"Goldy":     568644,
	}
	updatelist(info)
	fmt.Println("Memory Address", &name)
	m := &name
	fmt.Println("value ", *m)
	updateinfo(m)
	fmt.Println(name)

}
