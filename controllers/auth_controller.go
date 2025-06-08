package controllers

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/akhmad-ardi/be-evowish/database"
	"github.com/akhmad-ardi/be-evowish/models"
	"github.com/akhmad-ardi/be-evowish/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("SECRET_JWT_KEY") // bisa juga pakai dari os.Getenv()

// @Summary Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body validations.RegisterRequest true "Register Request"
// @Success 201 {object} validations.RegisterResponse
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var input validations.RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorsMap := make(map[string]string)
			for _, fe := range ve {
				errorsMap[fe.Field()] = validations.ValidationMessage(fe)
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorsMap})
			return
		}

		// fallback error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kata sandi dan konfirmasi kata sandi tidak cocok"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashed)}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pengguna gagal terdaftar"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pengguna berhasil terdaftar"})
}

// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body validations.LoginRequest true "Login Request data"
// @Success 200 {object} validations.LoginResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var input validations.LoginRequest
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
