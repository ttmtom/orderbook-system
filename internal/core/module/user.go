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
	Repository *repository.UserRepository
	Service    *service.UserService
	Controller *controller.UserController
}

func NewUserModule(
	connection *gorm.DB,
	validator *validator.Validate,
	commonModule *CommonModule,
) *UserModule {
	userRepository := repository.NewUserRepository(connection)
	userService := service.NewUserService(userRepository, commonModule.Service)
	userController := controller.NewUserController(validator, userService)

	return &UserModule{
		Repository: userRepository,
		Service:    userService,
		Controller: userController,
	}
}

func (m *UserModule) InitUserRoute(e *echo.Echo) {
	user := e.Group("/users")
	{
		user.POST("", m.Controller.Register)
		user.POST("/login", m.Controller.Login)
		/* TODO add auth
		authUser := user.Group("/").Use(authMiddleware())
		*/
		authUser := user.Group("")
		{
			authUser.GET("/:idHash", m.Controller.GetUser)
		}
	}
}
