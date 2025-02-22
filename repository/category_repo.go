package repository

import "go-struct/entity"

type CategoryRepository interface {
	Create(category entity.Category) (entity.Category, error)
	FindById(id uint64) (entity.Category, error)
	FindAll() ([]entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
	Delete(category entity.Category) (entity.Category, error)
}
type categoryRepositoryImpl struct {
	categoryRepository CategoryRepository
}

func (c *categoryRepositoryImpl) Create(category entity.Category) (entity.Category, error) {
	return c.categoryRepository.Create(category)
}
func (c *categoryRepositoryImpl) FindById(id uint64) (entity.Category, error) {
	return c.categoryRepository.FindById(id)
}
func (c *categoryRepositoryImpl) FindAll() ([]entity.Category, error) {
	return c.categoryRepository.FindAll()
}
func (c *categoryRepositoryImpl) Update(category entity.Category) (entity.Category, error) {
	return c.categoryRepository.Update(category)
}
func (c *categoryRepositoryImpl) Delete(category entity.Category) (entity.Category, error) {
	return c.categoryRepository.Delete(category)
}
