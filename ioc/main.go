package main

import (
	"context"
	"design-patterns/ioc/controllers"
	"design-patterns/ioc/repos"
	"log"

	"go.uber.org/fx"
)

func Run(userController controllers.IUserController) {
	userController.AddAndPrint()
}

func main() {
	
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			repos.NewUserRepo,
			controllers.NewUserController,
		),
		fx.Invoke(
			Run,
		),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
