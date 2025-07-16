package dependency

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Route struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderIds string `json:"order_ids"`
	Cost     uint64 `json:"cost"`
	Distance uint64 `json:"distance"`
	Time     uint64 `json:"time"`
}

var routeDB *gorm.DB

func initRouteDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/your_db_name?parseTime=true"
	var err error
	routeDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := routeDB.AutoMigrate(&Route{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}

func GetRouteByID(routeID uint64) (*Route, error) {
	var route Route
	if err := orderDB.First(&route, routeID).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func CreateRoute(r *Route) error {
	return orderDB.Create(r).Error
}

func DeleteRoute(routeID uint64) error {
	return orderDB.Delete(&Route{}, routeID).Error
}
