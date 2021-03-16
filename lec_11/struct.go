//A struct is a collection of field or properties

package main

import "fmt"
/*
var w1 weight

//func main(){


w1=30.50
fmt.Println(w1,name)
*/

//}



type book struct{
title string
author string
isbn string
price float32
pages int
}


//func main(){

//var b1 book
//b1.title="an introduction to programming in go"
//b1.author="CALLEB DOXEY"
//b1.isbn="978-1478355823"
//b1.price=10.50
//b1.pages=165
//fmt.Println(b1)
//fmt.Println(b1.title)
//}

//func main(){
/*
b1:=struct{
title string
author string
isbn string
price float32
pages int
}{
 title:"an introduction to programming in go",
 author:"CALLEB DOXEY",
 isbn:"978-1478355823",
 price:10.50,
 pages:165,
}
fmt.Println(b1)
fmt.Println(b1.title)
fmt.Println(b1.author)
fmt.Println(b1.isbn)
fmt.Println(b1.price)
fmt.Println(b1.pages)

*/
func main(){

var b1=book {title:"an introduction to programming in go",
 author:"CALLEB DOXEY",
 isbn:"978-1478355823",
 price:10.50,
 pages:165,}


fmt.Println(b1)
fmt.Println(b1.title)
fmt.Println(b1.author)
fmt.Println(b1.isbn)
fmt.Println(b1.price)
fmt.Println(b1.pages)

}