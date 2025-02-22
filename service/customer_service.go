package service

import (
	"go-struct/entity"
	"go-struct/repository"
)

type CustomerService interface {
	CreateCustomer(customer *entity.Customer) error
	GetAllCustomers() ([]entity.Customer, error)
	GetCustomerByID(id string) (*entity.Customer, error)
	UpdateCustomer(id string, customer *entity.Customer) error
	DeleteCustomer(id string) error
}
type customerServiceImpl struct {
	Repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerServiceImpl{repo}
}

func (c *customerServiceImpl) CreateCustomer(customer *entity.Customer) error {
	return c.Repo.Create(customer)
}

func (c *customerServiceImpl) GetAllCustomers() ([]entity.Customer, error) {
	return c.Repo.GetAll()
}

func (c *customerServiceImpl) GetCustomerByID(id string) (*entity.Customer, error) {
	return c.Repo.GetByID(id)
}

func (c *customerServiceImpl) UpdateCustomer(id string, customer *entity.Customer) error {
	return c.Repo.Update(id, customer)
}

func (c *customerServiceImpl) DeleteCustomer(id string) error {
	return c.Repo.Delete(id)
}
