package internal

import "routing/algorithm/src/dependency"

func CreateNewOrder(order *dependency.Order) error {
	_, err := dependency.GetClientByID(order.ClientID)
	if err != nil {
		return err
	}

	err = dependency.CreateOrder(order)
	if err != nil {
		return err
	}

	return nil
}

func GetOrder(id uint64) (*dependency.Order, error) {
	order, err := dependency.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
