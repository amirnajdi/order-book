package OrderBook

import (
	"time"

	database "github.com/amirnajdi/order-book/Database"
	"gorm.io/gorm"
)

type Order struct {
	ID        uint `gorm:"primarykey"`
	Side      string
	Symbol    string
	Amount    float64
	Price     float64
	Uuid      string
	CreatedAt time.Time
}

const table string = "orders"


func (order *Order) Insert() (error) {
	var connection *gorm.DB = database.GetConnectionInstance()
	connection.Table("orders").Omit("CreatedAt", "Uuid").Create(order)
	if connection.Error != nil {
		return connection.Error
	}

	return nil
}
