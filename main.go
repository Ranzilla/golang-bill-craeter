package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose an option: (a) add item | (b) save bill | (t) add tip ", reader)
	fmt.Println(opt)
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
	//fmt.Println(myBill)
}
