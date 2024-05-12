package command

import "backend/internal/users/domain/user"

type registerUser struct {
	uuid string
	user user.User
}

type registerUserUseCase struct {
	userRepository user.Repository
}

func NewUseCaseRegisterUser(repository user.Repository) registerUserUseCase {
	return registerUserUseCase{userRepository: repository}
}

func (ru registerUserUseCase) Execute(cmd registerUser) error {
	return ru.userRepository.Create(cmd.uuid, cmd.user, func(u user.User) (user.User, error) {
		err := u.Validate()
		if err != nil {
			return user.User{}, err
		}
		return u, nil
	})
}
