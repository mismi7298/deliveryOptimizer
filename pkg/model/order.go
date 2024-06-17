package model

type Order struct {
	Id         string
	Restaurant Restaurant
	Customer   Customer
	PrepTime   int64
}

type OrderEvent struct {
	OrderId string
}
