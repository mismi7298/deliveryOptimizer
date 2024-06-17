package factory

import (
	"deliveryOptimzer/pkg/customer"
	deliveryexecutive "deliveryOptimzer/pkg/delivery_executive"
	delExDL "deliveryOptimzer/pkg/delivery_executive/dl"
	"deliveryOptimzer/pkg/handler"
	"deliveryOptimzer/pkg/model"
	"deliveryOptimzer/pkg/order"
	orderDL "deliveryOptimzer/pkg/order/dl"
	ordermapper "deliveryOptimzer/pkg/order_mapper"
	orderMapperDL "deliveryOptimzer/pkg/order_mapper/dl"
	orderprocessor "deliveryOptimzer/pkg/order_processor"
	"deliveryOptimzer/pkg/restaurants"
	"deliveryOptimzer/pkg/router"
	routerDL "deliveryOptimzer/pkg/router/dl"
)

func FacotrySetup() {

	var restaurants = restaurants.NewRestaurantClient()
	var customers = customer.NewCustomerClient()
	var orderDL = orderDL.NewDL()
	var orderBL = order.NewBL(orderDL)

	var delExDL = delExDL.NewDL()
	var delExBL = deliveryexecutive.NewBL(delExDL)

	var orderMapperDL = orderMapperDL.NewDL()
	var orderMapperBL = ordermapper.NewBL(delExBL, orderMapperDL)

	NewInitClient(orderDL, delExDL, restaurants, customers).InitHandler()

	var routerDL = routerDL.NewDL()
	var routerBL = router.NewBL(orderBL, orderMapperBL, restaurants, routerDL)

	var orderProcessorBL = orderprocessor.NewBL(orderMapperBL, routerBL, orderBL)

	var orderqueue = make(chan model.OrderEvent, 10)
	var handler = handler.NewHandler(orderBL, orderProcessorBL, &orderqueue)

	handler.Handler()
}
