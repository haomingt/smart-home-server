package auth

import (
	"errors"
	"smart-home-server/config"
	"smart-home-server/models"

	"golang.org/x/crypto/bcrypt"
	"time"
)

func RegisterUser(username, password string) error {
	// 检查是否已存在
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err == nil {
		return errors.New("user already exists")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		Username:  username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	return config.DB.Create(&newUser).Error
}

func LoginUser(username, password string) (string, error) {
	// 查找用户
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return "", errors.New("user not found")
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// 生成 token
	token, err := GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
