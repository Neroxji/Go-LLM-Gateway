package main

import (
	"time"
)

// 发给大模型的请求体
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 大模型发回来的请求体
type ChatResponse struct {
	Choices []Choice `json:"choices"`
}
type Choice struct {
	Delta Delta `json:"delta"`
}
type Delta struct {
	Content string `json:"content"`
}

// 数据库用户
type User struct {
	ID        uint    `gorm:"primarykey"`
	Username  string  `gorm:"uniqueIndex;size:50"`
	Balance   float64 `gorm:"type:decimal(10,4);default:0"`
	Status    bool    `gorm:"default:1"`
	CreatedAt time.Time
}

// 令牌表
type Token struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `gorm:"index;not null"`
	Name      string `gorm:"size:50;default:'默认密钥'"`
	Key       string `gorm:"uniqueIndex;size:100;not null"`
	Status    int    `gorm:"default:1"`
	CreatedAt time.Time
}

// 数据库日志
type RequestLog struct {
	ID               uint   `gorm:"primarykey"`
	UserID           uint   `gorm:"index"`
	TokenID          uint   `gorm:"index"`
	TargetProvider   string `gorm:"size:20"`
	TargetModel      string `gorm:"size:50"`
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
	Latency          int64
	StatusCode       int
	ErrorMessage     string    `gorm:"type:text"`
	CreatedAt        time.Time `gorm:"index"`
}
