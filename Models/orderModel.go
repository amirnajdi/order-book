package OrderModel

import (
	"time"

	database "github.com/amirnajdi/order-book/Database"
	"gorm.io/gorm"
)

type list struct { 
    BUY string
    SELL string
}

// Enum for public use
var SIDE = &list{ 
    BUY: "BUY",
    SELL: "SELL",
}

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

func (order *Order) Insert() error {
	var connection *gorm.DB = database.GetConnectionInstance()
	connection.Table("orders").Omit("CreatedAt", "Uuid").Create(order)
	if connection.Error != nil {
		return connection.Error
	}

	return nil
}

func GetAllOrders(symbol string, limit int) ([]Order, error) {
	var connection *gorm.DB = database.GetConnectionInstance()
	var orders []Order
	connection.Table(table).Select("id", "price", "amount", "side").Where("symbol = ?", symbol).Order("id desc").Limit(limit).Find(&orders)
	if connection.Error != nil {
		return []Order{}, nil
	}

	return orders, nil
}
