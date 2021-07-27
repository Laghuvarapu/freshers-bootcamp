package Models
import (


	"freshers-bootcamp/day4/Config"

)



//CreateUser ... Insert New data
func CreateProduct(user *Product) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(user *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetProducts(user *[]Product) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
