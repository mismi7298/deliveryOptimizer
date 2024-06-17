package restaurants

import "deliveryOptimzer/pkg/model"

type RestaurantClient struct {
	restaurants []model.Restaurant
}

func NewRestaurantClient() *RestaurantClient {
	return &RestaurantClient{
		restaurants: make([]model.Restaurant, 0),
	}
}

func (r *RestaurantClient) AddRestaurant(restaurant model.Restaurant) {
	r.restaurants = append(r.restaurants, restaurant)
}

func (r *RestaurantClient) GetRestaurantById(id string) model.Restaurant {

	for i := range r.restaurants {
		if r.restaurants[i].Id == id {
			return r.restaurants[i]
		}
	}
	return model.Restaurant{}
}

func (r *RestaurantClient) GetRestaurantByLocationId(id string) model.Restaurant {

	for i := range r.restaurants {
		if r.restaurants[i].Location.Id == id {
			return r.restaurants[i]
		}
	}
	return model.Restaurant{}
}
