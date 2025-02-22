package controller

import (
	"go-struct/entity"
	"go-struct/repository"
)

type CategoryController struct {
	repo repository.CategoryRepository
}

func (controller *CategoryController) Create(category entity.Category) (entity.Category, error) {
	return controller.repo.Create(category)
}

func (controller *CategoryController) Update(category entity.Category) (entity.Category, error) {
	return controller.repo.Update(category)
}

func (controller *CategoryController) Delete(category entity.Category) (entity.Category, error) {
	return controller.repo.Delete(category)
}

func (controller *CategoryController) GetById(id uint64) (entity.Category, error) {
	return controller.repo.FindById(id)
}

func (controller *CategoryController) GetAll() ([]entity.Category, error) {
	return controller.repo.FindAll()
}
