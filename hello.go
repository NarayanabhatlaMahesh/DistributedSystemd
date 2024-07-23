package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Variable declaration")
	// variable declaration
	 var var1='z';
	 // checking the type of the variable
    fmt.Println(reflect.TypeOf(var1));
	var2 := 1

	//printing two variables in same println
	fmt.Println("var2" , var2);
	//checking the type of variable
    fmt.Println(reflect.TypeOf(var2))
	fmt.Println("\nLoops\nfor-loop");

	/* looping 
		statements */
	var i = 0
	//for loop
	for i=0;i<5;i++{
		fmt.Print("i-");fmt.Println(i);
	}
	s :=""
	fmt.Println("\nwhile-loop")
	//while loop wth sprintf
	for i > var2{
		s+= fmt.Sprintf("i is --- %d",i);
		s+="\n"
		i=i-1;
	}
	fmt.Print(s);

	//functions
	fmt.Println("\nFunctions")
	fmt.Print(sum(2,3))

}

func sum(num1 int ,num2 int) int{
	return num1+num2
}