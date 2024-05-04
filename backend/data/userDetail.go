package data

import "database/sql"

type userDetailRepository struct {
	db *sql.DB
}

func NewUserDetailRepository(db *sql.DB) *userDetailRepository {
	return &userDetailRepository{db: db}
}

func (repo *userDetailRepository) GetUserDetails(uuid string) (UserDetail, error) {
	userDetail := UserDetail{}
	if uuid == "" {
		return UserDetail{}, ErrMissingUuid{}
	}
	err := repo.db.QueryRow("SELECT * FROM UserDetails WHERE uuid=?", uuid).Scan(&userDetail.Uuid, &userDetail.FirstName, &userDetail.LastName)
	if err != nil {
		return UserDetail{}, err
	}
	return userDetail, nil
}
func (repo *userDetailRepository) Create(userDetail UserDetail) error {
	if userDetail.Uuid == "" {
		return ErrMissingUuid{}
	}
	_, err := repo.db.Exec("INSERT INTO UserDetails (uuid, firstName, lastName) VALUES (?, ?, ?)", userDetail.Uuid, userDetail.FirstName, userDetail.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userDetailRepository) UpdateFirstName(userDetail UserDetail) error {
	if userDetail.Uuid == "" {
		return ErrMissingUuid{}
	}
	if userDetail.FirstName == "" {
		return ErrMissingFirstName{userDetail: userDetail}
	}
	_, err := repo.db.Exec("UPDATE UserDetails SET firstName=? WHERE uuid=?", userDetail.FirstName, userDetail.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userDetailRepository) UpdateLastName(userDetail UserDetail) error {
	if userDetail.Uuid == "" {
		return ErrMissingUuid{}
	}
	if userDetail.LastName == "" {
		return ErrMissingLastName{userDetail: userDetail}
	}
	_, err := repo.db.Exec("UPDATE UserDetails SET lastName=? WHERE uuid=?", userDetail.LastName, userDetail.Uuid)
	if err != nil {
		return err
	}
	return nil
}
