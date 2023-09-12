package core_cli

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli/v2"

	core_repositories "github.com/abusquets/ab-kyc/internal/core/adapters/spi/repositories"
	core_ports "github.com/abusquets/ab-kyc/internal/core/domain/ports/repositories"
	core_services "github.com/abusquets/ab-kyc/internal/core/domain/services"
	core_user_use_cases "github.com/abusquets/ab-kyc/internal/core/domain/use_cases/user"
	"github.com/abusquets/ab-kyc/internal/db"
	"github.com/abusquets/ab-kyc/pkg/console"
)

func createUser(db *sqlx.DB, email string, password string, IsAdmin bool) {
	userRepo := core_repositories.NewUserRepository(db)
	userService := core_services.NewUserService(userRepo)
	presenter := NewCreateUserPresenter()
	createUserUseCase := core_user_use_cases.NewCreateUserUseCase(userService, presenter)
	inData := core_ports.CreateUserInDTO{
		Email:    email,
		Password: &password,
		IsActive: true,
		IsAdmin:  IsAdmin,
	}
	error := createUserUseCase.Execute(inData)
	if error != nil {
		fmt.Println(error.AsMessage().Message)
	} else {
		res := presenter.Result()
		fmt.Printf("The User [%s] has been created. Uuid: %s\n", res.Email, res.Uuid)
	}
}

func Commands() []*cli.Command {
	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "email",
			Value:    "", // no default value
			Usage:    "Email for username",
			Required: true,
		},
	}

	commands := []*cli.Command{
		{
			Name:  "create-user",
			Usage: "Create a new user",
			Flags: myFlags,
			Action: func(cCtx *cli.Context) error {
				//password := console.StringPrompt("Pasword:")

				// return cli.NewExitError("it is not in the soup", 86)

				return nil
			},
		},
		{
			Name:  "create-admin",
			Usage: "Create a new admin",
			Flags: myFlags,
			Action: func(cCtx *cli.Context) error {
				dbManager := cCtx.App.Metadata["dbManager"].(db.DBManager)
				password := console.StringPrompt("Pasword:")
				createUser(dbManager.Database(), cCtx.String("email"), password, true)
				return nil
			},
		},
	}
	return commands
}
