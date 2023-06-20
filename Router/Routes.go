package router

import (
	"net/http"

	orderModel "github.com/amirnajdi/order-book/Models"
	"github.com/gin-gonic/gin"
)

type ordersRequest struct {
	Limit int `form:"limit,default=100" binding:"max=5000"`
}

func DefineRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})

	router.GET("/orders/:symbol", func(c *gin.Context) {

		symbole := c.Param("symbol")
		var ordersRequest ordersRequest
		if err := c.ShouldBindQuery(&ordersRequest); err != nil {
			if ordersRequest.Limit > 5000 {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"status":  "fail",
					"message": "Value of the limit field must be less than 5000",
				})
				return
			}
		}

		orders, err := orderModel.GetAllOrders(symbole, ordersRequest.Limit)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "fail",
				"message": err,
			})
			return
		}

		var bids, asks [][]float64

		for _, order := range orders {
			if order.Side == orderModel.SIDE.BUY {
				bids = append(bids, []float64{order.Price, order.Amount})
			} else {
				asks = append(asks, []float64{order.Price, order.Amount})
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status":        "success",
			"lastUpdatedId": orders[0].ID,
			"bids":          bids,
			"asks":          asks,
		})
	})

}
