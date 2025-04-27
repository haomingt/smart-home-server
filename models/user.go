package models

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`                      // 主键
	Username  string    `json:"username" db:"username"`          // 用户名
	Password  string    `json:"-" db:"password"`                 // 密码哈希（不返回给前端）
	CreatedAt time.Time `json:"created_at" db:"created_at"`      // 注册时间
}
