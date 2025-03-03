package main

import (
	"fmt"
	"log"
	"os"

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
	log.Fatal(app.Listen(":6969"))
}
