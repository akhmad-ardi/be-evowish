package main

import (
	"be-undangan-digital/config"
	"be-undangan-digital/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat .env file")
	}

	config.ConnectDatabase()

	err_migrate := config.DB.AutoMigrate(
		&models.User{},
		&models.Template{},
		&models.Invitation{},
		&models.InvitationLink{},
		&models.SharedSocial{},
		&models.GuestView{},
	)

	if err_migrate != nil {
		log.Fatal("Migaration failed: ", err_migrate)
	}

	log.Println("Migration success!!!")
}
