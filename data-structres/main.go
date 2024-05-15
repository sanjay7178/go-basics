package main

import (
	"data-strtuctures/genericlist"
	"fmt"
)

func main() {
	glist := genericlist.New[string]() // Assuming New() returns a new list

	glist.Insert("bob")
	glist.Insert("foo")
	glist.Insert("bar")
	glist.Insert("alice")
	glist.RemoveByValue("alice")

	fmt.Printf("%+v\n", glist) // Use Printf for formatted output
}
