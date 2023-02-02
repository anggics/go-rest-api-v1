package customer

import (
	"errors"
	customer_repos "rest-api/domain/customer/repository"
	"rest-api/entity"
)

type serviceCustomerImpl struct {
	customerRepos customer_repos.CustomerRepository
}

var (
	outAll []entity.Customer
	out    *entity.Customer
)

func CustomerService(
	repoCustomer customer_repos.CustomerRepository,
) ServiceCustomer {
	return &serviceCustomerImpl{customerRepos: repoCustomer}
}

// GetAllCustomer implements ServiceCustomer
func (sc *serviceCustomerImpl) GetAllCustomer() ([]entity.Customer, error) {
	result, err := sc.customerRepos.All()
	if err != nil {
		return outAll, err
	}
	outAll = result
	return outAll, err
}

// GetCustomerById implements ServiceCustomer
func (sc *serviceCustomerImpl) GetCustomerById(in int) (entity.Customer, error) {

	result, err := sc.customerRepos.FindById(in)
	if err != nil {
		return *out, err
	}
	out = &result
	return *out, err
}

// SaveCustomer implements ServiceCustomer
func (sc *serviceCustomerImpl) SaveCustomer(in entity.Customer) error {
	result, err := sc.customerRepos.FindByName(in.FirstName)
	if err != nil {
		return errors.New("internal server error")
	}
	if len(result.FirstName) > 0 || len(result.LastName) > 0 {
		return errors.New("data sudah ada")
	}

	customers := entity.Customer{
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}
	err = sc.customerRepos.SaveOrUpdate(customers)
	if err != nil {
		return err
	}

	return err
}

// UpdateCustomerById implements ServiceCustomer
func (sc *serviceCustomerImpl) UpdateCustomerById(id int, in entity.Customer) error {
	_, err := sc.customerRepos.FindById(id)
	if err != nil {
		return errors.New("internal server error")
	}

	err = sc.customerRepos.SaveOrUpdate(in)
	if err != nil {
		return err
	}
	return err
}

func (sc *serviceCustomerImpl) DeleteCustomerById(id int) error {

	err := sc.customerRepos.Delete(id)
	if err != nil {
		return err
	}

	return err
}
