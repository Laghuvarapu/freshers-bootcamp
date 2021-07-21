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
	fmt.Println("Full-Time employee salary: " )
	return f.basic* f.days
}
func (f contractor) salaryperMonth() int {
	fmt.Println("Contractor salry:")
	return f.basic* f.days
}
func (f freeLancer) salaryperMonth() int {
	fmt.Println("Free Lnacer Salary:")
	return f.basic* f.hours

}

func main() {
	var p pay
	p= fulltimeEmp{500,28}
	fmt.Println(p.salaryperMonth())
	p = contractor{100,28}
	fmt.Println(p.salaryperMonth())

	p = freeLancer{10,15}
	fmt.Println(p.salaryperMonth())






}
