package main
import "fmt"

func main(){
	coun :=1
	for coun <5 {
		coun +=coun
	}
	fmt.Println(coun)
}