package repository

import "go-struct/entity"

type CustomerRepository interface {
	Create(customer entity.Customer) (entity.Customer, error)
	FindAll() ([]entity.Customer, error)
	FindById(id string) (entity.Customer, error)
	Update(customer entity.Customer) (entity.Customer, error)
	Delete(customer entity.Customer) error
}

type customerRepositoryImpl struct {
	customerRepository CustomerRepository
}

func (c *customerRepositoryImpl) Create(customer entity.Customer) (entity.Customer, error) {
	return c.customerRepository.Create(customer)
}
func (c *customerRepositoryImpl) FindAll() ([]entity.Customer, error) {
	return c.customerRepository.FindAll()
}
