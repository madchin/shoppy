package data

import (
	"database/sql"
)

func GetPhones(db *sql.DB, uuid string) (Phones, error) {
	if uuid == "" {
		return nil, &ErrMissingUuid{}
	}
	var phones Phones
	rows, err := db.Query("SELECT id, number FROM Phones WHERE uuid=?", uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		phone := &Phone{}
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

func (p *Phone) Create(db *sql.DB) error {
	if p.Uuid == "" {
		return &ErrMissingUuid{}
	}
	_, err := db.Exec("INSERT INTO Phones (uuid,number) VALUES (?,?)", p.Uuid, p.Number)
	if err != nil {
		return err
	}
	return nil
}

func (p *Phone) Update(db *sql.DB) error {
	if p.Uuid == "" {
		return &ErrMissingUuid{}
	}
	if p.Number == "" {
		return &ErrMissingPhoneNumber{}
	}
	_, err := db.Exec("UPDATE Phones SET number=? WHERE uuid=?", p.Number, p.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (p *Phone) Delete(db *sql.DB) error {
	if p.Uuid == "" {
		return &ErrMissingUuid{}
	}
	if p.Number == "" {
		return &ErrMissingPhoneNumber{phone: p}
	}
	_, err := db.Exec("DELETE FROM Phones WHERE uuid=? AND number=?", p.Uuid, p.Number)
	if err != nil {
		return err
	}
	return nil
}
