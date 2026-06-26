package users

import (
	"errors"
	"fmt"
	"myapp/internal/auth"
	"myapp/internal/users/dto"
)

var ErrInvalideCredentials = errors.New("User email not found")

// It is a set of rule
// Service level DTO use করবে (best practice)
type Service interface {
	CreateUser(req dto.CreateRequest) (*dto.Response, error)
}

// that is a object. service take Repository
type service struct {
	repo Repository
	jwtService auth.JWTService
}

// constructer function
func NewService(repo Repository, jwtService auth.JWTService) *service {
	return &service{
		repo: repo,
		jwtService: jwtService,
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

func (s *service) LoginUser(req dto.LoginRequest) (*dto.Response, error){
	// before login check user is exist in database
	// and database related work do in repository.go inside

	user, err := s.repo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, ErrInvalideCredentials
	}

	if err := user.CheckPassword(req.Password); err != nil {
		return nil, err
	}

	token, err := s.jwtService.GenerateToken(user.ID,user.Name,user.Email)

	if err != nil {
		return  nil, fmt.Errorf("token is not valied %w",err)
	}

	response := dto.Response {
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
		Token: token,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}