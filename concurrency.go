package main

import "fmt"
// goroutines <-->  threads in go 
//Go has built-in channels which are used to share data between goroutines.

func cookingGopher(id int, c chan int){
	fmt.Println("gopher", id, "started cooking")
	c <- id   // Send a value back to main   ----out----
}

func main(){
	c:= make(chan int) // Create a channel to pass ints

	for i:=0; i<5;i++{
		go cookingGopher(i,c)
	}

	for i:=0; i<5;i++{
		gopherID := <- c //Recieve a value from a channel  and  continue
		fmt.Println("gopher",gopherID,"finished the dish")
	} // All goroutines are finished at this point ??
}


