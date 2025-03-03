package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	Id        int    `gorm:"primaryKey" json:"id"` // the `json:"id"` this line says that from json input this id Should be mapped to Id like wise others
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

type User struct {
	Id        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

var db *gorm.DB

func ConnectDB() {
	// connection string
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Loading env")
	}
	cloud_connection_str := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(cloud_connection_str), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database : ", err)
	}
	fmt.Println("Connected to the database successfully")
	db.AutoMigrate(&Todo{})
	db.AutoMigrate(&User{})
}

func main() {
	ConnectDB()
	// todos := []Todo{}
	fmt.Println("Hello World")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello world"})
	})

	app.Post("/api/addtodo", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil { // this specific line c.BodyParser() does
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Enter the body"})
		}
		// todo.Id = len(todos) + 1
		// todos = append(todos, *todo)
		result := db.Create(todo)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to store in the database"})
		}
		return c.Status(200).JSON(fiber.Map{"message": "Added succesfully", "data": todo})
	})
	app.Get("/api/gettodos", func(c *fiber.Ctx) error {
		var todo []Todo
		db.Find(&todo)
		return c.Status(200).JSON(fiber.Map{"Todos": todo})
	})
	app.Post("/api/createuser", func(c *fiber.Ctx) error {
		user := &User{}
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err})
		}
		if user.Name == "" || user.Email == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Give the name or email"})
		}
		result := db.Create(user)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to store in the database"})
		}
		return c.Status(200).JSON(fiber.Map{"message": "Added succesfully", "data": user})
	})
	app.Get("/api/getusers", func(c *fiber.Ctx) error {
		var user []User
		db.Find(&user)
		return c.Status(200).JSON(fiber.Map{"Todos": user})
	})
	log.Fatal(app.Listen(":6969"))
}
