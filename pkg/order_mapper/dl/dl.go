package dl

import (
	"deliveryOptimzer/pkg/model"
)

type DL struct {
	orderMapper map[string][]string
}

func NewDL() *DL {

	return &DL{
		orderMapper: make(map[string][]string),
	}
}

// read the orders.csv file and return the orders
func (dl *DL) AssignOrderToDeliveryExecutive(order model.Order, de model.DeliveryExecutive) (err error) {

	dl.orderMapper[de.Id] = append(dl.orderMapper[de.Id], order.Id)

	return nil
}

// read the orders.csv file and return the orders
func (dl *DL) GetOrdersForDeliveryExecutive(deExId string) (orderIds []string, err error) {
	// records, err := utils.ReadCSVFile(model.OrderDeliveryExMappingFilePath)
	// if err != nil {
	// 	log.Fatalf("GetOrdersForDeliveryExecutive(): Error in getting order delivery executive mapping records for %s: %s", deExId, err)
	// 	return
	// }
	// for i := range records {
	// 	if records[i][0] == deExId {
	// 		orderIds = append(orderIds, records[i][1])
	// 	}
	// }
	return dl.orderMapper[deExId], nil
}
