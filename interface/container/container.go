package container

import (
	customer_repos "rest-api/domain/customer/repository"
	"rest-api/shared/database"
	"rest-api/usecase/customer"

	"github.com/jinzhu/gorm"
)

type Container struct {
	CustomerService customer.ServiceCustomer
}

func SetUpContainer(tx *gorm.DB) Container {
	database.Migrate(tx)
	customerServices := customer.CustomerService(customer_repos.SetupCustomerRepo(tx))
	return Container{CustomerService: customerServices}
}
