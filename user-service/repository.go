package main

import (
	pb "github.com/jakebjorke/shipper/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

//Repository is the repository interface
type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

//UserRepository is the implementation of the Repository interface for users.
type UserRepository struct {
	db *gorm.DB
}

//GetAll user
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

//Get a user by ID
func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

//GetByEmailAndPassword gets a user by email and password
func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

//Create creates a new user
func (repo *UserRepository) Create(user *pb.User) error {
	return repo.db.Create(user).Error
}
