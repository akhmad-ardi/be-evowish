package main

import (
	"log"

	"be-undangan-digital/config"
	"be-undangan-digital/models"

	"github.com/joho/godotenv"
)

func DropTables() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat .env file")
	}

	// Connect DB
	config.ConnectDatabase()

	// Drop tables
	if err := config.DB.Migrator().DropTable(
		&models.User{},
		&models.Template{},
		&models.Invitation{},
		&models.InvitationLink{},
		&models.SharedSocial{},
		&models.GuestView{},
	); err != nil {
		log.Fatal("Gagal drop tabel: ", err)
	}

	log.Println("Semua tabel berhasil dihapus!")
}
