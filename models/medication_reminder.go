package models

import "time"

// MedicationReminder 表示宠物的燥剂更换提醒信息
type MedicationReminder struct {
	ID             uint      `gorm:"primaryKey"`
	PetID          uint      `json:"pet_id"`         // 宠物ID
	ReminderMessage string   `json:"reminder_message"` // 提醒信息
	ReminderTime   time.Time `json:"reminder_time"`   // 提醒时间
	CreatedAt      time.Time `json:"created_at"`      // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
}
