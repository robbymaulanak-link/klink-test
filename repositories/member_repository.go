package repositories

import "test-k-link-indonesia/models"

type MemberRepository interface {
	CreateGender(gender models.Gender) (models.Gender, error)
	CreateLevel(level models.Level) (models.Level, error)
	CreateMember(member models.Member) (models.Member, error)
	GetGender(gender string) (models.Gender, error)
	GetLevel(level string) (models.Level, error)
	GetMemberById(id int) (models.Member, error)
	GetAllMembers() ([]models.Member, error)
	GetUser(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

func (r *repository) CreateGender(gender models.Gender) (models.Gender, error) {
	err := r.db.Create(&gender).First(&gender).Error

	return gender, err
}

func (r *repository) CreateLevel(level models.Level) (models.Level, error) {
	err := r.db.Create(&level).First(&level).Error

	return level, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).First(&user).Error

	return user, err
}

func (r *repository) CreateMember(member models.Member) (models.Member, error) {
	err := r.db.Create(&member).First(&member).Error

	return member, err
}

func (r *repository) GetGender(gender string) (models.Gender, error) {
	var data models.Gender

	err := r.db.Where("gender =?", gender).First(&data).Error

	return data, err
}

func (r *repository) GetLevel(level string) (models.Level, error) {
	var data models.Level

	err := r.db.Where("level =?", level).First(&data).Error

	return data, err
}

func (r *repository) GetMemberById(id int) (models.Member, error) {
	var data models.Member

	err := r.db.Preload("User").Preload("Gender").Preload("Level").Where("id =?", id).First(&data).Error

	return data, err
}

func (r *repository) GetAllMembers() ([]models.Member, error) {
	var data []models.Member

	err := r.db.Preload("User").Preload("Level").Preload("Gender").Find(&data).Error

	return data, err
}

func (r *repository) GetUser(id int) (models.User, error) {
	var data models.User

	err := r.db.Preload("Member").First(&data, "id = ?", id).Error

	return data, err
}
