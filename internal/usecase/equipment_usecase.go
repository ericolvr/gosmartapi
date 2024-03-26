package usecase

import (
	"fmt"
	"net/http"

	"github.com/ericolvr/goapi/internal/domain"
	"github.com/ericolvr/goapi/internal/repository"
)

type EquipmentUsecase interface {
	CreateEquipment(equipment *domain.Equipment) error
	GetEquipments() ([]*domain.Equipment, error)
	GetEquipmentByID(id int64) (*domain.Equipment, error)
	UpdateEquipment(equipment *domain.Equipment) error
	DeleteEquipment(id int64) error
}

type equipmentUsecase struct {
	equipmentRepo repository.EquipmentRepository
}

type DuplicateIdentiferError struct {
	Identifier string
	StatusCode int
}

func (e *DuplicateIdentiferError) Error() string {
	return fmt.Sprintf("equipment with identifier %s already exists", e.Identifier)
}

func (e *DuplicateIdentiferError) Status() int {
	return e.StatusCode
}

func NewEquipmentUsecase(equipmentRepo repository.EquipmentRepository) EquipmentUsecase {
	return &equipmentUsecase{
		equipmentRepo: equipmentRepo,
	}
}

func (ue *equipmentUsecase) CreateEquipment(equipment *domain.Equipment) error {
	existingIdentifier, err := ue.equipmentRepo.GetByIdentifier(equipment.Identifier)

	if err != nil {
		return err
	}

	if existingIdentifier != nil {
		return &DuplicateIdentiferError{
			Identifier: equipment.Identifier,
			StatusCode: http.StatusConflict,
		}
	}
	return ue.equipmentRepo.Create(equipment)
}

func (ue *equipmentUsecase) GetEquipments() ([]*domain.Equipment, error) {
	return ue.equipmentRepo.GetEquipments()
}

func (ue *equipmentUsecase) GetEquipmentByID(id int64) (*domain.Equipment, error) {
	return ue.equipmentRepo.GetEquipmentByID(id)
}

func (ue *equipmentUsecase) UpdateEquipment(equipment *domain.Equipment) error {
	return ue.equipmentRepo.UpdateEquipment(equipment)
}

func (ue *equipmentUsecase) DeleteEquipment(id int64) error {
	return ue.equipmentRepo.DeleteEquipment(id)
}
