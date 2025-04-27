package monitoring

import (
	"encoding/json"
	"fmt"
	"log"
	"smart-home-server/config"
	"smart-home-server/models"
	"net/http"
	"time"
)

// FetchMedicationReminderFromOtherGroup 定时获取外部小组的宠物燥剂更换提醒数据并更新本地数据库
func FetchMedicationReminderFromOtherGroup() error {
	// 假设获取其他小组的宠物燥剂更换提醒接口URL
	url := "https://other-group-api.com/medication-reminders"

	// 发送HTTP请求，获取宠物燥剂更换提醒数据
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch medication reminders from other group: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应的宠物燥剂更换提醒数据
	var reminderData struct {
		Reminders []struct {
			PetID           uint      `json:"pet_id"`
			ReminderMessage string    `json:"reminder_message"`
			ReminderTime    time.Time `json:"reminder_time"`
		} `json:"reminders"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&reminderData); err != nil {
		return fmt.Errorf("failed to decode medication reminder data: %v", err)
	}

	// 1. 删除现有的提醒记录
	if err := config.DB.Where("1 = 1").Delete(&models.MedicationReminder{}).Error; err != nil {
		return fmt.Errorf("failed to delete existing medication reminders: %v", err)
	}

	// 2. 插入新的提醒记录
	for _, reminder := range reminderData.Reminders {
		newReminder := models.MedicationReminder{
			PetID:           reminder.PetID,
			ReminderMessage: reminder.ReminderMessage,
			ReminderTime:    reminder.ReminderTime,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		// 插入新的提醒记录
		if err := config.DB.Create(&newReminder).Error; err != nil {
			log.Printf("Failed to insert medication reminder for pet_id %d: %v", reminder.PetID, err)
		} else {
			log.Printf("Inserted new medication reminder for pet_id %d: %s", reminder.PetID, reminder.ReminderMessage)
		}
	}

	return nil
}
