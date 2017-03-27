package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter guess: ")
	guess, _ := reader.ReadString('\n')
	fmt.Println("You entered:", guess)
}
