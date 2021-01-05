package main

import "fmt"

func main() {

	myList := []string{"windows", "osx", "linux"}

	// syntax: for {key}, value := range {list}
	for _, animal := range myList {
		fmt.Println("My preferred OS is: ", animal)
	}

	// same for a map
	myOtherList := map[string]string{
		"dog":      "woof",
		"cat":      "meow",
		"hedgehog": "sniff",
	}

	for animal, noise := range myOtherList {
		fmt.Println("The", animal, "went", noise)
	}

}
