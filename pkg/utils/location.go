package utils

import (
	"strconv"

	"deliveryOptimzer/pkg/model"
)

func GetLocation(latitute, longitude string) (location model.Location, err error) {

	var (
		lat  float64
		long float64
	)

	lat, err = strconv.ParseFloat(latitute, 64)
	if err != nil {
		return
	}
	long, err = strconv.ParseFloat(longitude, 64)
	if err != nil {
		return
	}
	return model.Location{
		Latitude:  lat,
		Longitude: long,
	}, nil
}
