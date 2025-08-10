package services

import (
	"be-undangan-digital/config"
	"be-undangan-digital/models"
	"be-undangan-digital/requests"
	"errors"
	"fmt"
	"reflect"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func CreateInvitationService(IdUser string, data *requests.CreateInvitationRequest) (*models.Invitation, error) {
	id_invitation, err := gonanoid.New(8)
	if err != nil {
		return nil, errors.New("id not generate")
	}

	// if data.DataInvitation["time"] != nil {
	// 	_, errParseTime := lib.ParseMapTimeAndReplace(data.DataInvitation, "time", []string{"15:04:05", "15:04"})
	// 	if errParseTime != nil {
	// 		return nil, errParseTime
	// 	}
	// }

	// if data.DataInvitation["date"] != nil {
	// 	_, errParseDate := lib.ParseMapTimeAndReplace(data.DataInvitation, "date", []string{"2006-01-02"})
	// 	if errParseDate != nil {
	// 		return nil, errParseDate
	// 	}
	// }

	// jsonDataInvitation, err := json.Marshal(data.DataInvitation)
	// if err != nil {
	// 	return nil, errors.New("data invitation is error")
	// }

	new_invitation := models.Invitation{
		IdInvitation: id_invitation,
		IdTemplate:   data.IdTemplate,
		IdUser:       IdUser,
		Name:         data.Name,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := config.DB.Create(new_invitation).Error; err != nil {
		return nil, err
	}

	return &new_invitation, nil
}

func GetInvitationsService(IdUser string) (*[]models.Invitation, error) {
	var invitations []models.Invitation

	err := config.DB.Preload("InvitationLink").Where("id_user = ?", IdUser).Find(&invitations).Error
	if err != nil {
		return nil, err
	}

	return &invitations, nil
}

func GetInvitationService(IdInvitation string) (*models.Invitation, error) {
	var invitation models.Invitation

	err := config.DB.Preload("InvitationLink").Where("id_invitation = ?", IdInvitation).First(&invitation).Error
	if err != nil {
		println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("data undangan tidak ditemukan")
		}
		return nil, err
	}

	return &invitation, nil
}

func UpdateInvitationService(IdInvittion string, updates map[string]interface{}) (*models.Invitation, error) {
	var invitation models.Invitation

	// Cari data dulu
	if err := config.DB.First(&invitation, "id_invitation = ?", IdInvittion).Error; err != nil {
		return nil, errors.New("undangan tidak ditemukan")
	}

	// Filter agar hanya field yang tidak kosong/nil
	cleanUpdates := make(map[string]interface{})
	for k, v := range updates {
		if v == nil {
			continue
		}
		// Kalau string kosong, skip
		if str, ok := v.(string); ok && str == "" {
			continue
		}
		// Kalau zero value untuk tipe lain, skip
		if reflect.ValueOf(v).IsZero() {
			continue
		}
		cleanUpdates[k] = v
	}

	// Update
	if err := config.DB.Model(&invitation).Updates(cleanUpdates).Error; err != nil {
		return nil, err
	}

	return &invitation, nil
}

func DeleteInvitationService(IdInvitation string) (*models.Invitation, error) {
	var invitation models.Invitation

	err := config.DB.Delete(&invitation, "id_invitation = ?", IdInvitation).Error
	if err != nil {
		return nil, err
	}

	return &invitation, nil
}
