// internal/repository/user_repository.go\npackage repository\n\n// Define your repository methods here\n
// services/user-service/internal/repository/user_repository.go
package repository

import (
	"errors"
	"sync"

	"github.com/Himanshu4432/ecommerce-microservices/services/user-service/internal/models"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	SaveUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id string) (models.User, error)
}

// InMemoryUserRepository is an in-memory implementation of UserRepository
type InMemoryUserRepository struct {
	users map[string]models.User
	mu    sync.RWMutex
}

// NewUserRepository creates a new InMemoryUserRepository
func NewUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]models.User),
	}
}

// SaveUser saves a new user
func (repo *InMemoryUserRepository) SaveUser(user models.User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	// Check if user already exists
	for _, u := range repo.users {
		if u.Email == user.Email {
			return errors.New("user with this email already exists")
		}
	}

	repo.users[user.ID] = user
	return nil
}

// GetUserByEmail retrieves a user by email
func (repo *InMemoryUserRepository) GetUserByEmail(email string) (models.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, user := range repo.users {
		if user.Email == email {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

// GetUserByID retrieves a user by ID
func (repo *InMemoryUserRepository) GetUserByID(id string) (models.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	user, exists := repo.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
