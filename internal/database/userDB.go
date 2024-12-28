package database

import "exporter/internal/models"

func (s *service) CreateUser(user *models.User) error {
	return s.gormDB.Create(user).Error
}

func (s *service) FindUser(user *models.User, email string) {
	s.gormDB.First(user, "email = ?", email)
}

func (s *service) ModifyUserPass(user *models.User, passwordHash string) {
	s.gormDB.Model(user).Where("userId = ?", user.UserId).Update("passwordHash", passwordHash)
}
