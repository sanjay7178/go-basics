package main

import "fmt"

func addUsers (users []string){
	for _, user :=  range users {
		fmt.Println(user)
	}
}


// func main (){
// 	addUsers("")
ee