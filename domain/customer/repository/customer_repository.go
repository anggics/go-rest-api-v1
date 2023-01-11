package customer_repos

import "rest-api/entity"

type CustomerRepository interface {
	All() (out []entity.Customer, err error)
	FindById(in int) (out entity.Customer, err error)
	FindByName(in string) (out entity.Customer, err error)
	SaveOrUpdate(in entity.Customer) error
	Delete(in int) error
}
