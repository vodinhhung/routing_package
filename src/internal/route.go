package internal

import (
	"routing/algorithm/src/dependency"
	"routing/algorithm/utils"
)

func CreateNewRoute(route *dependency.Route) error {
	inputOrderIds, err := utils.StringToUint64Slice(route.OrderIds)
	if err != nil {
		return err
	}
	finalOrderIds, cost, distance, time := findMinimumCost(inputOrderIds, route.WarehouseAddrLatitude, route.WarehouseAddrLongitude)
	route.OrderIds = utils.Uint64SliceToString(finalOrderIds)
	route.Cost = cost
	route.Distance = distance
	route.Time = time

	err = dependency.CreateRoute(route)
	if err != nil {
		return err
	}

	return nil
}

func GetRoute(id uint64) (*dependency.Route, error) {
	route, err := dependency.GetRouteByID(id)
	if err != nil {
		return nil, err
	}

	return route, nil
}
