package main
import (
	"fmt"

	"freshers-bootcamp/day3/problem2/Config"
	"freshers-bootcamp/day3/problem2/Models"
	"freshers-bootcamp/day3/problem2/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}

