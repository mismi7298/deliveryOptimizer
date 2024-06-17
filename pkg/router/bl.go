package router

import (
	"fmt"
	"log"
	"deliveryOptimzer/pkg/model"
	"deliveryOptimzer/pkg/order"
	ordermapper "deliveryOptimzer/pkg/order_mapper"
	"deliveryOptimzer/pkg/router/dl"
	"math"
)

type BL struct {
	dl            *dl.DL
	orderBL       *order.BL
	orderMapperBL *ordermapper.BL
}

func NewBL(orderBL *order.BL, orderMapperBL *ordermapper.BL, dl *dl.DL) *BL {
	return &BL{
		dl:            dl,
		orderBL:       orderBL,
		orderMapperBL: orderMapperBL,
	}
}

func (bl *BL) SetDeliveryRoute(de model.DeliveryExecutive) {

	orderIds, err := bl.orderMapperBL.GetOrdersForDeliveryExecutive(de)
	if err != nil {
		log.Fatalf("Error in fetching available delivery executive : %s", err)
	}
	fmt.Println("SetDeliveryRoute orderIds are ", orderIds)

	orders, err := bl.orderBL.GetOrders(orderIds)
	if err != nil {
		log.Fatalf("SetDeliveryRoute Error in fetching available delivery executive : %s", err)
	}

	route := bl.FindBestRoute(de, orders)

	fmt.Println("SetDeliveryRoute best route is ", route)

	// sets the delivery routes and saves it

	bl.dl.SaveDeliveryRoute(de, route)
}

func (bl *BL) FindBestRoute(de model.DeliveryExecutive, orders []model.Order) (locations []model.Location) {

	allPaths := findAllOrders(de, orders)
	var bestPath []model.Location
	bestDistance := 1e9
	for _, path := range allPaths {
		currDistance := 0.0
		prevLocation := de.Location
		for _, currentLocation := range path {
			prepTime := findRestaurantByLocationId(orders, currentLocation.Id).PrepTime
			currDistance += calculateDistance(prevLocation, currentLocation) + prepTimeToDistance(prepTime)
			prevLocation = currentLocation
		}
		if bestDistance > currDistance {
			bestDistance = currDistance
			bestPath = path
		}
	}

	return bestPath
}

func calculateDistance(loc1, loc2 model.Location) float64 {
	const R = 6371.0 // Earth radius in km

	lat1, lon1 := loc1.Latitude, loc1.Longitude
	lat2, lon2 := loc2.Latitude, loc2.Longitude

	dlat := toRadians(lat2 - lat1)
	dlon := toRadians(lon2 - lon1)

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(toRadians(lat1))*math.Cos(toRadians(lat2))*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return distance
}

func prepTimeToDistance(preptime float64) float64 {
	const avgVehicleSpeed = 20 // Will be dynamic

	return preptime * avgVehicleSpeed
}

func toRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func findRestaurantByLocationId(orders []model.Order, locationId string) model.Restaurant {
	for _, order := range orders {
		if order.Restaurant.Location.Id == locationId {
			return order.Restaurant
		}
	}
	return model.Restaurant{}
}
