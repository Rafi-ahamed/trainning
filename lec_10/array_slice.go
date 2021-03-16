package main 

import "fmt"
import "reflect"

func main( ){
 //array syntax:

//var students [3]string
//students[0]="rafi"
//students[1]="simanto"
//students[2]="maruf"


 //fmt.Println(students)
 //fmt.Println(len(students))
 //fmt.Println(students[2])
//}
 //array string literal:for representing a fixed value in source code
//students:=[...]string{"rafi","simanto","maruf","maruf 2","raju" }
 //fmt.Println(students,len(students))

 //slice syntax:

 //x:= make([]string,0)
var fruits[] string
fruits=append(fruits,"banana","mango")
 //fmt.Println(fruits)
//fmt.Println(len(fruits))
 //fmt.Printf("%T\n",fruits)
 //fmt.Printf("%T",students)

a:=reflect.TypeOf(fruits).kind().String()
 fmt.Println(a)
}