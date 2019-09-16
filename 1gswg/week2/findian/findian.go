package main

import(
	"fmt"
	"strings"
)

const (
	found = "Found!"
	notFound = "Not Found!"
)

func main(){
	fmt.Println("Enter a string")
	var s string
	fmt.Scan(&s)
	if !strings.HasPrefix(s, "i") && !strings.HasPrefix(s, "I") {
		fmt.Println(notFound)
		return
	}

	if !strings.HasSuffix(s, "n") && !strings.HasSuffix(s, "N") {
		fmt.Println(notFound)
		return
	}

	if !strings.ContainsAny(s, "aA") {
		fmt.Println(notFound)
		return
	}

	fmt.Println(found)
}
