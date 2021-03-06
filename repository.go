package main

import (
	pb "github.com/caquillo07/grpc-demo-user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type Users []*pb.User

type Repository interface {
	GetAll() (Users, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo UserRepository) GetAll() (Users, error) {
	var users Users
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepository) Create(user *pb.User) error {
	return repo.db.Create(user).Error
}

func (repo UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
