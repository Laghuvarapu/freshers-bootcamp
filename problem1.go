package main

import (
	"encoding/json"
	"fmt"
)



type Matrix struct{
	rows int
	columns int
	twoDM [][]int
}
func (d Matrix) getNumberofRows() int{
	return d.rows
}

func (d Matrix) setElement(rowIndex, columnIndex , reqNum int) {
	d.twoDM[rowIndex][columnIndex]=reqNum
}

func (d Matrix) getNumberofColumns() int{
	return d.columns
}

func (d Matrix) addMatrices(newM [][]int)[][]int{

	for i:=0; i<d.rows ; i++ {
		for j:=0 ; j<d.columns ; j++ {
			d.twoDM[i][j]=  newM[i][j]+ d.twoDM[i][j]
		}
	}
	return d.twoDM

}
func (d *Matrix) printMatrix(){
	matrix,err:= json.MarshalIndent(d,"","")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(matrix))
}



func main() {
        s := Matrix{rows:2, columns: 2,twoDM :[][]int{{1,2},{3,4}}}
        r := s.getNumberofRows()
        c := s.getNumberofColumns()

        fmt.Println(r)
        fmt.Println(c)
        s.setElement(1,1,6)
        newM := [][]int{{5,6}, {7,8}}
        fmt.Println(s.addMatrices(newM))
        s.printMatrix()












}
