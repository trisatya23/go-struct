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
	if len(order.OrderItems) == 0 {
		return errors.New("order harus memiliki minimal satu item")
	}
	if order.TotalAmount <= 0 {
		return errors.New("total amount harus lebih dari 0")
	}
	for _, item := range order.OrderItems {
		if item.ProductID == "" {
			return errors.New("setiap item harus memiliki product_id")
		}
	}
	return o.repo.Create(order)
}

func (o *orderServiceImpl) GetOrderByID(id string) (*entity.Order, error) {
	if id == "" {
		return nil, errors.New("order ID tidak boleh kosong")
	}
	order, err := o.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *orderServiceImpl) GetAllOrders() ([]entity.Order, error) {
	orders, err := o.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderServiceImpl) UpdateOrder(id string, order *entity.Order) error {
	if id == "" {
		return errors.New("order ID tidak boleh kosong")
	}
	if order == nil {
		return errors.New("order harus memiliki minimal satu item")
	}
	if order.TotalAmount <= 0 {
		return errors.New("total amount harus lebih dari 0")
	}

	existingOrder, err := o.repo.GetById(id)
	if err != nil {
		return errors.New("order tidak ditemukan")
	}
	existingOrder.CustomerID = order.CustomerID
	existingOrder.OrderItems = order.OrderItems
	existingOrder.TotalAmount = order.TotalAmount
	existingOrder.OrderItems = order.OrderItems

	return o.repo.Update(id, existingOrder)
}

func (o *orderServiceImpl) DeleteOrder(id string) error {
	if id == "" {
		return errors.New("order ID tidak boleh kosong")
	}
	_, err := o.repo.GetById(id)
	if err != nil {
		return errors.New("order tidak ditemukan")
	}
	return o.repo.Delete(id)
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderServiceImpl{repo}
}
