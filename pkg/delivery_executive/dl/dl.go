package dl

import (
	"deliveryOptimzer/pkg/model"
)

type DL struct {
	deliveryExecutives []model.DeliveryExecutive
}

func NewDL() *DL {
	return &DL{
		deliveryExecutives: make([]model.DeliveryExecutive, 0),
	}
}

// read the deliveryexecutive.csv and return the available delivery executive
func (dl *DL) GetAvailableDeliveryExecutive() (de []model.DeliveryExecutive, err error) {

	// records, err := utils.ReadCSVFile(model.DeliveryExecutiveFilePath)
	// if err != nil {
	// 	log.Fatalf("GetAvailableDeliveryExecutive(): Error in getting delivery executive records for : %s", err)
	// 	return
	// }
	// var delEx []model.DeliveryExecutive
	// for i := range records {

	// 	var location, _ = utils.GetLocation(records[i][2], records[i][3])
	// 	delEx = append(delEx, model.DeliveryExecutive{
	// 		Id:       records[i][0],
	// 		Name:     records[i][1],
	// 		Location: location,
	// 	})
	// }
	return dl.deliveryExecutives, nil
}

func (dl *DL) AddDeliveryExecutive(de model.DeliveryExecutive) {
	dl.deliveryExecutives = append(dl.deliveryExecutives, de)
}
