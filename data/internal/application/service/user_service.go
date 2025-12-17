package service

import (
	"context"
	"data-service/internal/application/dto"
	"data-service/internal/domain/entity"
	"data-service/internal/domain/repository"
	"errors"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

// UserService сервис для работы с пользователями
type UserService struct {
	repo repository.UserRepository
}

// NewUserService создает новый экземпляр UserService
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser создает нового пользователя
func (s *UserService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	// Проверяем существование пользователя
	exists, err := s.repo.Exists(ctx, req.Username, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	user := &entity.User{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
		UserRole:       req.UserRole,
	}

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return s.toUserResponse(createdUser), nil
}

// GetUserByID получает пользователя по ID
func (s *UserService) GetUserByID(ctx context.Context, id int) (*dto.UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return s.toUserResponse(user), nil
}

// GetUserByUsername получает пользователя по username
func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*dto.UserResponse, error) {
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return s.toUserResponse(user), nil
}

// GetUserByEmail получает пользователя по email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return s.toUserResponse(user), nil
}

// ListUsers получает список всех пользователей
func (s *UserService) ListUsers(ctx context.Context, filter *entity.UserFilter) ([]*dto.UserResponse, error) {
	users, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, s.toUserResponse(user))
	}

	return responses, nil
}

// UpdateUser обновляет данные пользователя
func (s *UserService) UpdateUser(ctx context.Context, id int, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	// Обновляем только переданные поля
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.HashedPassword != nil {
		user.HashedPassword = *req.HashedPassword
	}
	if req.UserRole != nil {
		user.UserRole = *req.UserRole
	}

	updatedUser, err := s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return s.toUserResponse(updatedUser), nil
}

// DeleteUser удаляет пользователя по ID
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// CheckUserExists проверяет существование пользователя
func (s *UserService) CheckUserExists(ctx context.Context, username, email string) (bool, error) {
	return s.repo.Exists(ctx, username, email)
}

// ListStaff получает список всех сотрудников
func (s *UserService) ListStaff(ctx context.Context) ([]*dto.StaffResponse, error) {
	users, err := s.repo.ListStaff(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.StaffResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, s.toStaffResponse(user))
	}

	return responses, nil
}

// toUserResponse преобразует entity.User в dto.UserResponse
func (s *UserService) toUserResponse(user *entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		UserRole:       user.UserRole,
	}
}

// toStaffResponse преобразует entity.User в dto.StaffResponse
func (s *UserService) toStaffResponse(user *entity.User) *dto.StaffResponse {
	return &dto.StaffResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserRole:  user.UserRole,
	}
}

