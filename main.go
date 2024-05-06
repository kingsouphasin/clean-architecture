package main

import (
	"fmt"
	router "king/cmd/router/user"
	model "king/internal/domain/model/user"
	repository "king/internal/domain/repository/user"
	handler "king/internal/handler/user"
	usecase "king/internal/usecase/user"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


const (
  host     = "localhost"  // or the Docker service name if running in another container
  port     = 3306         // default PostgreSQL port
  user     = "root"     // as defined in docker-compose.yml
  password = "password" // as defined in docker-compose.yml
  dbname   = "Golang_fiber" // as defined in docker-compose.yml
)

func main(){
	app := fiber.New();
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
		SlowThreshold:      time.Second,
		LogLevel:           logger.Info,
		Colorful:           true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:   newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{});
	fmt.Println(db)

	userRepo := repository.NewGormUserRepository(db);
	userUsecase := usecase.NewUserRepository(userRepo);
	userhandler := handler.NewUserHandler(userUsecase);

	router.SetupUserRoutes(app, *userhandler)

	// app.Post("/register", userhandler.RegisterUser);
	// app.Get("/order/:id", userhandler.GetByID);
	// app.Put("/order/:id", userhandler.Update);
	// app.Delete("/order/:id", userhandler.Deleted);

	app.Listen(":3000")

}