package query

import (
	"backend/internal/users/domain/user"
)

type RetrieveUser struct {
	Uuid string
}

type RetrieveUserHandler struct {
	retrieveUserUseCase retrieveUserUseCase
}
type retrieveUserUseCase struct {
	userRepository user.Repository
}

func NewRetrieveUserHandler(repository user.Repository) RetrieveUserHandler {
	return RetrieveUserHandler{retrieveUserUseCase{repository}}
}

func (rh RetrieveUserHandler) Execute(retrieveUser RetrieveUser) (user.User, error) {
	u, err := rh.retrieveUserUseCase.userRepository.Get(retrieveUser.Uuid, func(u user.User) (user.User, error) {
		err := u.Validate()
		if err != nil {
			return user.User{}, err
		}
		return u, nil
	})
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}
