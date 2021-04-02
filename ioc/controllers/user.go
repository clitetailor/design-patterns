package controllers

import (
	"design-patterns/ioc/entities"
	"design-patterns/ioc/repos"
	"fmt"
)

func NewUserController(userRepo repos.IUserRepo) IUserController {
	return &UserController{
		userRepo: userRepo,
	}
}

type IUserController interface {
	AddAndPrint()
}

type UserController struct {
	userRepo repos.IUserRepo
}

func (controller *UserController) AddAndPrint() {
	controller.userRepo.AddUser(&entities.User{
		Name: "Gate Anderson",
		Age:  21,
	})
	controller.userRepo.AddUser(&entities.User{
		Name: "Sierra",
		Age:  16,
	})

	users := controller.userRepo.GetUsers()

	for _, user := range users {
		fmt.Println(user)
	}
}
