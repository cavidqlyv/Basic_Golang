package main

import "fmt"
//return is writen in here  ⬇️⬇️⬇️⬇️
func average(x []float64) (avg float64){
	total := 0.0
	if len(x) == 0{
		avg = 0
	} else {
		for _,v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}

	return // avg
}

func main() {
	x := []float64{2.71, 3.14, 42.0, 29.5}
	fmt.Println(average(x))
}