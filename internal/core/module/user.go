package module

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"orderbook/internal/adapter/controller"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/service"
)

type UserModule struct {
	repository *repository.UserRepository
	service    *service.UserService
	controller *controller.UserController
}

func NewUserModule(connection *gorm.DB, validator *validator.Validate) *UserModule {
	userRepository := repository.NewUserRepository(connection)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(validator, userService)

	return &UserModule{
		repository: userRepository,
		service:    userService,
		controller: userController,
	}
}

func (m *UserModule) InitUserRoute(e *echo.Echo) {
	user := e.Group("/users")
	{
		user.POST("", m.controller.Register)
		/* TODO add auth
		authUser := user.Group("/").Use(authMiddleware())
		*/
		authUser := user.Group("")
		{
			authUser.GET("/:idHash", m.controller.GetUser)
		}
	}
}
