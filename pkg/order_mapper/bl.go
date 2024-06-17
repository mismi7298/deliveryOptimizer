package ordermapper

import (
	"log"

	deliveryexecutive "deliveryOptimzer/pkg/delivery_executive"
	"deliveryOptimzer/pkg/model"
	"deliveryOptimzer/pkg/order_mapper/dl"
)

type BL struct {
	dl      *dl.DL
	delExBL *deliveryexecutive.BL
}

func NewBL(delExBL *deliveryexecutive.BL, dl *dl.DL) *BL {
	return &BL{
		dl:      dl,
		delExBL: delExBL,
	}
}

func (bl *BL) OrderDeliveryExecutiveMapper(order model.Order) (de model.DeliveryExecutive, err error) {

	// get the available delivery executive
	de, err = bl.delExBL.GetAvailableDeliveryExecutive()
	if err != nil {
		log.Fatalf("OrderDeliveryExecutiveMapper Error in fetching available delivery executive : %s", err)
	}

	// save the details in db
	err = bl.dl.AssignOrderToDeliveryExecutive(order, de)
	if err != nil {
		log.Fatalf("OrderDeliveryExecutiveMapper failed to get delivery executives %s", err)
	}
	return
}

func (bl *BL) GetOrdersForDeliveryExecutive(de model.DeliveryExecutive) (orderIds []string, err error) {

	orderIds, err = bl.dl.GetOrdersForDeliveryExecutive(de.Id)
	if err != nil {
		log.Fatalf("GetOrdersForDeliveryExecutive Error in fetching available delivery executive : %s", err)
	}
	return
}
