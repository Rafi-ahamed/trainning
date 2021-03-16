package main 

import "fmt"
//import "reflect"

func main( ){

//maps syntax:
x:=make(map[string]string)
x["name"]="rafi"
x["height"]="10.5"
delete(x,"height")

fmt.Println(x["name"])
}