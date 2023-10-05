package user

type UserRepository interface {
	GetUserByID(id string) (UserDetails, error)
}

type UserDetails struct {
	ID   string
	Name string
}

type UserService struct {
	Repo UserRepository
}

func (s *UserService) GetUserNameByID(id string) (string, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
