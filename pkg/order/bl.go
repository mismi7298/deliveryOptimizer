package order

import (
	"fmt"
	"log"
	"deliveryOptimzer/pkg/model"
	"deliveryOptimzer/pkg/order/dl"
)

type BL struct {
	dl *dl.DL
}

func NewBL(dl *dl.DL) *BL {
	return &BL{
		dl: dl,
	}
}

func (bl *BL) GetNewOrders() (order []model.Order, err error) {
	order, err = bl.dl.GetNewOrders()
	if err != nil {
		log.Fatalf("GetNewOrders Error in fetching orders %s", err)
	}
	return
}

func (bl *BL) GetOrders(orderIds []string) (orders []model.Order, err error) {
	orders, err = bl.dl.GetOrders(orderIds)
	if err != nil {
		log.Fatalf("GetOrders Error in fetching orders %s", err)
	}
	return
}

func (bl *BL) GetOrder(orderId string) (order model.Order, err error) {
	orders, err := bl.dl.GetOrders([]string{orderId})
	if err != nil {
		log.Fatalf("GetOrder Error in fetching orders %s", err)
	}

	if len(orders) < 1 {
		err = fmt.Errorf("GetOrder no order found for %s", orderId)
		return
	}
	return orders[0], nil
}
