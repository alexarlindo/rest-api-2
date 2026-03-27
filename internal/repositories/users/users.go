package users
import "rest-api-2/internal/models"



type Users struct {
	users []models.User

}

func New () * Users {
	return &Users{}
}

func (u Users)  GetAll() []models.User{
	return u.users
}

func (u *Users) Add(newUser models.User) {
	u.users = append(u.users, newUser)

}







