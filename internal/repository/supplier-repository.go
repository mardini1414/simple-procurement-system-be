package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"gorm.io/gorm"
)

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

func (r *SupplierRepository) Create(req dto.CreateOrUpdateSupplierRequest) (*model.Supplier, error) {
	supplier := model.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	err := r.db.Create(&supplier).Error

	if err != nil {
		return &supplier, nil
	}

	return nil, err
}

func (r *SupplierRepository) Update(req dto.CreateOrUpdateSupplierRequest, id uuid.UUID) (*model.Supplier, error) {
	supplier := model.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	err := r.db.Where("id = ?", id).Updates(&supplier).Error

	if err != nil {
		return &supplier, nil
	}

	return nil, err
}

func (r *SupplierRepository) Delete(id uuid.UUID) (*model.Supplier, error) {
	var supplier model.Supplier

	err := r.db.Where("id = ?", id).Delete(&supplier).Error

	if err != nil {
		return &supplier, nil
	}

	return nil, err
}

func (r *SupplierRepository) FindById(id uuid.UUID) (*model.Supplier, error) {
	var supplier model.Supplier

	err := r.db.Take(&supplier, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &supplier, nil
}

func (r *SupplierRepository) FindByEmail(email string) (*model.Supplier, error) {
	var supplier model.Supplier

	err := r.db.Take(&supplier, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &supplier, nil
}

func (r *SupplierRepository) FindAll() (*[]model.Supplier, error) {
	var suppliers []model.Supplier

	if err := r.db.Find(&suppliers).Error; err != nil {
		return nil, err
	}

	return &suppliers, nil
}
