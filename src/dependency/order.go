package dependency

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	ID                uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	DeliveryLatitude  float64 `json:"delivery_latitude"`
	DeliveryLongitude float64 `json:"delivery_longitude"`
	ClientID          uint64  `json:"client_id"`
	DropOffStart      uint64  `json:"drop_off_start"`
	DropOffEnd        uint64  `json:"drop_off_end"`
	RegionCode        string  `json:"region_code"`
	RouteID           uint64  `json:"route_id"`
}

var orderDB *gorm.DB

func initOrderDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/your_db_name?parseTime=true"
	var err error
	orderDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate schema
	err = orderDB.AutoMigrate(&Order{})
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}
}

func GetOrdersByIDs(ids []uint64) ([]Order, error) {
	var orders []Order
	if err := orderDB.Where("id IN ?", ids).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrdersByRouteID(routeID uint64) ([]Order, error) {
	var orders []Order
	if err := orderDB.Where("route_id = ?", routeID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(orderID uint64) (*Order, error) {
	var order Order
	if err := orderDB.First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func CreateOrder(o *Order) error {
	return orderDB.Create(o).Error
}

func UpdateOrder(o *Order) error {
	return orderDB.Save(o).Error
}

func DeleteOrder(orderID uint64) error {
	return orderDB.Delete(&Order{}, orderID).Error
}
