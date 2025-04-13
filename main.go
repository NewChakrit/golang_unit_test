package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"regexp"
)

var validate = validator.New()

// User struct with validator tags.
type User struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"fullname" validate:"required,fullname"`
	Age      int    `json:"age" validate:"required,numeric,min=1"`
}

// setup function initializes the Fiber app.
func setUp() *fiber.App {
	app := fiber.New()

	// Register the custom validator functin for 'fullname'
	validate.RegisterValidation("fullname", validateFullname)

	app.Post("/users", func(ctx *fiber.Ctx) error {
		user := new(User)

		if err := ctx.BodyParser(user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		if err := validate.Struct(user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(user)
	})
	return app
}

func main() {
	//fmt.Println(Add(2, 3))
	//fmt.Println((Factorial(2))) // 2
	//fmt.Println((Factorial(3))) //6

	//app := setUp()
	//Gorm()
	//app.Listen(":8080")
	Postgres()

}

func validateFullname(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(fl.Field().String())
}

func Add(a, b int) int {
	return a + b
}

func Factorial(n int) (result int) {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	return n * Factorial(n-1)
}

// Factorial = !5 =5*4*3*2*1
