package main

import (
	"os"
	"bufio"
	"fmt"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	fmt.Println("Hello, World.")
	fmt.Println(inputString)
}
