package Models
import (


	"freshers-bootcamp/day4/Config"

)



//CreateUser ... Insert New data
func CreateOrder(user *Order) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByID(user *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
