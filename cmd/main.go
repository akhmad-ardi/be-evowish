package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat .env file")
	}

	if len(os.Args) < 2 {
		log.Fatal("❌ Perintah tidak ditemukan. Gunakan: migrate | drop")
	}

	cmd := os.Args[1]

	switch cmd {
	case "migrate":
		MigrateTables()
	case "drop":
		DropTables()
	default:
		log.Fatalf("❌ Perintah tidak dikenali: %s\nGunakan: migrate | drop", cmd)
	}
}
