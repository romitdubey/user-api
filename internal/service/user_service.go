package service

import (
	"time"

	"github.com/romitdubey1/user-api/internal/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Birthday not occurred yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *UserService) AddAge(user models.User) models.User {
	user.Age = s.CalculateAge(user.DOB)
	return user
}
