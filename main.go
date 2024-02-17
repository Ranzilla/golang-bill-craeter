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

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose an option - (a) add item | (s) save bill | (t) add tip: ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Printf("'%v' is not a valid input. The price must be a number.\n", price)
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Printf("Item '%v' has been added to the bill priced at '$%0.2f'.\n", name, p)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You saved the bill: ", b.name)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Printf("'%v' is not a valid input. The price must be a number.\n", tip)
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Printf("A tip of '$%0.2f' has been added to the bill.\n", t)
		promptOptions(b)

	default:
		fmt.Printf("'%v' is not a valid option. Choose (a), (s) or (t).\n", opt)
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Bill created: ", b.name)
	return b
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
