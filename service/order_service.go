package service

import (
	"errors"
	"go-struct/entity"
	"go-struct/repository"
)

type OrderService interface {
	CreateOrder(order *entity.Order) error
	GetOrderByID(id string) (*entity.Order, error)
	GetAllOrders() ([]entity.Order, error)
	UpdateOrder(id string, order *entity.Order) error
	DeleteOrder(id string) error
}

type orderServiceImpl struct {
	repo repository.OrderRepository
}

func (o *orderServiceImpl) CreateOrder(order *entity.Order) error {
	if order == nil {
		return errors.New("Order tidak boleh kosong")
	}
	if order.TotalAmount <= 0 {
		return errors.New("total amount harus lebih dari 0")
	}
	return o.repo.Create(order)
}

func (o *orderServiceImpl) GetOrderByID(id string) (*entity.Order, error) {
	return o.repo.GetById(id)
}

func (o *orderServiceImpl) GetAllOrders() ([]entity.Order, error) {
	return o.repo.GetAll()
}

func (o *orderServiceImpl) UpdateOrder(id string, order *entity.Order) error {
	return o.repo.Update(id, order)
}

func (o *orderServiceImpl) DeleteOrder(id string) error {
	return o.repo.Delete(id)
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderServiceImpl{repo}
}
