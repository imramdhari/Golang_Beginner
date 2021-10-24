package main

import "fmt"

func main() {

	mybill := newBill("Monty bill")
	mybill.addItems("onion soup", 4.50)
	mybill.addItems("Cake", 6.50)
	mybill.updateTip(10.43)
	fmt.Println(mybill.format())
}
