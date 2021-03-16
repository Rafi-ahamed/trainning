package main

import"fmt"
func update(a *int){
fmt.Println(a)
*a=*a+10
}
func main( ){

 var x int
 var y *int 
 fmt.Println("x memory address is",&x)
 fmt.Println("x value is", x)

 fmt.Println("y memory address is",&y)
 fmt.Println("y value is",y)


 x=10 //(assignment or initialization)
 y=&x //(pointer or referencing)
 fmt.Println("x value is", x) //(accessing)
 fmt.Println("y value is",y)
 fmt.Println("y is",*y) //(dereferencing)
  update(y)

 fmt.Println(x)
}