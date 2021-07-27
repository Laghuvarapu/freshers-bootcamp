package Controllers
import (
	"fmt"
	"net/http"

	"freshers-bootcamp/day4/Models"
	"github.com/gin-gonic/gin"
)



//CreateUser ... Create User
func CreateProduct(c *gin.Context) {
	var user Models.Product
	c.BindJSON(&user)
	err := Models.CreateProduct(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Product
	err := Models.GetProductByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetProducts(c *gin.Context) {
	var user []Models.Product
	err := Models.GetProducts(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateOrder(c *gin.Context) {
	var user Models.Order
	c.BindJSON(&user)
	err := Models.CreateOrder(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Order
	err := Models.GetOrderByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}