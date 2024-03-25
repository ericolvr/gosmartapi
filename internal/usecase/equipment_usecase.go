package usecase

import (
	"fmt"
	"net/http"

	"github.com/ericolvr/goapi/internal/domain"
	"github.com/ericolvr/goapi/internal/repository"
)

type EquipmentUsecase interface {
	CreateEquipment(equipment *domain.Equipment) error
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
