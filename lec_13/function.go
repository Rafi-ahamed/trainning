 //  # Function means function is a group of statement that exit within a program for the  purpose  of performing a specific  task, at a high level, a function  task an input & returns an  output.
//exam:1
package main 
import "fmt"

//func add (x int,y int)int{
//r:=x+y
//return r
//}
//func  main (){
//x:=add(10,30)
//fmt.Println(x)
//}

   //#exam:2
/*
func add(x,y int)int{
r:=x+y
return r
}
func main ( ){
x:=add(10,30)
fmt.Println(x)
}
*/

   // #exam3:
/*
func add(x,y int)(r int){
r=x+y
return r
}
func main( ){
x:=add (10,30)
fmt.Println(x)
}
*/
   // #exam4:
/*
func add(x,y int)(r int){
r=x+y
return 
}
func main( ){
x:=add (10,30)
fmt.Println(x)
}
*/
    // MULTIPLE RETURN VALUE:
/*
func rectangle(l int,b int) (area int,parameter int){
parameter=2*(1+b)
area=1*b
return
}
func main( ){
a,p:=rectangle(10,10)
fmt.Println(a,p)
}
*/

    //# passing adress to afunction:
//pointer means a special variable where includes other variable memory adress..
/*
func update (a *int,t *string){
 *a=*a+5
 *t=*t+"doe"
 return
}
func main( ){
number:=10
name:="rafi"
update(&number,&name)
fmt.Println(number,name)
}
*/


var(
a=func(x ,y int)(r int){
r=x*y
return
}
)
func main( ){


fmt.Println(a(10,10))
}    