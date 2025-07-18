package dependency

type Route struct {
	ID                     uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderIds               string  `json:"order_ids"`
	Cost                   float64 `json:"cost"`
	Distance               float64 `json:"distance"`
	Time                   float64 `json:"time"`
	WarehouseAddrLongitude float64 `json:"warehouse_address_longitude"`
	WarehouseAddrLatitude  float64 `json:"warehouse_addr_latitude"`
}

func GetRouteByID(routeID uint64) (*Route, error) {
	var route Route
	if err := db.First(&route, routeID).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func CreateRoute(r *Route) error {
	return db.Create(r).Error
}

func DeleteRoute(routeID uint64) error {
	return db.Delete(&Route{}, routeID).Error
}
