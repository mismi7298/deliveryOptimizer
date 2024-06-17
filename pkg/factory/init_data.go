package factory

import (
	"deliveryOptimzer/pkg/customer"
	delExDL "deliveryOptimzer/pkg/delivery_executive/dl"
	"deliveryOptimzer/pkg/model"
	orderDL "deliveryOptimzer/pkg/order/dl"
	"deliveryOptimzer/pkg/restaurants"
)

type initClient struct {
	orderDL *orderDL.DL
	delExDL *delExDL.DL

	restaurants *restaurants.RestaurantClient
	customers   *customer.CustomerClient
}

func NewInitClient(orderDL *orderDL.DL, delExDL *delExDL.DL, restaurants *restaurants.RestaurantClient, customers *customer.CustomerClient) *initClient {
	return &initClient{
		orderDL:     orderDL,
		delExDL:     delExDL,
		customers:   customers,
		restaurants: restaurants,
	}
}

func (i *initClient) InitHandler() {
	i.InitCustomers()
	i.InitRestaurants()
	i.InitOrders()
	i.InitDeliveryExecutives()
}

func (i *initClient) InitCustomers() {
	i.customers.AddCustomer(model.Customer{
		CustomerId: "customer1",
		Location: model.Location{
			Id:        "cust1loc",
			Latitude:  23,
			Longitude: 14,
		},
	})
	i.customers.AddCustomer(model.Customer{
		CustomerId: "customer2",
		Location: model.Location{
			Id:        "cust2loc",
			Latitude:  7,
			Longitude: 9,
		},
	})
}

func (i *initClient) InitRestaurants() {
	i.restaurants.AddRestaurant(model.Restaurant{
		Id:       "R1",
		Name:     "VegRest",
		PrepTime: 2,
		Location: model.Location{
			Id:        "R1",
			Latitude:  14,
			Longitude: 13,
		},
	})

	i.restaurants.AddRestaurant(model.Restaurant{
		Id:       "R2",
		Name:     "NonVegRest",
		PrepTime: 3,
		Location: model.Location{
			Id:        "R2",
			Latitude:  4,
			Longitude: 7,
		},
	})
}

func (i *initClient) InitOrders() {
	i.orderDL.AddOrder(model.Order{

		Id:         "order1",
		Restaurant: i.restaurants.GetRestaurantById("R1"),
		Customer:   i.customers.GetCustomerById("customer1"),
		PrepTime:   2,
	})
	i.orderDL.AddOrder(model.Order{

		Id:         "order2",
		Restaurant: i.restaurants.GetRestaurantById("R2"),
		Customer:   i.customers.GetCustomerById("customer2"),
		PrepTime:   3,
	})
}

func (i *initClient) InitDeliveryExecutives() {
	i.delExDL.AddDeliveryExecutive(model.DeliveryExecutive{
		Id:   "delEx1",
		Name: "Suresh",
		Location: model.Location{
			Id:        "delEx1Location",
			Latitude:  1,
			Longitude: 1,
		},
	})
}
