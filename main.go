package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func main() {
	greetMessageEmpty := hello("")
	fmt.Println(aurora.Yellow(greetMessageEmpty))

	greetMessageDavid := hello("David")
	fmt.Println(aurora.Yellow(greetMessageDavid))
}
