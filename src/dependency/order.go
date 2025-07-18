package dependency

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

func GetOrdersByIDs(ids []uint64) ([]Order, error) {
	var orders []Order
	if err := db.Where("id IN ?", ids).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrdersByRouteID(routeID uint64) ([]Order, error) {
	var orders []Order
	if err := db.Where("route_id = ?", routeID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(orderID uint64) (*Order, error) {
	var order Order
	if err := db.First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func CreateOrder(o *Order) error {
	return db.Create(o).Error
}

func UpdateOrder(o *Order) error {
	return db.Save(o).Error
}

func DeleteOrder(orderID uint64) error {
	return db.Delete(&Order{}, orderID).Error
}
