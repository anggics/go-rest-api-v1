package server

import "rest-api/interface/container"

type handler struct {
	customerHandler *customerHandler
}

func setupHandler(container *container.Container) *handler {
	customerHandlers := newCustomerHandler(container.CustomerService)
	return &handler{customerHandler: customerHandlers}
}
