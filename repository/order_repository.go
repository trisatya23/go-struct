package repository

import (
	"go-struct/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *entity.Order) error
	GetById(orderID string) (*entity.Order, error)
	GetAll() ([]entity.Order, error)
	Update(id string, order *entity.Order) error
	Delete(orderID string) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db}
}

func (o *orderRepositoryImpl) Create(order *entity.Order) error {
	return o.db.Create(order).Error
}

func (o *orderRepositoryImpl) GetById(id string) (*entity.Order, error) {
	var order entity.Order
	err := o.db.Preload("OrderItems").First(&order, "order_id = ?", id).Error
	return &order, err
}

func (o *orderRepositoryImpl) GetAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := o.db.Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func (o *orderRepositoryImpl) Update(id string, order *entity.Order) error {
	return o.db.Save(&order).Error
}

func (o *orderRepositoryImpl) Delete(orderID string) error {
	return o.db.Delete(&entity.Order{}, "order_id=?", orderID).Error
}
