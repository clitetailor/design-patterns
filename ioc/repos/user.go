package repos

import "design-patterns/ioc/entities"

func NewUserRepo() IUserRepo {
	return &UserRepo{
		users: []*entities.User{
			{
				Name: "Tim Carousel",
				Age:  24,
			},
			{
				Name: "Lia",
				Age:  18,
			},
		},
	}
}

type IUserRepo interface {
	AddUser(user *entities.User)
	GetUsers() []*entities.User
}

type UserRepo struct {
	users []*entities.User
}

func (repo *UserRepo) AddUser(user *entities.User) {
	repo.users = append(repo.users, user)
}

func (repo *UserRepo) GetUsers() []*entities.User {
	return repo.users
}
