package monitoring

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"smart-home-server/config"
	"smart-home-server/models"
	"time"
)

func FetchPetInfoFromOtherGroup() error {
	// 假设获取其他小组的接口URL
	url := "https://other-group-api.com/pet-info"

	// 发送HTTP请求，获取宠物信息
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch pet info from other group: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应的宠物信息
	var petInfo []models.Pet
	if err := json.NewDecoder(resp.Body).Decode(&petInfo); err != nil {
		return fmt.Errorf("failed to decode pet info: %v", err)
	}

	// 删除数据库中所有的宠物信息
	if err := config.DB.Where("1 = 1").Delete(&models.Pet{}).Error; err != nil {
		log.Printf("Failed to delete existing pet records: %v", err)
		return err
	}

	// 插入新获取的宠物信息
	for _, pet := range petInfo {
		pet.CreatedAt = time.Now()
		err = config.DB.Create(&pet).Error
		if err != nil {
			log.Printf("Failed to insert new pet info: %v", err)
			continue
		}
		log.Printf("Successfully inserted new pet info: %s", pet.Name)
	}

	return nil
}
