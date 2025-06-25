package model

import "time"

type Message struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Author    string    `json:"author" gorm:"column:author"`
	Content   string    `json:"content" gorm:"column:content"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func (Message) TableName() string {
	return "messages"
}
