// shippy-user-service/repository.go

package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/seiji-thirdbridge/shippy-user-service/proto/user"
)

type repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
	GetByEmail(email string) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	var user *pb.User
	user.Email = email

	if err := repo.db.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
