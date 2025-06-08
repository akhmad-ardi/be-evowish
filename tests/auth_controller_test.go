package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/akhmad-ardi/be-evowish/controllers"
	"github.com/akhmad-ardi/be-evowish/database"
	"github.com/akhmad-ardi/be-evowish/validations"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/auth/register", controllers.Register)
	return r
}

func TestRegister_Success(t *testing.T) {
	database.DB = SetupTestDB()
	router := setupRouter()

	body := validations.RegisterRequest{
		Name:            "Test User",
		Email:           "test@example.com",
		Password:        "secret123",
		ConfirmPassword: "secret123",
	}

	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Pengguna berhasil terdaftar")
}

func TestRegister_PasswordMismatch(t *testing.T) {
	database.DB = SetupTestDB()
	router := setupRouter()

	body := validations.RegisterRequest{
		Name:            "Test User",
		Email:           "test@example.com",
		Password:        "secret123",
		ConfirmPassword: "wrong123",
	}

	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Kata sandi dan konfirmasi kata sandi tidak cocok")
}
