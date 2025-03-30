package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
	Comments  []Comment `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null;index"`
	User      User      `gorm:"foreignKey:UserID"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID                   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Content              string `gorm:"type:text;not null"`
	UserID               uint   `gorm:"not null;index"`
	PostID               uint   `gorm:"not null;index"`
	User                 User   `gorm:"foreignKey:UserID"`
	Post                 Post   `gorm:"foreignKey:PostID"`
	FlaggedForModeration bool   `gorm:"default:false"`
	ModerationStatus     string `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
