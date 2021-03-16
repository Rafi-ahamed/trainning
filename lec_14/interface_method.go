package main

import "fmt"

type Email interface{
send()string
read()
write(string,string,string)

}
type test struct{
to  string
from string
sub string
message string
}

/*
type doctor struct {
name string
education string
age int
salary float32

}
*/

func (t test)write(to string,from string,sub string){

 fmt.Println(to,from,sub)
}
func (t test)send() string{
 return t.to
 //fmt.Println(t.to,"email send")

}
func (t test)read(){

 fmt.Println(t.from,"receive from")

}
/*
type doctor struct{
name string
education string
age int
salary float32

}
func(d doctor) speak(){

 fmt.Println(d.name,"can speak")
}

func(d doctor) getname()string{
return d.name
}

func(d doctor) getsalaryinfo()float32{
return d.salary
}
*/

func main(){

/*
var d doctor
d.name="tareq"
d.education="MBBS"
d.age=30
d.salary=50000.20

var d=doctor{name:"tareq",education:"MBBS",age:30,salary:20000,} 

 fmt.Println(d.name,d.education,d.age)
 d.speak()

fmt.Println(d.getname())
fmt.Println(d.getsalaryinfo())

}
*/

//var e Email
//fmt.Println(e)

var tst test
tst.to="rafiahammed19992@gmail.com"
tst.from="cityinstitute1992@gmail.com"
tst.sub="test email"
tst.message="hello this is atest email"
tst.write(tst.to,tst.from ,tst.sub )



 
}