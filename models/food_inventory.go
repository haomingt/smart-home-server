package models

import "time"

// PetFoodInventory 表示宠物的粮食库存
type PetFoodInventory struct {
	ID              uint      `json:"id" gorm:"primaryKey"`            // 主键，自动递增
	PetID           uint      `json:"pet_id"`                          // 关联宠物ID
	RemainingAmount float64   `json:"remaining_amount"`                // 剩余粮食量
	CreatedAt       time.Time `json:"created_at"`                      // 创建时间
	UpdatedAt       time.Time `json:"updated_at"`                      // 更新时间
}
