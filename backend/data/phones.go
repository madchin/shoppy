package data

import "database/sql"

func GetPhones(db *sql.DB, uuid string) (Phones, error) {
	if uuid == "" {
		return nil, &ErrMissingUuid{}
	}
	var phones = Phones{}
	rows, err := db.Query("SELECT id,number FROM Phones WHERE uuid=?", uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		phone := &Phone{}
		rows.Scan(phone.Id, phone.Number)
		phones = append(phones, phone)
	}
	return phones, nil
}
