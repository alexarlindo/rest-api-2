package repositories

import "rest-api-2/internal/models"



type Repositories struct {
	User interface{
		GetAll() []models.User
		Add(newUser models.User) 
	}
}

func  New() *Repositories{
	return &Repositories{}
	
} 