package main

import (
	"log"

	"github.com/conley21p/AccelDatabase/internal/config"
	"github.com/conley21p/AccelDatabase/internal/database"
	"github.com/conley21p/AccelDatabase/internal/model"
	"github.com/conley21p/AccelDatabase/pkg/util"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseUrl)
	password, err := util.HashPassword("password")
	if err != nil {
		log.Fatalf("Error generating password: %v\n", err)
	}
	users := []model.User{
		{
			Username: "admin",
			Password: password,
		},
	}
	_, err = db.NamedExec(
		`insert into users (username, password)
		 values (:username, :password)`,
		users,
	)
	if err != nil {
		log.Fatalf("Error inserting users: %v\n", err)
	}
	log.Printf("Successfully inserted users: %v\n", users)
}
