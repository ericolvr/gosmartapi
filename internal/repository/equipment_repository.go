package repository

import (
	"database/sql"
	"errors"

	"github.com/ericolvr/goapi/internal/domain"
)

type EquipmentRepository interface {
	Create(equipment *domain.Equipment) error
	GetEquipments() ([]*domain.Equipment, error)
	GetEquipmentByID(id int64) (*domain.Equipment, error)
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

func (r *mysqlEquipmentRepository) GetEquipments() ([]*domain.Equipment, error) {
	query := "SELECT id, identifier, uniorg, code, used_code, environment, system_name, schedule, users, completed, created_at FROM equipments"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []*domain.Equipment
	for rows.Next() {
		var equipment domain.Equipment
		err := rows.Scan(&equipment.ID, &equipment.Identifier, &equipment.Uniorg, &equipment.Code, &equipment.UsedCode, &equipment.Environment, &equipment.SystemName, &equipment.Schedule, &equipment.Users, &equipment.Completed, &equipment.CreatedAt)
		if err != nil {
			return nil, err
		}
		equipments = append(equipments, &equipment)
	}

	return equipments, nil
}

func (r *mysqlEquipmentRepository) GetEquipmentByID(id int64) (*domain.Equipment, error) {
	query := "SELECT id, identifier, uniorg, code, used_code, " +
		"environment, system_name, schedule, users, completed," +
		"created_at FROM equipments WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var equipment domain.Equipment
	err := row.Scan(
		&equipment.ID,
		&equipment.Identifier,
		&equipment.Uniorg,
		&equipment.Code,
		&equipment.UsedCode,
		&equipment.Environment,
		&equipment.SystemName,
		&equipment.Schedule,
		&equipment.Users,
		&equipment.Completed,
		&equipment.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &equipment, nil
}
