package repository

import (
	"database/sql"
	"errors"

	"github.com/ericolvr/goapi/internal/domain"
)

type EquipmentRepository interface {
	Create(equipment *domain.Equipment) error
	GetByIdentifier(identifier string) (*domain.Equipment, error)
}

type mysqlEquipmentRepository struct {
	db *sql.DB
}

func NewMySQLEquipmentRepository(db *sql.DB) EquipmentRepository {
	return &mysqlEquipmentRepository{
		db: db,
	}
}

func (r *mysqlEquipmentRepository) Create(equipment *domain.Equipment) error {
	query := "INSERT INTO equipments (identifier, uniorg, code, used_code, " +
		"environment, system_name, schedule, users, completed) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(query,
		equipment.Identifier,
		equipment.Uniorg,
		equipment.Code,
		equipment.UsedCode,
		equipment.Environment,
		equipment.SystemName,
		equipment.Schedule,
		equipment.Users,
		equipment.Completed,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	equipment.ID = id
	return nil
}

func (r *mysqlEquipmentRepository) GetByIdentifier(identifier string) (*domain.Equipment, error) {
	query := "SELECT id, identifier FROM equipments WHERE identifier = ?"
	row := r.db.QueryRow(query, identifier)

	var equipment domain.Equipment
	err := row.Scan(&equipment.ID, &equipment.Identifier)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &equipment, nil
}
