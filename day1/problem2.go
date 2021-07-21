package main

import "fmt"

type pay interface{
	salaryperMonth() int
}


type fulltimeEmp struct{
	basic int
	days int

}

type contractor struct{
	basic int
	days int
}

type freeLancer struct{
	basic int
	hours int

}
func (f fulltimeEmp) salaryperMonth() int {
	fmt.Println("Full-Time employee salary: " ,f.basic * f.days)
	return 1
}
func (f contractor) salaryperMonth() int {
	fmt.Println("Contractor salry:", f.basic* f.days)
	return 1
}
func (f freeLancer) salaryperMonth() int{
	fmt.Println("Free Lnacer Salary:", f.basic* f.hours)
	return 1
}

func main() {
	var p pay
	p= fulltimeEmp{500,28}
	p.salaryperMonth()
	p = contractor{100,28}
	p.salaryperMonth()
	p = freeLancer{10,15}
	p.salaryperMonth()





}
