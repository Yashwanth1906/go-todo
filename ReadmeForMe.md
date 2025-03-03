To start a go project we need to get the go.mod file which is like a package.json file:

1. go mod init github.com/Yashwanth1906/project-name

To start running a go file or go api folder

we need to write a function called main() in the main package

1. go run .\main.go


To create a httpserver easily we are using github.com/gofiber/fiber/v2 module This module is actually fcking great 
It as lot of methods exactly like express use this


TO have a nodemon feel like in express we will be using github.com/cosmtrek/air@latest 


2. To connect to postgress :

use gorm :

go get gorm.io/gorm  -> this will get the gorm orm
go get gorm.io/driver/postgres -> this will add the driver for the postgres

make a pointer to point to the gorm.DB

connect with the db using

gorm.Open(postgres.Open(db_url),&gorm.Config{}) -> this is the database connection thing

to migrate db use db.Automigrate(the struct you defined) -> the struct defined should contain a gorm:"primaryKey"