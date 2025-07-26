package lib

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetUserIDFromContext(c *fiber.Ctx) (string, error) {
	localsID := c.Locals("id_user")
	idUser, ok := localsID.(string)
	if !ok || idUser == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: ID user tidak valid")
	}
	return idUser, nil
}

func RespondError(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"message_error": message,
	})
}

func RespondValidationError(c *fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"validation_errors": errors,
	})
}

func UploadImageFile(file *multipart.FileHeader, folderName string) (string, error) {
	// Validasi ukuran file (maks 2MB)
	if file.Size > 2*1024*1024 {
		return "", fmt.Errorf("ukuran gambar maksimal 2MB")
	}

	// Validasi ekstensi
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", fmt.Errorf("format gambar harus JPG/PNG")
	}

	// Buat nama file unik
	filename := fmt.Sprintf("invitation_%d%s", time.Now().UnixNano(), ext)

	// Ambil direktori kerja saat ini
	baseDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("gagal mendapatkan direktori kerja: %v", err)
	}

	// Tentukan path folder upload
	uploadDir := filepath.Join(baseDir, folderName)
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("gagal membuat folder upload: %v", err)
		}
	}

	// Path lengkap file
	savePath := filepath.Join(uploadDir, filename)

	// Simpan file
	if err := saveMultipartFile(file, savePath); err != nil {
		return "", fmt.Errorf("gagal menyimpan file: %v", err)
	}

	fmt.Println("Path:", savePath)
	fmt.Println("Filename:", filename)

	return filename, nil
}

// Helper: simpan file multipart
func saveMultipartFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = dst.ReadFrom(src)
	return err
}
