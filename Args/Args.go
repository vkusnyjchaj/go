// Reailization of Unix echo command. Prints the command line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var str string = os.Args[1]
	var separator string = " "

	for i := 2; i < len(os.Args); i++ {
		str += separator + os.Args[i]
	}

	fmt.Println(str)
}
