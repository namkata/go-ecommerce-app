package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// svc UserService
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of user service & inject to handler
	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
	}
	handler := UserHandler{
		svc: svc,
	}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoint
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)

	app.Post("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message": "please provide valid inputs",
			},
		)
	}

	token, err := h.svc.Signup(user)

	if err != nil {
		return ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message": "error on signup",
			},
		)
	}

	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": token,
		},
	)
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)

	if err != nil {
		return ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message": "please provide valid inputs",
			},
		)
	}
	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message": "please provide correct user id password",
			},
		)
	}

	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Login",
			"token":   token,
		},
	)
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "GetVerificationCode",
		},
	)
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Verify",
		},
	)
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "CreateProfile",
		},
	)
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "GetProfile",
		},
	)
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "AddToCart",
		},
	)
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "GetCart",
		},
	)
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "CreateOrder",
		},
	)
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "GetOrders",
		},
	)
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "GetOrder",
		},
	)
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "BecomeSeller",
		},
	)
}
