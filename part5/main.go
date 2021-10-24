package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Create a new Bill name:")
	//name, _ := reader.ReadString('\n')
	//name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created Bill-", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("choose option(a - add item, s - save bill, t - add tip):", reader)
	//switch statement

	switch opt {
	case "a":
		name, _ := getInput("Choose Name: ", reader)
		price, _ := getInput("Price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("This price must be number")
			promptOptions(b)
		}
		b.addItems(name, p)
		fmt.Println("You choose Add ITEM: "+name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Choose Name: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("This price must be number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Selected tip" + tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("Selected save bills", b.name)
	default:
		fmt.Println("Please select again")
		promptOptions(b)

	}
	//fmt.Println(opt)

}

func main() {
	myBill := createBill()
	promptOptions(myBill)

}
