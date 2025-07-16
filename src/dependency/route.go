package dependency

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Route struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderIds string `json:"order_ids"`
	Cost     uint64 `json:"cost"`
	Distance uint64 `json:"distance"`
	Time     uint64 `json:"time"`
}

var routeDB *gorm.DB

func InitRouteDB() error {
	dsn := "user:password@tcp(127.0.0.1:3306)/your_db_name?parseTime=true"
	var err error
	routeDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func GetRouteByID(routeID uint64) (*Route, error) {
	var route Route
	if err := routeDB.First(&route, routeID).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func CreateRoute(r *Route) error {
	return routeDB.Create(r).Error
}

func DeleteRoute(routeID uint64) error {
	return routeDB.Delete(&Route{}, routeID).Error
}
