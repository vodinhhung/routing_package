package internal

import (
	"errors"
	"log"
	"routing/algorithm/src/dependency"
	"routing/algorithm/utils"
)

const (
	defaultWarehouseLatitude  = 0.0
	defaultWarehouseLongitude = 0.0
	defaultCost               = 0
	defaultDistance           = 0
	defaultTime               = 0
	defaultRouteID            = 0
)

func FindMinimumCost(route *dependency.Route) (*dependency.Route, error) {
	inputOrderIds, err := utils.StringToUint64Slice(route.OrderIds)
	if err != nil {
		return nil, err
	}

	finalOrderIds, cost, distance, time, err := findMinimumCost(inputOrderIds, route.WarehouseAddrLatitude, route.WarehouseAddrLongitude)
	if err != nil {
		return nil, err
	}

	return &dependency.Route{
		OrderIds:               utils.Uint64SliceToString(finalOrderIds),
		Cost:                   cost,
		Distance:               distance,
		Time:                   time,
		WarehouseAddrLongitude: route.WarehouseAddrLongitude,
		WarehouseAddrLatitude:  route.WarehouseAddrLatitude,
	}, nil
}

func findMinimumCost(orderIds []uint64, wareHouseLatitude, warehouseLongitude float64) (
	finalOrderIds []uint64,
	cost, distance, time float64,
	err error,
) {
	var defaultOrderIds []uint64

	orders, err := dependency.GetOrdersByIDs(orderIds)
	if err != nil {
		return defaultOrderIds, defaultCost, defaultDistance, defaultTime, err
	}

	if len(orders) == 0 || len(orders) != len(orderIds) {
		err = errors.New("invalid order IDs provided")
		return defaultOrderIds, defaultCost, defaultDistance, defaultTime, nil
	}

	startCoord := Coord{
		Lat: wareHouseLatitude,
		Lng: warehouseLongitude,
	}

	regions, orderInRegion, finalOrders, distance := planOptimalRoute(startCoord, orders)
	log.Printf("Distance: %f\n", distance)
	log.Printf("Orders: %v\n", finalOrders)
	log.Printf("Regions: %v\n", regions)
	log.Printf("Order in region: %v\n", orderInRegion)

	mileDistance := distance * 0.621371
	averageSpeed := 30.0 // Average speed in miles per hour
	time = mileDistance / averageSpeed
	cost = mileDistance*18.75 + time*0.3

	finalOrderIds = make([]uint64, len(finalOrders))
	for i, order := range finalOrders {
		finalOrderIds[i] = order.ID
	}

	return finalOrderIds, cost, distance, time, nil
}
