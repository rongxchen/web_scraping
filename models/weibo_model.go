package models

import (
	"fmt"
	"gorm.io/gorm"
)

type BlogUser struct {
	gorm.Model
	UserId         int     `gorm:"column:user_id" json:"user_id"`
	Username       string  `gorm:"column:username" json:"username"`
	Description    string  `gorm:"column:description" json:"description"`
	FollowCount    float64 `gorm:"column:follow_count" json:"follow_count"`
	FollowersCount string  `gorm:"column:followers_count" json:"followers_count"`
	Avatar         string  `gorm:"column:avatar" json:"avatar"`
	Gender         string  `gorm:"column:gender" json:"gender"`
	Verified       bool    `gorm:"column:verified" json:"verified"`
	VerifiedReason string  `gorm:"column:verified_reason" json:"verified_reason"`
}

func (b BlogUser) String() string {
	return fmt.Sprintf("Id: %d, UserId: %d, Username: %s, Description: %s, FollowCount: %f, "+
		"FollowersCount: %s, Avatar: %s, Gender: %s, Verified: %t, VerifiedReason: %s",
		b.ID, b.UserId, b.Username, b.Description, b.FollowCount, b.FollowersCount, b.Avatar, b.Gender,
		b.Verified, b.VerifiedReason)
}

type Blog struct {
	gorm.Model
	MId           string `gorm:"column:m_id" json:"m_id"`
	Text          string `gorm:"column:text" json:"text"`
	PictureList   string `gorm:"column:picture_list" json:"picture_list"`
	LikeCount     int    `gorm:"column:like_count" json:"like_count"`
	CommentCount  int    `gorm:"column:comment_count" json:"comment_count"`
	RepostCount   int    `gorm:"column:repost_count" json:"repost_count"`
	BlogUserId    int    `gorm:"column:blog_user_id" json:"blog_user_id"`
	Source        string `gorm:"column:source" json:"source"`
	StatusCity    string `gorm:"column:status_city" json:"status_city"`
	StatusCountry string `gorm:"column:status_country" json:"status_country"`
	CreatedAt     string `gorm:"column:created_at" json:"created_at"`
}

func (b Blog) String() string {
	return fmt.Sprintf("Id: %d, MId: %s, Text: %s, PictureList: %v, LikeCount: %d, CommentCount: %d, "+
		"RepostCount: %d, BlogUserId: %d, Source: %s, StatusCity: %s, StatusCountry: %s, CreatedAt: %s",
		b.ID, b.MId, b.Text, b.PictureList, b.LikeCount, b.CommentCount, b.RepostCount, &b.BlogUserId,
		b.Source, b.StatusCity, b.StatusCountry, b.CreatedAt)
}
