package repositories

import "rest-api-2/internal/models"
import "rest-api-2/internal/repositories/users"



type Repositories struct {
	User interface{
		GetAll() []models.User
		Add(newUser models.User) 
		EmailExists(email string ) bool
	}
}

func  New() *Repositories{
	return &Repositories{
		User: users.New(),
	}
	
} 