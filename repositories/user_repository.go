package repositories

import "test-k-link-indonesia/models"

type UserRepository interface {
	Register(user models.User) (models.User, error)
	Login(username string) (models.User, error)
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).First(&user).Error

	return user, err
}

func (r *repository) Login(username string) (models.User, error) {
	var user models.User

	err := r.db.Where("username = ?", username).First(&user).Error

	return user, err
}
