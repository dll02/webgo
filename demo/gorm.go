package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"column:name"`
	Birthday time.Time `gorm:"column:birthday"`
}

func main() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	user := User{
		Name: "user5",
	}
	birthdayStr := "1988-11-20"
	birthday, err := time.Parse("2006-01-02", birthdayStr)
	if err != nil {
		log.Fatal(err)
	}
	user.Birthday = birthday

	result := db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Printf("Inserted user with ID: %d\n", user.ID)
}
