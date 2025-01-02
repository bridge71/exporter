package database

import "exporter/internal/models"

func (s *service) SaveUser(user *models.User) error {
	return s.gormDB.Save(user).Error
}

func (s *service) DeleteUser(user *models.User) error {
	return s.gormDB.Delete(user).Error
}

func (s *service) FindUser(user *[]models.User) {
	s.gormDB.Find(user)
}

func (s *service) FindUserByEmail(user *models.User, email string) {
	s.gormDB.First(user, "email = ?", email)
}

func (s *service) FindUserById(user *models.User, userId uint) {
	s.gormDB.First(user, "userId = ?", userId)
}
