package data

import "database/sql"

func GetUserDetails(db *sql.DB, uuid string) (*UserDetail, error) {
	userDetail := &UserDetail{}
	err := db.QueryRow("SELECT * FROM UserDetails WHERE uuid=?", uuid).Scan(&userDetail.Uuid, &userDetail.FirstName, &userDetail.LastName)
	if err != nil {
		return nil, err
	}
	return userDetail, nil
}
func (userDetail *UserDetail) Create(db *sql.DB) error {
	if userDetail.Uuid == "" {
		return &ErrMissingUuid{userDetail: userDetail}
	}
	_, err := db.Exec("INSERT INTO UserDetails (uuid, firstName, lastName) VALUES (?, ?, ?)", userDetail.Uuid, userDetail.FirstName, userDetail.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (userDetail *UserDetail) UpdateFirstName(db *sql.DB) error {
	if userDetail.FirstName == "" {
		return &ErrMissingFirstName{userDetail: userDetail}
	}
	_, err := db.Exec("UPDATE UserDetails SET firstName=? WHERE uuid=?", userDetail.FirstName, userDetail.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (userDetail *UserDetail) UpdateLastName(db *sql.DB) error {
	if userDetail.LastName == "" {
		return &ErrMissingLastName{userDetail: userDetail}
	}
	_, err := db.Exec("UPDATE UserDetails SET lastName=? WHERE uuid=?", userDetail.LastName, userDetail.Uuid)
	if err != nil {
		return err
	}
	return nil
}
