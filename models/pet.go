package models

import "time"

// Pet 表示宠物的基本信息
type Pet struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `json:"name"`         // 宠物名字
    Type        string    `json:"type"`         // 宠物类型
    Age         int       `json:"age"`          // 宠物年龄
    Weight      float64   `json:"weight"`       // 宠物体重
    CreatedAt   time.Time `json:"created_at"`   // 创建时间
    UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
}
