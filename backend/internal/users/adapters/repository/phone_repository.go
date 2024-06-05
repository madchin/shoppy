package adapters

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"database/sql"
)

type PhoneDTO struct {
	uuid   string
	number string
}

type PhonesDTO []PhoneDTO

type PhoneRepository struct {
	db *sql.DB
}

func NewPhoneRepository(db *sql.DB) user.PhoneRepository {
	return PhoneRepository{db}
}

func (p PhoneRepository) Get(userUuid string) (user.Phones, custom_error.ContextError) {
	rows, err := p.db.Query("SELECT * FROM Phones WHERE uuid=?", userUuid)

	if err == sql.ErrNoRows {
		return nil, custom_error.NewPersistenceError("retrieve user", "phone for specified user not found")
	}

	if err != nil {
		return nil, custom_error.UnknownPersistenceError("phone retrieve")
	}
	defer rows.Close()

	phonesDto := PhonesDTO{}
	for rows.Next() {
		phoneDto := PhoneDTO{}
		if err := rows.Scan(&phoneDto.uuid, &phoneDto.number); err != nil {
			return nil, custom_error.UnknownPersistenceError("phone retrieve")
		}

		phonesDto = append(phonesDto, phoneDto)
	}

	phones := mapPhonesDtoToDomainPhones(phonesDto)
	return phones, custom_error.ContextError{}
}

func (ur PhoneRepository) Create(userUuid string, phone user.Phone, validateFn func(user.Phone) []error) custom_error.ContextError {

	if errs := validateFn(phone); len(errs) > 0 {
		return custom_error.NewValidationErrors("user phone create", errs)
	}
	if _, err := ur.db.Exec("INSERT INTO Phones VALUES (?,?)", userUuid, phone.Number()); err != nil {
		return custom_error.UnknownPersistenceError("User phone create")
	}
	return custom_error.ContextError{}
}

func (ur PhoneRepository) Update(userUuid string, prevNumber string, phone user.Phone, validateFn func(user.Phone) []error) custom_error.ContextError {
	if errs := validateFn(phone); len(errs) > 0 {
		return custom_error.NewValidationErrors("user number update", errs)
	}

	if _, err := ur.Get(userUuid); err.Error() != "" {
		return err
	}

	if _, err := ur.db.Exec("UPDATE Phones SET number=? WHERE uuid=? AND number=?", phone.Number(), userUuid, prevNumber); err != nil {
		return custom_error.UnknownPersistenceError("update phone number")
	}

	return custom_error.ContextError{}
}

func (ur PhoneRepository) DeletePhone(userUuid string, phone user.Phone, deleteFn func(user.Phones) error) custom_error.ContextError {
	phones, err := ur.Get(userUuid)
	if err.Error() != "" {
		return err
	}

	if err := deleteFn(phones); err != nil {
		return custom_error.NewPersistenceError("user phone deletion", err.Error())
	}

	if _, err := ur.db.Exec("DELETE * FROM Phones WHERE uuid=? AND number=?", userUuid, phone.Number()); err != nil {
		return custom_error.UnknownPersistenceError("user phone delete")
	}

	return custom_error.ContextError{}
}

func (ur PhoneRepository) DeleteAll(userUuid string) custom_error.ContextError {
	if _, err := ur.Get(userUuid); err.Error() != "" {
		return err
	}

	if _, err := ur.db.Exec("DELETE * FROM Phones where uuid=?", userUuid); err != nil {
		return custom_error.UnknownPersistenceError("user all phones deletion")
	}

	return custom_error.ContextError{}
}

func mapPhonesDtoToDomainPhones(phonesDTO PhonesDTO) user.Phones {
	var phones user.Phones
	for _, phoneDto := range phonesDTO {
		phone := user.NewPhone(phoneDto.number)
		phones = append(phones, phone)
	}
	return phones
}
