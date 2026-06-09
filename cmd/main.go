package main

import (
	"myapp/internal/config"
	"myapp/internal/server"
)

// type CustomValidator struct {
// 	validator *validator.Validate
// }

// // error handler
// func parseBindError(err error) string {
// 	var unmarshalErr *json.UnmarshalTypeError

// 	if errors.As(err, &unmarshalErr) {
// 		switch unmarshalErr.Field {
// 		case "name":
// 			return "Name must be a string"
// 		case "email":
// 			return "Email must be a string and required"
// 		case "password":
// 			return "Password must be required"
// 		}
// 	}

// 	return "Invalid request body"
// }

// func (cv *CustomValidator) Validate(i any) error {
// 	if err := cv.validator.Struct(i); err != nil {
// 		// Optionally, you could return the error to give each route more control over the status code
// 		return echo.ErrBadRequest.Wrap(err)
// 	}
// 	return nil
// }

// type User struct{
// 	gorm.Model
// 	Name string `json:"name" validate:"required"`
// 	Email string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required" max:"6"`
// }

func main() {

	// dsn := "host=localhost user=postgres password=12345 dbname=go-crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// env := config.LoadEnv()
	// dns := env.DatabaseURL
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println("server error:",err)
	// } else {
	// 	fmt.Println("Server connected successfully!")
	// }

	//? step -> 1
	env := config.LoadEnv()

	//? step -> 2
	db := config.ConnectDatabase(env)

	// e := echo.New()
	// e.Use(middleware.RequestLogger())

	// e.Validator = &CustomValidator{validator: validator.New()}

	// get all user api
	// e.GET("/",func(c *echo.Context) error {
	// 	return c.String(http.StatusOK,"Hello, world! In Go Language inside!")
	// })

	// Post User
	// users.RegisterRoute(e, db)

	// userRepository := users.NewRepository(db)
	// userService := users.NewService(userRepository)
	// userHandler := users.NewHandler(userService)

	// e.POST("/users", userHandler.CreateUser)

	/*
	e.POST("/user", func(c *echo.Context) error {
		// this created a new empty User object
		// that like var u User
		newUsers := new(User)
		
		if err := c.Bind(newUsers); err != nil{
			return c.JSON(http.StatusBadRequest,map[string]string{
				"error": parseBindError(err),
			})
		}

		if err := c.Validate(newUsers); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Invalid JSON data",
				"details": err.Error(),
			})
		}

		result := db.Create(&newUsers)

		if result.Error != nil{
			return c.JSON(http.StatusBadRequest, map[string]any{"error": result.Error.Error()})
		}

		return c.JSON(http.StatusCreated, newUsers)
	})
	*/
	// if err := e.Start(":"+env.Port); err != nil{
	// 	e.Logger.Error("failed to server start","error", err)
	// }

	//? step -> 3
	server.Start(db, env)
}