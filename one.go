// Main package for go pit of trial
package main

import (
	"fmt"
	"os"
	"strings"
)

func one() {
	var allArgString, sep string

	for i:= 1; i < len(os.Args); i++ {
		allArgString += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])
	fmt.Println(len(os.Args))
	fmt.Println("Hello Go-pit")
	fmt.Println(allArgString)

	for index, arg := range os.Args[1:] {
		fmt.Printf("Index %d Argument %s\n", index, arg)
	}

	fmt.Printf("Joined string is \"%s\"\n", strings.Join(os.Args[1:], " "))
}
