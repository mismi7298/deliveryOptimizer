package deliveryexecutive

import (
	"log"
	"deliveryOptimzer/pkg/delivery_executive/dl"
	"deliveryOptimzer/pkg/model"
)

type BL struct {
	dl *dl.DL
}

func NewBL(dl *dl.DL) *BL {
	return &BL{
		dl: dl,
	}
}

func (bl *BL) GetAvailableDeliveryExecutive() (de model.DeliveryExecutive, err error) {

	delEx, err := bl.dl.GetAvailableDeliveryExecutive()
	if err != nil {
		log.Fatalf("GetAvailableDeliveryExecutive failed to get delivery executives %s", err)
		return
	}

	return delEx[0], nil
}
