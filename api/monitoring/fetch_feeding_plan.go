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

func FetchFeedingPlanFromOtherGroup() error {
	// 假设获取其他小组的宠物喂食计划接口URL
	url := "https://other-group-api.com/pet-feeding-plans"

	// 发送HTTP请求，获取宠物喂食计划
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch pet feeding plan from other group: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应的宠物喂食计划
	var feedingPlans []struct {
		PetID     uint    `json:"pet_id"`     // 关联宠物ID
		Amount    float64 `json:"amount"`     // 每次喂食的粮食量（克）
		Frequency string  `json:"frequency"`  // 喂食频率
	}

	if err := json.NewDecoder(resp.Body).Decode(&feedingPlans); err != nil {
		return fmt.Errorf("failed to decode pet feeding plan: %v", err)
	}

	// 删除数据库中所有的宠物喂食计划
	if err := config.DB.Where("1 = 1").Delete(&models.FeedingPlan{}).Error; err != nil {
		log.Printf("Failed to delete existing pet feeding plan records: %v", err)
		return err
	}

	// 将获取到的数据插入到本地数据库
	for _, plan := range feedingPlans {
		// 适配其他小组数据到自己的模型
		newPlan := models.FeedingPlan{
			PetID:     plan.PetID,
			Amount:    plan.Amount,
			Frequency: plan.Frequency,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = config.DB.Create(&newPlan).Error
		if err != nil {
			log.Printf("Failed to insert new pet feeding plan: %v", err)
			continue
		}
		log.Printf("Successfully inserted new pet feeding plan for PetID: %d", plan.PetID)
	}

	return nil
}
