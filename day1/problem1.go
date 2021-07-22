package main

import (
	"encoding/json"
	"fmt"
)
	type Matrix struct{
	Rows int "json:"rows"'
	 columns int
	TwoDM [][]int 'json:"data"'
}
func (d Matrix) getNumberofRows() int{
	return d.Rows
}

func (d Matrix) setElement(rowIndex, columnIndex , reqNum int) {
	d.TwoDM[rowIndex][columnIndex]=reqNum
}

func (d Matrix) getNumberofColumns() int{
	return d.Columns
}

func (d Matrix) addMatrices(newM [][]int)[][]int{

	for i:=0; i<d.Rows ; i++ {
		for j:=0 ; j<d.Columns ; j++ {
			d.TwoDM[i][j]=  newM[i][j]+ d.TwoDM[i][j]
		}
	}
	return d.TwoDM

}

func matrixToJson(d Matrix) error {
	matrix , err := json.MarshalIndent(d,""," ")
	if err != nil {
		return err
	}
	fmt.Println(string(matrix))
	return nil
}

func main() {
        s := Matrix{Rows:2, Columns: 2,TwoDM :[][]int{{1,2},{3,4}}}
        r := s.getNumberofRows()
        c := s.getNumberofColumns()

        fmt.Println(r)
        fmt.Println(c)
        s.setElement(1,1,6)

        newM := [][]int{{5,6}, {7,8}}
        fmt.Println(s.addMatrices(newM))

        err := matrixToJson(s)
	    if err != nil{
		fmt.Println("Error in converting matrix to json ",err)
	    }
}













}
