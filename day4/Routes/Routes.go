package Routes
import (
	"freshers-bootcamp/day4/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetProducts)

		grp1.POST("products", Controllers.CreateProduct)
		grp1.POST("orders", Controllers.CreateOrder)
		grp1.GET("products/:id", Controllers.GetProductByID)
		grp1.GET("orders/:id", Controllers.GetOrderByID)

	}
	return r
}
