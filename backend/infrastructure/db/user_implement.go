package db

import (
	"context"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	irepository "github.com/Temisaputra/warOnk/delivery/repository"
	"github.com/Temisaputra/warOnk/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	*TransactionRepository
}

func NewUserRepo(db *gorm.DB) irepository.UserRepository {
	return &UserRepository{
		TransactionRepository: NewTransactionRepo(db),
	}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]presenter.UserResponse, error) {
	var users []entity.Users
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	var presenters []presenter.UserResponse
	for _, user := range users {
		presenters = append(presenters, *user.ToPresenter())
	}
	return presenters, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user entity.Users) error {
	if err := r.Conn(ctx).WithContext(ctx).Model(&entity.Users{}).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (entity.Users, error) {
	var user entity.Users
	if err := r.Conn(ctx).WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user entity.Users) error {
	if err := r.Conn(ctx).WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	if err := r.Conn(ctx).WithContext(ctx).Delete(&entity.Users{}, id).Error; err != nil {
		return err
	}
	return nil
}
