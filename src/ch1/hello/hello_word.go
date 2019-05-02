package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) >1 {
		fmt.Printf("Hello World",os.Args[1])
	}
}