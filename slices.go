package main
import (
    "fmt"
    "reflect"
)
func main(){
	languages :=[9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // needs comma else "syntax error: unexpected newline, expecting comma or }"
	}

	classics :=languages[:3]

	modern := make([]string,4) //created 4 length array 
	modern = languages[3:7]   // := is used when creating new variable

	new := languages[7:]

	fmt.Printf("classic languagues: %v\n", classics) 
	fmt.Printf("modern languages: %v\n", modern)    
	fmt.Printf("new languages: %v\n", new)

	allLangs :=languages[:]
	fmt.Println(reflect.TypeOf(allLangs).Kind())   //Type = Slice

	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks  := frameworks[:4:4]    // length 4 capacity 4
	frameworks = append(frameworks, "Meteor") 

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", jsFrameworks)
}