package orderprocessor

import (
	"log"

	"deliveryOptimzer/pkg/model"
	"deliveryOptimzer/pkg/order"
	ordermapper "deliveryOptimzer/pkg/order_mapper"
	"deliveryOptimzer/pkg/router"
)

type BL struct {
	orderBL       *order.BL
	orderMapperBL *ordermapper.BL
	routerBL      *router.BL
}

func NewBL(orderMapperBL *ordermapper.BL, routerBL *router.BL, orderBL *order.BL) *BL {
	return &BL{
		orderMapperBL: orderMapperBL,
		routerBL:      routerBL,
		orderBL:       orderBL,
	}
}

func (bl *BL) ProcessOrder(orderEvent model.OrderEvent) {

	order, err := bl.orderBL.GetOrder(orderEvent.OrderId)
	if err != nil {
		log.Fatalf("ProcessOrder Error in fetching order details : %s", err)
	}
	de, err := bl.orderMapperBL.OrderDeliveryExecutiveMapper(order)

	if err != nil {
		log.Fatalf("ProcessOrder Error in mapping available delivery executive : %s", err)
	}
	bl.routerBL.SetDeliveryRoute(de)
}
