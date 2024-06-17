package handler

import (
	"fmt"
	"log"
	"sync"

	"deliveryOptimzer/pkg/model"
	"deliveryOptimzer/pkg/order"
	orderprocessor "deliveryOptimzer/pkg/order_processor"
)

type Handler struct {
	orders           *chan model.OrderEvent
	channelOpen      bool
	wg               *sync.WaitGroup
	orderBL          *order.BL
	orderProcessorBL *orderprocessor.BL
}

func NewHandler(orderBL *order.BL, orderprocessorBL *orderprocessor.BL, orderChan *chan model.OrderEvent) *Handler {
	return &Handler{
		orders:           orderChan,
		channelOpen:      true,
		wg:               &sync.WaitGroup{},
		orderBL:          orderBL,
		orderProcessorBL: orderprocessorBL,
	}
}

func (h *Handler) CloseChannel() {
	h.channelOpen = false

	close(*h.orders)
}

func (h *Handler) Handler() {

	go h.FetchNewOrders()
	h.wg.Add(1)

	go h.OrderHandler()
	h.wg.Add(1)

	h.wg.Wait()
}

func (h *Handler) FetchNewOrders() {

	for h.channelOpen {
		var newOrders, err = h.orderBL.GetNewOrders()
		if err != nil {
			log.Fatalf("FetchNewOrders Error in fetching new orders: %s", err)
			h.CloseChannel()
		}

		for i := range newOrders {
			*h.orders <- model.OrderEvent{
				OrderId: newOrders[i].Id,
			}
		}

		h.CloseChannel()
	}
	h.wg.Done()
}

func (h *Handler) OrderHandler() {
	for order := range *h.orders {
		fmt.Println("OrderHandler processing this order", order)
		h.orderProcessorBL.ProcessOrder(order)
	}
	h.wg.Done()
}
