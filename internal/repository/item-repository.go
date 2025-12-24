package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) Create(req dto.CreateOrUpdateItemRequest) (*model.Item, error) {
	item := model.Item{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	err := r.db.Create(&item).Error

	if err != nil {
		return &item, nil
	}

	return nil, err
}

func (r *ItemRepository) Update(req dto.CreateOrUpdateItemRequest, id uuid.UUID) (*model.Item, error) {
	item := model.Item{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	err := r.db.Where("id = ?", id).Updates(&item).Error

	if err != nil {
		return &item, nil
	}

	return nil, err
}

func (r *ItemRepository) Delete(id uuid.UUID) (*model.Item, error) {
	var item model.Item

	err := r.db.Where("id = ?", id).Delete(&item).Error

	if err != nil {
		return &item, nil
	}

	return nil, err
}

func (r *ItemRepository) FindById(id uuid.UUID) (*model.Item, error) {
	var item model.Item

	err := r.db.Take(&item, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) FindAll() (*[]model.Item, error) {
	var items []model.Item

	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}

	return &items, nil
}

func (r *ItemRepository) IncreaseStock(
	tx *gorm.DB,
	itemID uuid.UUID,
	qty int,
) error {
	return tx.Model(&model.Item{}).
		Where("id = ?", itemID).
		Update("stock", gorm.Expr("stock + ?", qty)).
		Error
}
