package customer

import "deliveryOptimzer/pkg/model"

type CustomerClient struct {
	customers []model.Customer
}

func NewCustomerClient() *CustomerClient {
	return &CustomerClient{
		customers: make([]model.Customer, 0),
	}
}

func (c *CustomerClient) AddCustomer(customer model.Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerClient) GetCustomerById(id string) model.Customer {

	for i := range c.customers {
		if id == c.customers[i].CustomerId {
			return c.customers[i]
		}
	}

	return model.Customer{}
}
