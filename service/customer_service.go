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
type customerRepositoryImpl struct {
	Repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerRepositoryImpl{repo}
}

func (c *customerRepositoryImpl) CreateCustomer(customer *entity.Customer) error {
	return c.Repo.Create(customer)
}

func (c *customerRepositoryImpl) GetAllCustomers() ([]entity.Customer, error) {
	return c.Repo.GetAll()
}

func (c *customerRepositoryImpl) GetCustomerByID(id string) (*entity.Customer, error) {
	return c.Repo.GetByID(id)
}

func (c *customerRepositoryImpl) UpdateCustomer(id string, customer *entity.Customer) error {
	return c.Repo.Update(id, customer)
}

func (c *customerRepositoryImpl) DeleteCustomer(id string) error {
	return c.Repo.Delete(id)
}
