package main

import "fmt"
var age int = 200
func main(){
	//string  formatting verbs 
	var name = "AMAKA"
	fmt.Printf("%s\n",name) //prints the value of plain text
	fmt.Printf("%q\n", name)  //prints the value of double quoted string
	fmt.Printf("%8s\n", name) //
	fmt.Printf("% x\n", age )
	fmt.Printf("%x\n",name) //used in hashing text
	//boolent formating verb
	var ada = true
	var obi = false
	fmt.Printf("%t\n",ada)
	fmt.Printf("%t\n", obi)
	//float formmatting verb
	//var i := 2.35
	//fmt.Printf("%e",i)

}


