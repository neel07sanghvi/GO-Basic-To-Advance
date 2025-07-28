package storage

import (
	"time"

	"github.com/neel07sanghvi/crud-api/models"
)

type UserStorage struct {
	users  map[int]*models.User
	nextID int
}

func New() *UserStorage {
	storage := &UserStorage{
		users:  make(map[int]*models.User),
		nextID: 1,
	}

	storage.CreateUser("Hit Shiroya", "hit@gmail.com")
	storage.CreateUser("Gautam Jivrajani", "gautam@gmail.com")

	return storage
}

func (s *UserStorage) CreateUser(name, email string) *models.User {
	user := &models.User{
		ID:        s.nextID,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	s.users[s.nextID] = user
	s.nextID++

	return user
}

func (s *UserStorage) GetAllUsers() []*models.User {
	users := make([]*models.User, 0, len(s.users))

	for _, user := range s.users {
		users = append(users, user)
	}

	return users
}

func (s *UserStorage) GetUserByID(id int) (*models.User, bool) {
	user, exists := s.users[id]

	return user, exists
}

func (s *UserStorage) UpdateUser(id int, name, email string) (*models.User, bool) {
	user, exists := s.users[id]

	if !exists {
		return nil, false
	}

	user.Name = name
	user.Email = email

	return user, true
}

func (s *UserStorage) DeleteUser(id int) bool {
	_, exists := s.users[id]

	if exists {
		delete(s.users, id)
	}

	return exists
}
