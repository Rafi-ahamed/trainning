package main 

import( 
  "fmt"
)
func main(){
/*
 fmt.Print("enter you are age: ")
 var age int
 fmt.Scanf("%d",&age)
*/

/*
 if  age <3 {
   fmt.Println("infant")
 }else if age>2 && age<13{

   fmt.Println("children")

  }else if age>12 && age<=19{
    
   fmt.Println("teen age") 

 }else{
  fmt.Println("adult")

}
 
*/
 //fmt.Println(age)


//fixed value:
/*
switch age{
case 2:
  fmt.Println("infant")
      fallthrough

case 3,4,5,6,7,8,9,10,11,12:
 fmt.Println("children")

case 13,14,15,16,17,18,19:
 fmt.Println("teen age")

default:
 fmt.Println("adult")

}
*/

//for loop:
/*
for i:=1; i<=9; i++{

  fmt.Println(i)

}
*/

/*
students:=[]string{"rafi","asgor","asif"}

for i,std:=range students{

 fmt.Println(i,std)

}
*/
i:=0

for i<10 {
 fmt.Println(i,"hello")
i++

}



}