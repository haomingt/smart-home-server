package models

import "time"

// FeedingPlan 表示宠物的喂食计划
type FeedingPlan struct {
    ID          uint      `gorm:"primaryKey"`
    PetID       uint      `json:"pet_id"`       // 关联宠物ID
    Amount      float64   `json:"amount"`       // 每次喂食的粮食量（克）
    Frequency   string    `json:"frequency"`    // 喂食频率，如每日、每周等
    CreatedAt   time.Time `json:"created_at"`   // 创建时间
    UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
}
