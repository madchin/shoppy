package data

import (
	"database/sql"
)

type phoneRepository struct {
	db *sql.DB
}

func NewPhoneRepository(db *sql.DB) *phoneRepository {
	return &phoneRepository{db: db}
}

func (repo *phoneRepository) GetPhones(uuid string) (Phones, error) {
	if uuid == "" {
		return nil, ErrMissingUuid{}
	}
	var phones Phones
	rows, err := repo.db.Query("SELECT id, number FROM Phones WHERE uuid=?", uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		phone := Phone{}
		if err := rows.Scan(&phone.Uuid, &phone.Id, &phone.Number); err != nil {
			return nil, err
		}
		phones = append(phones, phone)
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return phones, nil
}

func (repo *phoneRepository) Create(phone Phone) error {
	if phone.Uuid == "" {
		return ErrMissingUuid{}
	}
	_, err := repo.db.Exec("INSERT INTO Phones (uuid,number) VALUES (?,?)", phone.Uuid, phone.Number)
	if err != nil {
		return err
	}
	return nil
}

func (repo *phoneRepository) Update(phone Phone) error {
	if phone.Uuid == "" {
		return ErrMissingUuid{}
	}
	if phone.Number == "" {
		return ErrMissingPhoneNumber{phone}
	}
	_, err := repo.db.Exec("UPDATE Phones SET number=? WHERE uuid=?", phone.Number, phone.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (repo *phoneRepository) Delete(phone Phone) error {
	if phone.Uuid == "" {
		return ErrMissingUuid{}
	}
	if phone.Number == "" {
		return ErrMissingPhoneNumber{phone}
	}
	_, err := repo.db.Exec("DELETE FROM Phones WHERE uuid=? AND number=?", phone.Uuid, phone.Number)
	if err != nil {
		return err
	}
	return nil
}
