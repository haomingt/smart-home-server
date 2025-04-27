package models

import "time"

// FeedingRecord 表示每次出粮的记录
type FeedingRecord struct {
    ID          uint      `gorm:"primaryKey"`
    PetID       uint      `json:"pet_id"`       // 关联宠物ID
    Amount      float64   `json:"amount"`       // 出粮量
    Timestamp   time.Time `json:"timestamp"`    // 出粮时间
    CreatedAt   time.Time `json:"created_at"`   // 创建时间
    UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
}
