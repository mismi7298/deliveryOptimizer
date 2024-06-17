package dl

import "deliveryOptimzer/pkg/model"

type DL struct {
	routeMap map[string][]model.Location
}

func NewDL() *DL {

	return &DL{
		routeMap: make(map[string][]model.Location),
	}
}

func (dl *DL) SaveDeliveryRoute(de model.DeliveryExecutive, locations []model.Location) {

	dl.routeMap[de.Id] = locations
}
