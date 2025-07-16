package internal

import (
	"routing/algorithm/src/dependency"
	"routing/algorithm/utils"
)

func FindMinimumCost(route *dependency.Route) (*dependency.Route, error) {
	inputOrderIds, err := utils.StringToUint64Slice(route.OrderIds)
	if err != nil {
		return nil, err
	}

	finalOrderIds, cost, distance, time := findMinimumCost(inputOrderIds, route.WarehouseAddrLatitude, route.WarehouseAddrLongitude)
	return &dependency.Route{
		OrderIds:               utils.Uint64SliceToString(finalOrderIds),
		Cost:                   cost,
		Distance:               distance,
		Time:                   time,
		WarehouseAddrLongitude: route.WarehouseAddrLongitude,
		WarehouseAddrLatitude:  route.WarehouseAddrLatitude,
	}, nil
}

func findMinimumCost(orderIds []uint64, wareHouseLatitude, warehouseLongitude float64) (finalOrderIds []uint64, cost, distance, time uint64) {
	return
}
