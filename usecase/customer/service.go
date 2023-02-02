package customer

import (
	"rest-api/entity"
)

type ServiceCustomer interface {
	GetAllCustomer() ([]entity.Customer, error)
	GetCustomerById(in int) (entity.Customer, error)
	SaveCustomer(in entity.Customer) error
	DeleteCustomerById(in int) error
	UpdateCustomerById(id int, in entity.Customer) error

}
