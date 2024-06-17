package dl

import (
	"deliveryOptimzer/pkg/model"
)

type DL struct {
	orders []model.Order
}

func NewDL() *DL {
	return &DL{
		orders: make([]model.Order, 0),
	}
}

// read the orders.csv file and return the orders
func (dl *DL) GetNewOrders() ([]model.Order, error) {
	// records, err := utils.ReadCSVFile(model.OrdersFilePath)
	// if err != nil {
	// 	log.Fatalf("GetNewOrders: Error in getting GetNewOrders records for : %s", err)
	// 	return
	// }
	// for i := range records {
	// 	orders = append(orders, model.Order{
	// 		Id: records[i][0],
	// 	})
	// }
	// if len(orders) ==0 {
	// 	return
	// }
	return dl.orders, nil
}

func (dl *DL) GetOrders(orderIds []string) (res []model.Order, err error) {
	// // read the orders.csv file and return the orders
	// records, err := utils.ReadCSVFile(model.OrdersFilePath)
	// if err != nil {
	// 	log.Fatalf("GetOrders: Error in getting order detail records for %s: %s", orderIds, err)
	// 	return
	// }
	// for i := range records {

	// 	startLocation, _ := utils.GetLocation(records[i][2], records[i][3])
	// 	endLocation, _ := utils.GetLocation(records[i][5], records[i][6])
	// 	prepTime, _ := strconv.ParseFloat(records[i][7], 64)
	// 	for j := range orderIds {

	// 		if records[i][0] == orderIds[j] {
	// 			orders = append(orders, model.Order{
	// 				Id: records[i][0],
	// 				Restaurant: model.Restaurant{
	// 					Name:     records[i][1],
	// 					PrepTime: prepTime,
	// 					Location: startLocation,
	// 				},
	// 				Customer: model.Customer{
	// 					CustomerId: records[i][4],
	// 					Location:   endLocation,
	// 				},
	// 			})
	// 		}
	// 	}
	// }

	for i := range dl.orders {
		for j := range orderIds {
			if dl.orders[i].Id == orderIds[j] {
				res = append(res, dl.orders[i])
			}
		}
	}
	return res, nil
}

func (dl *DL) AddOrder(o model.Order) {
	dl.orders = append(dl.orders, o)
}
