package users

import "myapp/internal/users/dto"

// IT is a set of rule
// Service level DTO use করবে (best practice)
type Service interface {
	CreateUser(req dto.CreateRequest) (*dto.Response, error)
}

// that is a object. service take Repository
type service struct {
	repo Repository
}

// constructer function
func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

// Reciver function
func (s *service) CreateUser(req dto.CreateRequest) (*dto.Response, error){

	user := User{
		Name: req.Name,
		Email: req.Email,
	}

	if err := user.hashPassword(req.Password); err != nil {
		return nil, err
	}

	if err :=s.repo.CreateUser(&user); err != nil {
		return nil, err
	}
	
	response := dto.Response {
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}