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

// FetchPetFoodInventoryFromOtherGroup 定时获取外部小组的每个宠物剩余粮食量数据并更新本地数据库
func FetchPetFoodInventoryFromOtherGroup() error {
	// 假设获取其他小组的每个宠物剩余粮食量接口URL
	url := "https://other-group-api.com/pet-food-inventory"

	// 发送HTTP请求，获取宠物剩余粮食量数据
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch pet food inventory from other group: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应的宠物剩余粮食量数据
	var inventoryData struct {
		Pets []struct {
			PetID         uint    `json:"pet_id"`
			RemainingAmount float64 `json:"remaining_amount"`
		} `json:"pets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&inventoryData); err != nil {
		return fmt.Errorf("failed to decode pet food inventory data: %v", err)
	}

	// 遍历每个宠物的剩余粮食量数据
	for _, petInventory := range inventoryData.Pets {
		// 检查数据库中是否已有该宠物的粮食量记录
		var existingInventory models.PetFoodInventory
		err := config.DB.Where("pet_id = ?", petInventory.PetID).First(&existingInventory).Error
		if err == nil {
			// 如果记录存在，则更新其剩余粮食量
			existingInventory.RemainingAmount = petInventory.RemainingAmount
			existingInventory.UpdatedAt = time.Now()

			// 更新数据库中的记录
			if err := config.DB.Save(&existingInventory).Error; err != nil {
				log.Printf("Failed to update pet food inventory for pet_id %d: %v", petInventory.PetID, err)
			} else {
				log.Printf("Updated pet food inventory for pet_id %d with RemainingAmount: %.2f", petInventory.PetID, petInventory.RemainingAmount)
			}
		} else {
			// 如果记录不存在，则插入新记录
			newInventory := models.PetFoodInventory{
				PetID:          petInventory.PetID,
				RemainingAmount: petInventory.RemainingAmount,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			if err := config.DB.Create(&newInventory).Error; err != nil {
				log.Printf("Failed to insert pet food inventory for pet_id %d: %v", petInventory.PetID, err)
			} else {
				log.Printf("Inserted new pet food inventory for pet_id %d with RemainingAmount: %.2f", petInventory.PetID, petInventory.RemainingAmount)
			}
		}
	}

	return nil
}
