package customer_repos

import (
	"errors"

	"rest-api/entity"

	"github.com/jinzhu/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func SetupCustomerRepo(db *gorm.DB) CustomerRepository {
	return &customerRepo{
		db: db,
	}
}

// All implements customer.CustomerRepository
func (c *customerRepo) All() (out []entity.Customer, err error) {
	err = c.db.Model(&entity.Customer{}).Find(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New("failed to get List")
	}
	return out, err
}

// Delete implements customer.CustomerRepository
func (c *customerRepo) Delete(in int) error {
	err := c.db.Where("customer_id = ?", in).Delete(entity.Customer{}).Error
	if err != nil {
		return errors.New("failed to delete data")
	}
	return err
}

// FindById implements customer.CustomerRepository
func (c *customerRepo) FindById(in int) (out entity.Customer, err error) {
	err = c.db.Model(&entity.Customer{}).Where("customer_id = ? ", in).First(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New("failed to get data")
	}
	return out, err
}

// FindByName implements customer.CustomerRepository
func (c *customerRepo) FindByName(in string) (out entity.Customer, err error) {
	err = c.db.Model(&entity.Customer{}).Where("first_name = ? ", in).First(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return out, nil
		}
		return out, errors.New("failed to get data")
	}
	return out, err
}

// SaveOrUpdate implements customer.CustomerRepository
func (c *customerRepo) SaveOrUpdate(in entity.Customer) error {
	err := c.db.Model(entity.Customer{}).Save(&in).Error
	if err != nil {
		return errors.New("failed to save or update data")
	}
	return err
}
