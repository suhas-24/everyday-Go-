package main

import (
	"fmt"
	"os"
)

func main(){
	name := os.Getenv("USER")
	fmt.Printf("Welcome MR. %s ", name)
}